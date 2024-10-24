package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)


type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
}




func main() {
	addr:=flag.String("addr",":4000","HTTP network address")
	flag.Parse()

	infoLog:=log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	errorLog:=log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Ltime|log.Llongfile)

	//Initialize a new instance of the application struct
	//containing the dependencies
	app:=&application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	
	
	mux := http.NewServeMux()
	fileServer:=http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/snipview", app.snipview)
	mux.HandleFunc("/snippet/createsnip", app.createsnip)
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