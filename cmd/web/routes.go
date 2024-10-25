package main

import "net/http"

// Moved the route declarations for the application into a standalone routes.go
// The routes() method returns a servemux containing our application routes
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/snipview", app.snipview)
	mux.HandleFunc("/snippet/createsnip", app.createsnip)

	return mux
}
