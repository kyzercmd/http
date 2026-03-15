package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w , "<h1>HELLO THE</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact Page</h1> <p>To contact me message me on Discord: <a href=\"youinspireme\">discord</a></p>")
}

type Router struct {}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {
	case "/": homeHandler(w, r)
	case "/contact": contactHandler(w, r)
	default: 
	http.Error(w, "page not found", http.StatusNotFound)
	}
}

func main() {
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	var router Router
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}