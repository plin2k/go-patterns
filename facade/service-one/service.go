package service_one

type User struct {
	Name string
	Role int
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

func (u *User) IsAdmin() bool {
	if u.Role == 777 {
		return true
	}
	return false
}

func (u *User) IsManager() bool {
	if u.Role == 666 {
		return true
	}
	return false
}
