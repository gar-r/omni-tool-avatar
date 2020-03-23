package main

import "time"

type session struct {
	ID      string
	Order   []*avatar
	Current int
	Expires time.Time
}

func newSession() (*session, error) {
	s := &session{
		ID:      newID(),
		Order:   newOrder(),
		Current: 0,
		Expires: time.Now().Add(sessionExpiry),
	}
	err := sessionStore.save(s)
	return s, err
}

func getSession(id string) (*session, error) {
	s, err := sessionStore.find(id)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *session) moveNext() error {
	s.Current = (s.Current + 1) % len(avatars)
	return sessionStore.save(s)
}

func (s *session) avatar() *avatar {
	return s.Order[s.Current]
}

func (s *session) remaining() int {
	return len(avatars) - (s.Current + 1)
}
