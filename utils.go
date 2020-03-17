package main

import (
	"encoding/binary"
	"log"
	"net"
	"net/http"
)

func getClientID(r *http.Request) uint32 {
	ip := getIP(r)
	return ip2int(ip)
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func getIP(r *http.Request) net.IP {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("unable to parse client IP: %s\n", err)
	}
	return net.ParseIP(ip)
}
