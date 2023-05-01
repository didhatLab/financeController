package group

type Member struct {
	UserId   int
	IsAdmin  bool
	Username string
}

func NewGroupMember(userId int, isAdmin bool) Member {
	return Member{UserId: userId, IsAdmin: isAdmin}
}
