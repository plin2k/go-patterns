package main

import (
	"github.com/plin2k/go-patterns/singletone/service"
)

func main() {
	instance := service.GetInstance()
	instance.Greeter()
}
