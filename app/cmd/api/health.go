package main

import "net/http"

func (app *Application) healtCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK!"))
}
