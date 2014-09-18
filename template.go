package main

import (
	"html/template"
	"path/filepath"
)

const templatesPath = "templates/*"

var templates *template.Template

func init() {
	files, err := filepath.Glob(templatesPath)
	if err != nil {
		panic(err.Error())
	}

	templates = template.Must(
		template.ParseFiles(files...),
	)
}
