package modules

import (
	"net/http"

	"github.com/YagoSchramm/gopher-social/internal/infrastructure/router"
	"github.com/gorilla/mux"
)

type healthModule struct{}

func NewHealthModule() router.Module {
	return healthModule{}
}

func (m healthModule) Name() string {
	return "Health"
}

func (m healthModule) Path() string {
	return ""
}

func (m healthModule) Routes() []router.RouteDefinition {
	return []router.RouteDefinition{
		{
			Path:        "/health",
			Description: "Check whether the API is alive",
			Handler:     m.healthCheckHandler,
			HttpMethods: []string{http.MethodGet},
			Public:      true,
		},
	}
}

func (m healthModule) Middlewares() []mux.MiddlewareFunc {
	return nil
}

func (m healthModule) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("OK!"))
}
