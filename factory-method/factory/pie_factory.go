package factory

import (
	"errors"
	"github.com/plin2k/go-patterns/factory-method/factory/apple_pie"
	"github.com/plin2k/go-patterns/factory-method/factory/interfaces"
	"github.com/plin2k/go-patterns/factory-method/factory/pumpkin_pie"
)

func GetPie(pieType string) (interfaces.IPie, error) {
	switch pieType {
	case apple_pie.Name:
		return apple_pie.New(), nil
	case pumpkin_pie.Name:
		return pumpkin_pie.New(), nil
	default:
		return nil, errors.New("unknown pie type")
	}
}
