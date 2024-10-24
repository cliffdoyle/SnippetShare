package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application)home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path !="/"{
		http.NotFound(w,r)
		return
	}

	files:=[]string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts,err:=template.ParseFiles(files...)
	if err !=nil{
		app.errorLog.Print(err.Error())
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		return
	}

	err=ts.ExecuteTemplate(w,"base",nil)
	if err !=nil{
		app.errorLog.Print(err.Error())
		http.Error(w,"internal Server Error",http.StatusInternalServerError)
	}

	// fmt.Fprintln(w,"Hello fellow coders")


}
//This handler will handle the viewing of snippets
func (app *application)snipview(w http.ResponseWriter, r *http.Request){
	id,err:=strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0{
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w,"Display a specific code snippet with ID %d..",id)
}

func (app *application)createsnip(w http.ResponseWriter, r *http.Request){
	if r.Method !=http.MethodPost{
		w.Header().Add("Allow","POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed"))
		http.Error(w,"Methodos Not Allowed",http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new code snippet"))
}


