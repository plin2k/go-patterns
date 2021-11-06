package prototype

type Prototype interface {
	Print(string)
	Clone() Prototype
}
