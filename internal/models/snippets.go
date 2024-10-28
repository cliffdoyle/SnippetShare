package models


import (
	"database/sql"
	"time"
)
//Snippet type to hold the data for an individual snippet
type Snippet struct{
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}
//Wraps a sql.DB connection pool
type SnippetModel struct{
	DB *sql.DB
}
//Inserts a new snippet into the database
func (m *SnippetModel)Insert(title,content string, expires int)(int,error){
	return 0,nil
}
//returns a specific snippet based on its id
func(m *SnippetModel)Get(id int)(*Snippet,error){
	return nil,nil
}
//Returns the 10 most recently created snippets
func(m *SnippetModel)Latest()([]*Snippet,error){
	return nil,nil
}