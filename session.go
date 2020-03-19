package main

import (
	"math/rand"
	"time"
)

const expiry = 24 * time.Hour

type session struct {
	ID      uint32
	Current int
	Expires time.Time
}

func getSession(id uint32) (*session, error) {
	s, err := store.find(id)
	if err != nil {
		return nil, err
	}
	if s == nil {
		s = newSession(id)
		store.save(s)
	}
	return s, nil
}

func newSession(id uint32) *session {
	return &session{
		ID:      id,
		Current: 0,
		Expires: time.Now().Add(expiry),
	}
}

func (s *session) moveNext() error {
	s.Current = (s.Current + 1) % len(avatars)
	return store.save(s)
}

func (s *session) avatar() *avatar {
	return s.calcOrder()[s.Current]
}

func (s *session) remaining() int {
	return len(avatars) - (s.Current + 1)
}

func (s *session) calcOrder() []*avatar {
	y, m, d := time.Now().Date()
	seed := int64(s.ID) + int64(y) + int64(m) + int64(d)
	rng := rand.New(rand.NewSource(seed))
	res := make([]*avatar, len(avatars))
	copy(res, avatars)
	rng.Shuffle(len(avatars), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}
