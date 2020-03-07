package main

import "math/rand"

type avatar struct {
	Name, Logon, Msid string
}

func makeAvatar(name, logon, msid string) *avatar {
	return &avatar{name, logon, msid}
}

func makeOrder() []*avatar {
	s := make([]*avatar, len(avatars))
	copy(s, avatars)
	rand.Shuffle(len(avatars), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}

func nextAvatar(s *session) *avatar {
	a := s.order[s.current]
	s.current = (s.current + 1) % len(s.order)
	return a
}

type context struct {
	Avatar    avatar
	Remaining int
}
