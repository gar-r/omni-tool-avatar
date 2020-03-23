package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func newID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func newOrder() []*avatar {
	order := make([]*avatar, len(avatars))
	copy(order, avatars)
	rand.Shuffle(len(order), func(i, j int) {
		order[i], order[j] = order[j], order[i]
	})
	return order
}

func setCookie(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, getCookie(sessionID))
}

func clearCookie(w http.ResponseWriter, sessionID string) {
	c := getCookie(sessionID)
	c.Expires = time.Unix(0, 0)
	http.SetCookie(w, c)
}

func getCookie(sessionID string) *http.Cookie {
	return &http.Cookie{
		Name:     cookieName,
		Value:    sessionID,
		HttpOnly: true,
		Expires:  time.Now().Add(sessionExpiry),
	}
}
