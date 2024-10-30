package main

import "github.com/cliffdoyle/SnippetShare.git/internal/models"

type templateData struct{
	Snippet *models.Snippet
	Snippets []*models.Snippet
}