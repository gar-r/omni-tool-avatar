package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	id := getClientID(r)
	session := getSession(id)
	executePageTemplate(w, session.avatar(), session.remaining())
	session.moveNext()
}
