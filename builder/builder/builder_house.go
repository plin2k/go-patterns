package builder

import (
	"github.com/plin2k/go-patterns/builder/builder/advanced_house"
	"github.com/plin2k/go-patterns/builder/builder/simple_house"
	"github.com/plin2k/go-patterns/builder/house"
)

type HouseBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() house.House
}

func NewHouseBuilder(builderType string) HouseBuilder {
	switch builderType {
	case simple_house.BuilderName:
		return simple_house.New()
	case advanced_house.BuilderName:
		return advanced_house.New()
	default:
		return nil
	}
}
