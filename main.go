package main

import (
	"errors"
	"fmt"
	"io"
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
	executeErrorTemplate(w, errors.New("test error"))
	id := getClientID(r)
	session, err := getSession(id)
	if !check(err, w) {
		return
	}
	executePageTemplate(w, session.avatar(), session.remaining())
	err = session.moveNext()
	if !check(err, w) {
		return
	}
}

func check(err error, w io.Writer) bool {
	if err != nil {
		log.Println(err)
		executeErrorTemplate(w, err)
		return false
	}
	return true
}
