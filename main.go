package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kyzercmd/http/controllers"
	"github.com/kyzercmd/http/views"
)


func main() {

	//Create new Variable of type chi.mux with chi.newRouter
	r := chi.NewRouter()
	//Import and use the Logger middleware
	r.Use(middleware.Logger)

	//Parsing templates
	homeTmpl := views.Must(views.ParseTemplate(filepath.Join("templates", "home.gohtml"))) 
	
	contactTmpl := views.Must(views.ParseTemplate(filepath.Join("templates", "contact.gohtml")))

	aboutTmpl := views.Must(views.ParseTemplate(filepath.Join("templates", "about.gohtml")))
	
	//Patterns and handlerFunctions for those patterns
	r.Get("/", controllers.HomeHandler(homeTmpl))
	r.Get("/contact", controllers.ContactHandler(contactTmpl))
	r.Get("/contact/{contactID}",  controllers.ContactHandler(contactTmpl))
	r.Get("/about", controllers.AboutHandler(aboutTmpl))

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