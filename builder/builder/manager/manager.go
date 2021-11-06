package manager

import (
	"github.com/plin2k/go-patterns/builder/builder"
	"github.com/plin2k/go-patterns/builder/house"
)

type Manager struct {
	builder builder.HouseBuilder
}

func New(b builder.HouseBuilder) *Manager {
	return &Manager{
		builder: b,
	}
}

func (m *Manager) SetBuilder(b builder.HouseBuilder) {
	m.builder = b
}

func (m *Manager) BuildHouse() house.House {
	m.builder.SetDoorType()
	m.builder.SetWindowType()
	m.builder.SetNumFloor()
	return m.builder.GetHouse()
}
