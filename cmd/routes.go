package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.HomePage)
	mux.HandleFunc("/create", app.CreationPage)
	mux.HandleFunc("/show", app.Show)
	mux.HandleFunc("/change", app.Change)
	return mux
}
