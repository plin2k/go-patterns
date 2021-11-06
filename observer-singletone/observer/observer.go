package observer

import "sync"

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

var (
	once           sync.Once
	singleInstance *Observe
)

type Observe struct {
	observerList map[string]map[string][]Observer
}

func New() *Observe {
	once.Do(func() {
		singleInstance = &Observe{
			observerList: make(map[string]map[string][]Observer),
		}
	})
	return singleInstance
}

func (o *Observe) Register(name string, observer Observer, event string) {
	if _, ok := o.observerList[name]; !ok {
		o.observerList[name] = make(map[string][]Observer)
	}
	for _, obs := range o.observerList[name][event] {
		if observer.GetId() == obs.GetId() {
			return
		}
	}
	o.observerList[name][event] = append(o.observerList[name][event], observer)
}

func (o *Observe) Unregister(name string, observer Observer, event string) {
	if _, ok := o.observerList[name]; !ok {
		return
	}
	for i, obs := range o.observerList[name][event] {
		if obs.GetId() == observer.GetId() {
			if len(o.observerList[name][event]) > 0 {
				o.observerList[name][event] = append(o.observerList[name][event][:i], o.observerList[name][event][i+1:]...)
			}
			if len(o.observerList[name][AllEventsIndex]) > 0 {
				o.observerList[name][AllEventsIndex] = append(o.observerList[name][AllEventsIndex][:i], o.observerList[name][AllEventsIndex][i+1:]...)
			}
		}
	}
}

func (o *Observe) Notify(name string, event string, data interface{}) {
	if _, ok := o.observerList[name]; !ok {
		return
	}
	observerList := o.observerList[name][event]
	if event != AllEventsIndex {
		observerList = append(observerList, o.observerList[name][AllEventsIndex]...)
	}
	for _, obs := range observerList {
		obs.Update(data, event)
	}
}
