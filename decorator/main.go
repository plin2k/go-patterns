package main

import (
	"fmt"
	decorator_one "github.com/plin2k/go-patterns/decorator/decorators/decorator-one"
	decorator_two "github.com/plin2k/go-patterns/decorator/decorators/decorator-two"
	"github.com/plin2k/go-patterns/decorator/service"
)

func main() {

	user := &service.User{
		Name: "Aleksandr Kalinkin",
		Role: 666,
	}

	adminUser := &decorator_one.AdminUser{
		User: user,
	}

	prefixAdminUser := &decorator_two.PrefixAdminUser{
		User: adminUser,
	}

	fmt.Printf("%s with role %d\n", prefixAdminUser.GetName(), prefixAdminUser.GetRole())
}