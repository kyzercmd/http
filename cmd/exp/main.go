package main

import (
	"html/template"
	"os"
)

type user struct {
	Name string
}

func main() {
	tmpl, err := template.ParseFiles("test.gohtml")
	if err != nil {
		panic(err)
	}

	user := struct{
		Name string
	}{
		Name: "Kaiser",
	}

	err = tmpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}