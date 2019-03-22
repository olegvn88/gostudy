package users

type User struct {
	ID   int
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(newName string) {
	u.Name = newName
}

type UserInterface interface {
	GetName() string
	SetName(string)
}
