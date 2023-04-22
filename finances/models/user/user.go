package user

type User struct {
	UserId int
}

func NewUser(userId int) User {
	return User{UserId: userId}
}

type FactoryImp struct{}

func (f FactoryImp) CreateUser(userId int, username string) User {
	return User{UserId: userId}
}

type Factory interface {
	CreateUser(userId int, username string) User
}
