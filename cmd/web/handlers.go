package main


import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path !="/"{
		http.NotFound(w,r)
		return
	}
	w.Write([]byte(`hello from SnippetShare`))
}
//This handler will handle the viewing of snippets
func snipview(w http.ResponseWriter, r *http.Request){
	id,err:=strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0{
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w,"Display a specific code snippet with ID %d..",id)
}

func createsnip(w http.ResponseWriter, r *http.Request){
	if r.Method !=http.MethodPost{
		w.Header().Add("Allow","POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed"))
		http.Error(w,"Methodos Not Allowed",http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new code snippet"))
}

