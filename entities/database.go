package entities

type Member struct {
	Nama string
	Pass int
}

type LinkedlistMember struct {
	Data Member
	Next *LinkedlistMember
}
