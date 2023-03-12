package finance

type Spending struct {
	UserId int
	Name   string
	Type   string
}

type FactoryImp struct{}

func (fi FactoryImp) CreateSpending(userId int, name string, Type string) Spending {
	return Spending{Name: name, Type: Type, UserId: userId}
}

type Factory interface {
	CreateSpending(userId int, name string, Type string) Spending
}
