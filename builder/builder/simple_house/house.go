package simple_house

import "github.com/plin2k/go-patterns/builder/house"

const (
	BuilderName = "simple"
)

type SimpleHouseBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func New() *SimpleHouseBuilder {
	return &SimpleHouseBuilder{}
}

func (b *SimpleHouseBuilder) SetWindowType() {
	b.windowType = "Wooden Window"
}

func (b *SimpleHouseBuilder) SetDoorType() {
	b.doorType = "Wooden Door"
}

func (b *SimpleHouseBuilder) SetNumFloor() {
	b.floor = 1
}

func (b *SimpleHouseBuilder) GetHouse() house.House {
	return house.House{
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}
}
