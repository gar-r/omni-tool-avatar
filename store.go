package main

import (
	"bytes"
	"encoding/gob"
	"os"
	"sync"
	"time"
)

var sessionStore store = newFileStore("sessions.gob")

type store interface {
	find(id string) (*session, error)
	save(session *session) error
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

func (s *fileStore) find(id string) (*session, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	sessions, err := s.load()
	if err != nil {
		return nil, err
	}
	session, ok := sessions[id]
	if !ok {
		return nil, &notFound{}
	}
	return session, nil
}

func (s *fileStore) save(session *session) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	sessions, err := s.load()
	if err != nil {
		return err
	}
	sessions[session.ID] = session
	return s.persist(sessions)
}

func (s *fileStore) load() (map[string]*session, error) {
	f, err := os.Open(s.fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]*session), nil
		}
		return nil, err
	}
	d := gob.NewDecoder(f)
	var result map[string]*session
	err = d.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *fileStore) persist(sessions map[string]*session) error {
	s.purge(sessions)
	buf := new(bytes.Buffer)
	e := gob.NewEncoder(buf)
	err := e.Encode(sessions)
	if err != nil {
		return err
	}
	f, err := os.Create(s.fileName)
	if err != nil {
		return err
	}
	_, err = f.Write(buf.Bytes())
	return err
}

func (s *fileStore) purge(sessions map[string]*session) {
	for k, v := range sessions {
		if v.Expires.Before(time.Now()) {
			delete(sessions, k)
		}
	}
}

type notFound struct {
}

func (*notFound) Error() string {
	return "session not found"
}
