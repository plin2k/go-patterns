package decorator_one

import "github.com/plin2k/go-patterns/decorator/decorators"

type AdminUser struct {
	User decorators.User
}

func (u *AdminUser) GetName() string {
	return u.User.GetName()
}

func (u *AdminUser) GetRole() int {
	if u.User.GetRole() < 777 {
		return 777
	}
	return 666
}
