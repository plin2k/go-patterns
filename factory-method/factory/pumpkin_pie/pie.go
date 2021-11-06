package pumpkin_pie

import (
	"github.com/plin2k/go-patterns/factory-method/factory/interfaces"
	"github.com/plin2k/go-patterns/factory-method/pie"
)

const Name = "pumpkin"

type ApplePie struct {
	pie.Pie
}

func New() interfaces.IPie {
	return &ApplePie{
		Pie: pie.Pie{
			Size: 15,
			Form: pie.PieFormNameSquare,
			Fill: pie.FillPumpkin,
		},
	}
}
