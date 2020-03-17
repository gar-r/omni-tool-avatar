package main

import (
	"html/template"
	"io"
	"log"
)

var page = template.Must(
	template.ParseFiles("templates/page.html",
		"templates/body.html"))

func executePageTemplate(w io.Writer, avatar *avatar, remaining int) {
	err := page.Execute(w, &context{
		Avatar:    *avatar,
		Remaining: remaining,
	})
	if err != nil {
		log.Println(err)
	}
}
