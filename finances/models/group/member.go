package group

type Member struct {
	UserId  int
	IsAdmin bool
}

func NewGroupMember(userId int, isAdmin bool) Member {
	return Member{UserId: userId, IsAdmin: isAdmin}
}
