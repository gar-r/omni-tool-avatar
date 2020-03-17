package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"sync"
)

var store sessionStore = newFileStore("sessions.gob")

type sessionStore interface {
	find(id uint32) *session
	save(session *session)
}

type fileStore struct {
	fileName string
	mux      sync.Mutex
}

func newFileStore(fileName string) *fileStore {
	return &fileStore{
		fileName: fileName,
	}
}

func (s *fileStore) find(id uint32) *session {
	s.mux.Lock()
	defer s.mux.Unlock()
	sessions := s.load()
	session, ok := sessions[id]
	if !ok {
		return nil
	}
	return session
}

func (s *fileStore) save(session *session) {
	s.mux.Lock()
	defer s.mux.Unlock()
	sessions := s.load()
	sessions[session.ID] = session
	buf := new(bytes.Buffer)
	e := gob.NewEncoder(buf)
	err := e.Encode(sessions)
	check(err)
	f, err := os.Create(s.fileName)
	check(err)
	_, err = f.Write(buf.Bytes())
	check(err)
}

func (s *fileStore) load() map[uint32]*session {
	f, err := os.Open(s.fileName)
	if os.IsNotExist(err) {
		return make(map[uint32]*session)
	}
	check(err)
	d := gob.NewDecoder(f)
	var result map[uint32]*session
	err = d.Decode(&result)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
