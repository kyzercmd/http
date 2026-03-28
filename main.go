package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kyzercmd/http/views"
)

//Function to parse and execute template
func executeTemplate (w http.ResponseWriter, filePath string){
	//From views package use ParseTemplate and save the returned Template type
	tmpl, err := views.ParseTemplate(filePath)
	if err != nil {
		fmt.Println("Error parsing html :", err)
		http.Error(w, "Error parsing html", http.StatusInternalServerError)
		return
	}

	//Execute method on Template type
	tmpl.Execute(w, nil)
	
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	//Set Content-Type Header on the ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//OS agnostic filePath using filepath.Join
	filePath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, filePath)
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//URLParam to grab contactID from the request
	contactID := chi.URLParam(r, "contactID")
	filePath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, filePath)
	fmt.Fprint(w, contactID)
}

func main() {

	//Create new Variable of type chi.mux with chi.newRouter
	r := chi.NewRouter()

	//Import and use the Logger middleware
	r.Use(middleware.Logger)

	//Patterns and handlerFunctions for those patterns
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/contact/{contactID}", contactHandler)

	//The NotFound page
	r.NotFound(func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "page not found", http.StatusNotFound)
	})

	//Start server on port 3333
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":3333", r)
	if err != nil {
		panic(err)
	}
}