package main

import (
	"log"
	"net/http"

	"github.com/YagoSchramm/gopher-social/internal/infrastructure/router"
	routermodules "github.com/YagoSchramm/gopher-social/internal/infrastructure/router/modules"
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

	modules := []router.Module{
		routermodules.NewHealthModule(),
	}

	for _, module := range modules {
		moduleRouter := r
		if module.Path() != "" {
			moduleRouter = r.PathPrefix(module.Path()).Subrouter()
		}

		protectedRouter := moduleRouter.NewRoute().Subrouter()
		for _, mw := range module.Middlewares() {
			protectedRouter.Use(mw)
		}

		for _, route := range module.Routes() {
			target := protectedRouter
			if route.Public {
				target = moduleRouter
			}

			target.HandleFunc(route.Path, route.Handler).Methods(route.HttpMethods...)
		}
	}

	return r
}
func (app *Application) run(r *mux.Router) error {
	log.Printf("server has started at %s", app.config.addr)
	return http.ListenAndServe(app.config.addr, r)
}
