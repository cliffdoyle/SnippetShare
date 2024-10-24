package main 

import (
	"net/http"
	"log"
)
func main() {

	mux := http.NewServeMux()
	fileServer:=http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/snipview", snipview)
	mux.HandleFunc("/snippet/createsnip", createsnip)
	// fmt.Println(`Hello world`)

	log.Print("starting server on :4040")
	err:=http.ListenAndServe(":4040",mux)
	log.Fatal(err)
}