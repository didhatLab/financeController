package user

type User struct {
	UserId   int
	Username string
}

type FactoryImp struct{}

func (f FactoryImp) CreateUser(userId int, username string) User {
	return User{UserId: userId, Username: username}
}

type Factory interface {
	CreateUser(userId int, username string) User
}
