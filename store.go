package main

import (
	"bytes"
	"encoding/gob"
	"os"
	"sync"
)

var store sessionStore = newFileStore("sessions.gob")

type sessionStore interface {
	find(id uint32) (*session, error)
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

func (s *fileStore) find(id uint32) (*session, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	sessions, err := s.load()
	if err != nil {
		return nil, err
	}
	session, ok := sessions[id]
	if !ok {
		return nil, nil
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
	buf := new(bytes.Buffer)
	e := gob.NewEncoder(buf)
	err = e.Encode(sessions)
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

func (s *fileStore) load() (map[uint32]*session, error) {
	f, err := os.Open(s.fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[uint32]*session), nil
		}
		return nil, err
	}
	d := gob.NewDecoder(f)
	var result map[uint32]*session
	err = d.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
