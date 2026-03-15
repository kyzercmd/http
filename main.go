package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w , "<h1>HELLO THE</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact Page</h1> <p>To contact me message me on Discord: <a href=\"youinspireme\">discord</a></p>")
}

func main() {

	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":3333", r)
	if err != nil {
		panic(err)
	}
}