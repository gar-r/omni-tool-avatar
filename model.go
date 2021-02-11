package main

var avatars = []*avatar{
	makeAvatar("Okki", "user1", "100"),
	makeAvatar("Neko", "user2", "101"),
	makeAvatar("Mako", "user3", "102"),
	makeAvatar("Kika", "user4", "103"),
	makeAvatar("Bobi", "user5", "104"),
	makeAvatar("Lufi", "user6", "105"),
}

type avatar struct {
	Name, Logon, Msid string
}

func makeAvatar(name, logon, msid string) *avatar {
	return &avatar{name, logon, msid}
}
