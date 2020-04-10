package main

import (
	"html/template"
	"io"
	"log"
)

var page = template.Must(
	template.ParseFiles("templates/page.html",
		"templates/body.html"))

var errpage = template.Must(
	template.ParseFiles("templates/error.html"),
)

func executePageTemplate(w io.Writer, avatar *avatar) {
	err := page.Execute(w, &avatar)
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
