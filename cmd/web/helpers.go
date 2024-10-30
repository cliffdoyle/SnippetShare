package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application)render(w http.ResponseWriter, status int, page string,data *templateData){
	// Retrieve the appropriate template set from the cache based on the page
// name (like 'home.tmpl'). If no entry exists in the cache with the
// provided name, then create a new error and call the serverError()helper
// method that we made earlier and return.

ts,ok:=app.templateCache[page]
if !ok{
	err:=fmt.Errorf("the template %s does not exist",page)
	app.serverError(w,err)
	return
}

w.WriteHeader(status)

//execute the template set and write the response body
err:=ts.ExecuteTemplate(w,"base",data)
if err !=nil{
	app.serverError(w,err)
}

}

func (app *application)serverError(w http.ResponseWriter, err error){
	trace:=fmt.Sprintf("%s\n%s",err.Error(),debug.Stack())
	app.errorLog.Output(2,trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
}

func(app *application)clientError(w http.ResponseWriter,status int){
	http.Error(w,http.StatusText(status),status)
}

func(app *application)NotFound(w http.ResponseWriter){
	app.clientError(w,http.StatusNotFound)
}