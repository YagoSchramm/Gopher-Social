package domain

type Password struct {
	Hash []byte `json:"-"`
}
