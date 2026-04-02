package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kyzercmd/http/views"
)

func HomeHandler(tmpl views.Template) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, nil)
	}
}

func ContactHandler(tmpl views.Template) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		contactID := chi.URLParam(r, "contactID")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, nil)
		fmt.Fprint(w, contactID)
	}
}

func AboutHandler(tmpl views.Template) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, nil)
	}
}