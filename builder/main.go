package main

import (
	"fmt"
	"github.com/plin2k/go-patterns/builder/builder"
	"github.com/plin2k/go-patterns/builder/builder/manager"
)

func main() {
	simpleBuilder := builder.NewHouseBuilder("simple")
	advancedBuilder := builder.NewHouseBuilder("advanced")

	managerBuilder := manager.New(simpleBuilder)

	simpleHouse := managerBuilder.BuildHouse()

	fmt.Printf("Simple House Door Type: %s\n", simpleHouse.GetDoorInfo())
	fmt.Printf("Simple House Window Type: %s\n", simpleHouse.GetWindowInfo())
	fmt.Printf("Simple House Num Floor: %d\n", simpleHouse.GetFoorInfo())

	managerBuilder.SetBuilder(advancedBuilder)

	advancedHouse := managerBuilder.BuildHouse()

	fmt.Printf("Advanced House Door Type: %s\n", advancedHouse.GetDoorInfo())
	fmt.Printf("Advanced House Window Type: %s\n", advancedHouse.GetWindowInfo())
	fmt.Printf("Advanced House Num Floor: %d\n", advancedHouse.GetFoorInfo())

}
