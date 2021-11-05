package observer

const (
	AllEventsIndex = "*"
)

type Observer interface {
	Update(interface{}, string)
	GetId() string
}

type Subject interface {
	Register(Observer)
	Unregister(Observer)
	Notify()
}
