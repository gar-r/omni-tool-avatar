package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

const sessionCookieName = "session-id"
const sessionExpiry = time.Hour

var sessions map[string]*session = make(map[string]*session)

type session struct {
	id      string
	expires time.Time
	order   []*avatar
	current int
}

func registerSession() *session {
	s := &session{
		id:      makeID(),
		expires: time.Now().Add(sessionExpiry),
		order:   makeOrder(),
		current: 0,
	}
	sessions[s.id] = s
	go func() {
		<-time.After(sessionExpiry)
		delete(sessions, s.id)
	}()
	return s
}

func makeID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
