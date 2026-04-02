package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//Template type with template.Template field
type Template struct {
	htmlTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(f fs.FS, patterns ...string) (Template, error){
	tmpl, err := template.ParseFS(f, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("Parsing FS: %w", err)
	}
	return Template{
		htmlTpl: tmpl,
	}, nil
}

func ParseTemplate (filePath string) (Template, error){

	//Parse from filePath and return a template.Template
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		//If err return empty Template and error
		return Template{}, fmt.Errorf("Error parsing: %w", err)
	}

	//Or if Success return Template with parsed template
	return Template{
		htmlTpl: tmpl,
	}, nil
}

//Method for Template type with any type parameter
func (t Template) Execute (w http.ResponseWriter, data interface{}){
	//Execute template on the Template type
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Error executing", http.StatusInternalServerError)
		return
	}
	
}