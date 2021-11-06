package service

import (
	"fmt"
	"github.com/plin2k/go-patterns/observer-singletone/observer"
)

const (
	ServiceName = "SomeService"
)

type Service struct {
	Name             string
	observerRegistry *observer.Observe
}

func New(name string, observer *observer.Observe) *Service {
	return &Service{
		Name:             name,
		observerRegistry: observer,
	}
}
func (s *Service) SomeFunction() {
	fmt.Println("Some Name Assigned")
	s.Name = "Some Name"
	s.observerRegistry.Notify(ServiceName, "some:event", s)
}
