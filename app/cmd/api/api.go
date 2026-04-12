package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Application struct {
	config Config
}
type Config struct {
	addr string
}

func (app *Application) mount() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", app.healtCheckHandler).Methods("GET")
	return r
}
func (app *Application) run(r *mux.Router) error {
	log.Printf("server has started at %s", app.config.addr)
	return http.ListenAndServe(app.config.addr, r)
}
