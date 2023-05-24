package entities

type Member struct {
	Nama string
	Pass int
}

type LinkedlistMember struct {
	Data Member
	Next *LinkedlistMember
}

type MemberVip struct {
	Nama string
	Pass int
}

type LinkedlistMemberVip struct {
	DataVip MemberVip
	Next    *LinkedlistMemberVip
}
