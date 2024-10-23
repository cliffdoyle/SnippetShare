package main

import (
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello from SnippetShare`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, home)
	// fmt.Println(`Hello world`)

	log.Print("starting server on :4040")
	err:=http.ListenAndServe(":4040",mux)
	log.Fatal(err)
}
