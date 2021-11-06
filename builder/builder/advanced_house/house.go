package advanced_house

import "github.com/plin2k/go-patterns/builder/house"

const (
	BuilderName = "advanced"
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
	b.windowType = "Plastic Window"
}

func (b *SimpleHouseBuilder) SetDoorType() {
	b.doorType = "Steel Door"
}

func (b *SimpleHouseBuilder) SetNumFloor() {
	b.floor = 3
}

func (b *SimpleHouseBuilder) GetHouse() house.House {
	return house.House{
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}
}
