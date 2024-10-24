package main 

import (
	"net/http"
	"log"
)
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