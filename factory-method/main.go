package main

import (
	"fmt"
	"github.com/plin2k/go-patterns/factory-method/factory"
	"github.com/plin2k/go-patterns/factory-method/factory/apple_pie"
	"github.com/plin2k/go-patterns/factory-method/factory/interfaces"
	"github.com/plin2k/go-patterns/factory-method/factory/pumpkin_pie"
)

func main() {
	applePie, _ := factory.GetPie(apple_pie.Name)
	pumpkinPie, _ := factory.GetPie(pumpkin_pie.Name)

	printDetails(applePie)
	printDetails(pumpkinPie)
}

func printDetails(g interfaces.IPie) {
	fmt.Printf("Form: %s\n", g.GetForm())
	fmt.Printf("Size: %d\n", g.GetSize())
	fmt.Printf("Fill: %s\n", g.GetFill())
}
