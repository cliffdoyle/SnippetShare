package main

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/cliffdoyle/SnippetShare.git/internal/models"
)

type templateData struct{
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

//Logic for parsing the html files once when application is starting
func newTemplateCache()(map[string]*template.Template,error){
	//initialize  anew map to act as the cache
	cache:=map[string]*template.Template{}

	//Use filepath.Glob() to get a slice of all filepaths that match pattern

	pages,err:=filepath.Glob("./ui/html/pages/*.html")
	if err !=nil{
		return nil,err
	}

	for _,page:=range pages{
		name:=filepath.Base(page)

		files:=[]string{
			"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		page,
		}

		ts,err:=template.ParseFiles(files...)
		if err !=nil{
			return nil,err
		}
		cache[name]=ts
	}

	// fmt.Println(cache)

return cache,nil


}