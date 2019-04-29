package goPaginate

import (
"html/template"
"log"
"net/http"
"path/filepath"
)

func LayoutFiles() []string {
	files, err := filepath.Glob("templates/layouts/*.tmpl")
	if err != nil {
		log.Panic(err)
	}
	return files
}

type View struct {
	Index  Page
}

type Page struct {
	Template *template.Template
	Layout string
}

func (self *Page) Render(w http.ResponseWriter, data interface{}) error {
	return self.Template.ExecuteTemplate(w, self.Layout, data)
}