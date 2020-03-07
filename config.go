package main

import (
	"html/template"
)

const port = 8080

var page = template.Must(
	template.ParseFiles("templates/page.html",
		"templates/body.html"))

var avatars = []*avatar{
	makeAvatar("Okker Makker", "okki", "100"),
	makeAvatar("Neko Vampire", "neko", "101"),
	makeAvatar("Richard Garai", "garric", "103"),
}
