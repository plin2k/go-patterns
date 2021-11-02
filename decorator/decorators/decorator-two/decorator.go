package decorator_two

import "github.com/plin2k/go-patterns/decorator/decorators"

type PrefixAdminUser struct {
	User decorators.User
}

func (u *PrefixAdminUser) GetName() string {
	if u.User.GetRole() == 777 {
		return "Admin - " + u.User.GetName()
	}
	return "Moderator - " + u.User.GetName()
}

func (u *PrefixAdminUser) GetRole() int {
	return u.User.GetRole()
}