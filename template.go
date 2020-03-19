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

func executePageTemplate(w io.Writer, avatar *avatar, remaining int) {
	err := page.Execute(w, &context{
		Avatar:    *avatar,
		Remaining: remaining,
	})
	if err != nil {
		log.Println(err)
	}
}

func executeErrorTemplate(w io.Writer, context error) {
	err := errpage.Execute(w, context.Error())
	if err != nil {
		log.Println(err)
	}
}
