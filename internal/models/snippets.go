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
	stm:=`INSERT INTO snippets(title,content,created,expires)
	Values(?,?, UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(),INTERNAL?DAY))`
	
	result,err:=m.DB.Exec(stm,title,content,expires)
	if err !=nil{
		return 0,err
	}

	id,err:=result.LastInsertId()
	if err !=nil{
		return 0,err
	}
	return int(id),nil
}
//returns a specific snippet based on its id
func(m *SnippetModel)Get(id int)(*Snippet,error){
	return nil,nil
}
//Returns the 10 most recently created snippets
func(m *SnippetModel)Latest()([]*Snippet,error){
	return nil,nil
}