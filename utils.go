package main

import (
	"hash/fnv"
	"log"
	"net"
	"net/http"
)

func getClientID(r *http.Request) uint32 {
	var token string
	token = getSMSession(r)
	if token == "" {
		token = getIP(r)
	}
	return hash(token)
}

func getSMSession(r *http.Request) string {
	c, err := r.Cookie("SMSESSION")
	if err != nil {
		return ""
	}
	return c.Value
}

func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("unable to parse client IP: %s\n", err)
	}
	return ip
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
