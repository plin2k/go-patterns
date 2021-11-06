package pie

const (
	FillApple   = "Apple"
	FillPumpkin = "Pumpkin"

	PieFormNameCircle = "circle"
	PieFormNameSquare = "square"
)

type Pie struct {
	Size int
	Form string
	Fill string
}

func (g *Pie) SetSize(size int) {
	g.Size = size
}

func (g *Pie) GetSize() int {
	return g.Size
}

func (g *Pie) SetForm(form string) {
	g.Form = form
}

func (g *Pie) GetForm() string {
	return g.Form
}

func (g *Pie) SetFill(fill string) {
	g.Fill = fill
}

func (g *Pie) GetFill() string {
	return g.Fill
}
