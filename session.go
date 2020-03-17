package main

import (
	"math/rand"
	"time"
)

const expiry = 24 * time.Hour

var sessions = make(map[uint32]*session)

type session struct {
	id      uint32
	current int
}

func getSession(id uint32) *session {
	s, ok := sessions[id]
	if !ok {
		s = newSession(id)
		sessions[id] = s
	}
	return s
}

func newSession(id uint32) *session {
	time.AfterFunc(expiry, func() {
		delete(sessions, id)
	})
	return &session{
		id:      id,
		current: 0,
	}
}

func (s *session) moveNext() {
	s.current = (s.current + 1) % len(avatars)
}

func (s *session) avatar() *avatar {
	return s.calcOrder()[s.current]
}

func (s *session) remaining() int {
	return len(avatars) - (s.current + 1)
}

func (s *session) calcOrder() []*avatar {
	y, m, d := time.Now().Date()
	seed := int64(s.id) + int64(y) + int64(m) + int64(d)
	rng := rand.New(rand.NewSource(seed))
	res := make([]*avatar, len(avatars))
	copy(res, avatars)
	rng.Shuffle(len(avatars), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}
