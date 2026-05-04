package middlewares

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/gorilla/mux"
)

type contextKey string

const claimsContextKey contextKey = "jwtClaims"

type Claims map[string]any

type jwtAuthMiddleware struct {
	secret []byte
}

func NewJWTAuthMiddlewareFromEnv() mux.MiddlewareFunc {
	secret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	if secret == "" {
		secret = strings.TrimSpace(os.Getenv("JWT"))
	}

	return NewJWTAuthMiddleware(secret)
}

func NewJWTAuthMiddleware(secret string) mux.MiddlewareFunc {
	m := &jwtAuthMiddleware{
		secret: []byte(strings.TrimSpace(secret)),
	}

	return m.middleware
}

func ClaimsFromContext(ctx context.Context) (Claims, bool) {
	claims, ok := ctx.Value(claimsContextKey).(Claims)
	return claims, ok
}

func (m *jwtAuthMiddleware) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(m.secret) == 0 {
			writeHTTPError(w, derr.NewInternalError("jwt secret is not configured"))
			return
		}

		claims, err := m.authorize(r)
		if err != nil {
			writeHTTPError(w, derr.UnauthorizedError)
			return
		}

		ctx := context.WithValue(r.Context(), claimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *jwtAuthMiddleware) authorize(r *http.Request) (Claims, error) {
	token, err := bearerToken(r.Header.Get("Authorization"))
	if err != nil {
		return nil, err
	}

	header, payload, signature, err := splitJWT(token)
	if err != nil {
		return nil, err
	}

	if err := verifyJWTHeader(header); err != nil {
		return nil, err
	}

	if err := verifyJWTSignature(m.secret, header, payload, signature); err != nil {
		return nil, err
	}

	claims, err := decodeClaims(payload)
	if err != nil {
		return nil, err
	}

	if err := validateRegisteredClaims(claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func bearerToken(value string) (string, error) {
	parts := strings.Fields(value)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.New("missing bearer token")
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", errors.New("empty bearer token")
	}

	return token, nil
}

func splitJWT(token string) (string, string, string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", "", "", errors.New("invalid jwt format")
	}

	if parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", errors.New("invalid jwt format")
	}

	return parts[0], parts[1], parts[2], nil
}

func verifyJWTHeader(encodedHeader string) error {
	raw, err := base64.RawURLEncoding.DecodeString(encodedHeader)
	if err != nil {
		return err
	}

	var header struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	if err := json.Unmarshal(raw, &header); err != nil {
		return err
	}

	if !strings.EqualFold(header.Alg, "HS256") {
		return errors.New("unsupported jwt algorithm")
	}

	if header.Typ != "" && !strings.EqualFold(header.Typ, "JWT") {
		return errors.New("invalid jwt type")
	}

	return nil
}

func verifyJWTSignature(secret []byte, encodedHeader string, encodedPayload string, encodedSignature string) error {
	signedPart := encodedHeader + "." + encodedPayload
	mac := hmac.New(sha256.New, secret)
	_, _ = mac.Write([]byte(signedPart))

	expected := mac.Sum(nil)
	got, err := base64.RawURLEncoding.DecodeString(encodedSignature)
	if err != nil {
		return err
	}

	if !hmac.Equal(expected, got) {
		return errors.New("invalid jwt signature")
	}

	return nil
}

func decodeClaims(encodedPayload string) (Claims, error) {
	raw, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return nil, err
	}

	claims := Claims{}
	if err := json.Unmarshal(raw, &claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func validateRegisteredClaims(claims Claims) error {
	now := time.Now().Unix()

	if exp, ok := numericClaim(claims, "exp"); ok && now >= exp {
		return errors.New("jwt token expired")
	}

	if nbf, ok := numericClaim(claims, "nbf"); ok && now < nbf {
		return errors.New("jwt token not active yet")
	}

	return nil
}

func numericClaim(claims Claims, key string) (int64, bool) {
	value, ok := claims[key]
	if !ok {
		return 0, false
	}

	switch v := value.(type) {
	case float64:
		return int64(v), true
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return 0, false
		}
		return n, true
	case int64:
		return v, true
	case int:
		return int64(v), true
	default:
		return 0, false
	}
}

func writeHTTPError(w http.ResponseWriter, err derr.HTTPError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	_ = json.NewEncoder(w).Encode(err)
}
