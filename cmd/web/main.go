package main

import (
	"flag"
	"log"
	"net/http"
)
func main() {
	addr:=flag.String("addr",":4000","HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	fileServer:=http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/snipview", snipview)
	mux.HandleFunc("/snippet/createsnip", createsnip)
	// fmt.Println(`Hello world`)

	log.Printf("starting server on %s",*addr)
	err:=http.ListenAndServe(*addr,mux)
	log.Fatal(err)
}