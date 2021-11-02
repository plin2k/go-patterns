package main

import (
	"fmt"
	"github.com/plin2k/go-patterns/facade/facade"
)

func main() {
	facadeVar := facade.NewFacade("Aleksandr Kalinkin", "Golang is Good!")

	if facadeVar.PublishPost() {
		fmt.Println("Success!")
	}

	facadeVar.ChangeUserName("plin2k")

	fmt.Println(facadeVar.User.Name)
}
