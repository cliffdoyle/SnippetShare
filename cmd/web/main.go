package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn:=flag.String("dsn","doyle:Kombewa@254@/snippetshare?parseTime=true","MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Ltime|log.Llongfile)

	// Initialize a new instance of the application struct
	// containing the dependencies

	db,err:=openDB(*dsn)
	if err !=nil{
		errorLog.Fatal(err)
	}
	defer db.Close()
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// fmt.Println(`Hello world`)
	serv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("starting server on %s", *addr)
	err = serv.ListenAndServe()
	errorLog.Fatal(err)
}
