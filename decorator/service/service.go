package service

type User struct {
	Name string
	Role int
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetRole() int {
	return u.Role
}
