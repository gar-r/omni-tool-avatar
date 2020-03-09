package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const cookieName = "session-id"

func main() {
	rand.Seed(time.Now().UnixNano())
	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(cookieName)
	if err != nil {
		c = newSession(w)
	}
	session, ok := sessions[c.Value]
	if !ok {
		c = newSession(w)
		session, _ = sessions[c.Value]
	}
	remaining := len(session.order) - session.current - 1
	avatar := nextAvatar(session)
	err = page.Execute(w, &context{
		Avatar:    *avatar,
		Remaining: remaining,
	})
	if err != nil {
		log.Println(err)
	}
}

func newSession(w http.ResponseWriter) *http.Cookie {
	s := registerSession()
	c := &http.Cookie{
		Name:    cookieName,
		Value:   s.id,
		Expires: s.expires,
	}
	http.SetCookie(w, c)
	return c
}
