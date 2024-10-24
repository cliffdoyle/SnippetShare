package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)
func main() {
	addr:=flag.String("addr",":4000","HTTP network address")
	flag.Parse()

	infoLog:=log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	errorLog:=log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Ltime|log.Llongfile)

	mux := http.NewServeMux()
	fileServer:=http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/snipview", snipview)
	mux.HandleFunc("/snippet/createsnip", createsnip)
	// fmt.Println(`Hello world`)
	serv:=&http.Server{
		Addr: *addr,
		Handler: mux,
		ErrorLog: errorLog,
	}

	infoLog.Printf("starting server on %s",*addr)
	err:=serv.ListenAndServe()
	errorLog.Fatal(err)
}