package main

var avatars = []*avatar{
	makeAvatar("Okker Makker", "okki", "100"),
	makeAvatar("Neko Vampire", "neko", "101"),
	makeAvatar("Richard Garai", "garric", "103"),
}

type avatar struct {
	Name, Logon, Msid string
}

func makeAvatar(name, logon, msid string) *avatar {
	return &avatar{name, logon, msid}
}

func (a *avatar) clone() *avatar {
	return makeAvatar(a.Name, a.Logon, a.Msid)
}

type context struct {
	Avatar    avatar
	Remaining int
}
