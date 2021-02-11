package main

import (
	"html/template"
	"io"
	"log"
)

type context struct {
	Current *avatar
	Next    *avatar
	Prev    *avatar
}

var page = template.Must(
	template.ParseFiles("templates/page.html",
		"templates/body.html"))

var errpage = template.Must(
	template.ParseFiles("templates/error.html"),
)

func executePageTemplate(w io.Writer, data *context) {
	err := page.Execute(w, &data)
	if err != nil {
		log.Println(err)
	}
}

func executeErrorTemplate(w io.Writer, e error) {
	err := errpage.Execute(w, e.Error())
	if err != nil {
		log.Println(err)
	}
}
