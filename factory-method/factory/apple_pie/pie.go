package apple_pie

import (
	"github.com/plin2k/go-patterns/factory-method/factory/interfaces"
	"github.com/plin2k/go-patterns/factory-method/pie"
)

const Name = "apple"

type ApplePie struct {
	pie.Pie
}

func New() interfaces.IPie {
	return &ApplePie{
		Pie: pie.Pie{
			Size: 10,
			Form: pie.PieFormNameCircle,
			Fill: pie.FillApple,
		},
	}
}
