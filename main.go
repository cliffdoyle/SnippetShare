package main

import (
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello from SnippetShare`))
}
//This handle will handle the viewing of snippets
func snipview(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a specific code snippet"))
}

func createsnip(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Create a new code snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, home)
	mux.HandleFunc("/snipview", snipview)
	mux.HandleFunc("/createsnip", createsnip)
	// fmt.Println(`Hello world`)

	log.Print("starting server on :4040")
	err:=http.ListenAndServe(":4040",mux)
	log.Fatal(err)
}
