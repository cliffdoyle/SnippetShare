package main

import (
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path !="/"{
		http.NotFound(w,r)
		return
	}
	w.Write([]byte(`hello from SnippetShare`))
}
//This handler will handle the viewing of snippets
func snipview(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a specific code snippet"))
}

func createsnip(w http.ResponseWriter, r *http.Request){
	if r.Method !=http.MethodPost{
		w.WriteHeader(405)
		w.Write([]byte("Method not Allowed"))
		return
	}
	w.Write([]byte("Create a new code snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/snipview", snipview)
	mux.HandleFunc("/snippet/createsnip", createsnip)
	// fmt.Println(`Hello world`)

	log.Print("starting server on :4040")
	err:=http.ListenAndServe(":4040",mux)
	log.Fatal(err)
}
