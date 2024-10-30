package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/cliffdoyle/SnippetShare.git/internal/models"
)

//Added snippets field to make the snippetmodel object
//available to the handlers
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
	templateCache map[string]*template.Template
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

	//Initialize a new template cahe...
	templateCache,err:=newTemplateCache()
	if err !=nil{
		errorLog.Fatal(err)
	}
	//Initialize a models.SnippetModel instance and add
	//it to the application dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
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
