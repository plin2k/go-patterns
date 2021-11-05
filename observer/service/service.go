package service

import (
	"fmt"
	"github.com/plin2k/go-patterns/observer/observer"
)

type Service struct {
	observerList map[string][]observer.Observer
	Name         string
}

func New(name string) *Service {
	return &Service{
		Name:         name,
		observerList: make(map[string][]observer.Observer),
	}
}
func (s *Service) SomeFunction() {
	fmt.Println("Some Name Assigned")
	s.Name = "Some Name"
	s.Notify("some:event")
}

func (s *Service) Register(o observer.Observer, event string) {
	for _, obs := range s.observerList[event] {
		if o.GetId() == obs.GetId() {
			return
		}
	}
	s.observerList[event] = append(s.observerList[event], o)
}

func (s *Service) Unregister(o observer.Observer, event string) {
	for i, obs := range s.observerList[event] {
		if obs.GetId() == o.GetId() {
			if len(s.observerList[event]) > 0 {
				s.observerList[event] = append(s.observerList[event][:i], s.observerList[event][i+1:]...)
			}
			if len(s.observerList[observer.AllEventsIndex]) > 0 {
				s.observerList[observer.AllEventsIndex] = append(s.observerList[observer.AllEventsIndex][:i], s.observerList[observer.AllEventsIndex][i+1:]...)
			}
		}
	}
}

func (s *Service) Notify(event string) {
	observerList := s.observerList[event]
	if event != observer.AllEventsIndex {
		observerList = append(observerList, s.observerList[observer.AllEventsIndex]...)
	}
	for _, obs := range observerList {
		obs.Update(s, event)
	}
}
