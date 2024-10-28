package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application)home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path !="/"{
		app.NotFound(w)
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
		app.serverError(w,err)
		return
	}

	err=ts.ExecuteTemplate(w,"base",nil)
	if err !=nil{
		// app.errorLog.Print(err.Error())
		app.serverError(w,err)
	}

	// fmt.Fprintln(w,"Hello fellow coders")


}
//This handler will handle the viewing of snippets
func (app *application)snipview(w http.ResponseWriter, r *http.Request){
	id,err:=strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0{
		app.NotFound(w)
	}
	fmt.Fprintf(w,"Display a specific code snippet with ID %d..\n",id)
}

func (app *application)createsnip(w http.ResponseWriter, r *http.Request){
	if r.Method !=http.MethodPost{
		w.Header().Add("Allow","POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed"))
		app.clientError(w,http.StatusMethodNotAllowed)
		return
	}

	title:="O snail"
	content:= `O snail\n|Climb Mount Fuji,\nBut slowly, slowly!\n\n–
Kobayashi Issa`
	expires:=7

	id,err:=app.snippets.Insert(title,content,expires)
	if err !=nil{
		app.serverError(w,err)
		return
	}
	http.Redirect(w,r,fmt.Sprintf("/snippet/snipview?id=%d",id),http.StatusSeeOther)


	// w.Write([]byte("Create a new code snippet"))
}


