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

func (s *session) moveBack() error {
	s.Current = s.Current - 1
	if s.Current < 0 {
		s.Current = len(avatars) - 1
	}
	return sessionStore.save(s)
}

func (s *session) skip() error {
	current := s.Order[s.Current]
	s.Order = append(s.Order[:s.Current], s.Order[s.Current+1:]...)
	s.Order = append(s.Order, current)
	return sessionStore.save(s)
}

func (s *session) avatar() *avatar {
	return s.Order[s.Current]
}
