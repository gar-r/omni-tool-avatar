package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const port = 8080
const cookieName = "session-id"
const sessionExpiry = 6 * time.Hour

func main() {
	fs := http.FileServer(http.Dir("assets"))
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/reset", resetHander)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	s, err := extractSession(r)
	if err != nil {
		handleError(err, w)
		return
	}
	setCookie(w, s.ID)
	avatar := s.avatar().clone()
	remaining := s.remaining()
	err = s.moveNext()
	if err != nil {
		handleError(err, w)
		return
	}
	executePageTemplate(w, avatar, remaining)
}

func resetHander(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(cookieName)
	if err == nil {
		clearCookie(w, c.Value)
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func extractSession(r *http.Request) (*session, error) {
	c, err := r.Cookie(cookieName)
	if err != nil {
		return newSession()
	}
	return getSession(c.Value)
}

func handleError(err error, w io.Writer) {
	log.Println(err)
	executeErrorTemplate(w, err)
}
