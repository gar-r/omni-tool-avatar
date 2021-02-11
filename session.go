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
	s.Current = s.nextIndex()
	return sessionStore.save(s)
}

func (s *session) moveBack() error {
	s.Current = s.prevIndex()
	return sessionStore.save(s)
}

func (s *session) nextIndex() int {
	return (s.Current + 1) % len(avatars)
}

func (s *session) prevIndex() int {
	index := s.Current - 1
	if index < 0 {
		return len(avatars) - 1
	}
	return index
}

func (s *session) skip() error {
	current := s.Order[s.Current]
	s.Order = append(s.Order[:s.Current], s.Order[s.Current+1:]...)
	s.Order = append(s.Order, current)
	return sessionStore.save(s)
}

func (s *session) context() *context {
	return &context{
		Current: s.Order[s.Current],
		Next:    s.Order[s.nextIndex()],
		Prev:    s.Order[s.prevIndex()],
	}
}
