package service

import (
	"fmt"
	"github.com/plin2k/go-patterns/observer-singletone/observer"
)

type Service struct {
	Name             string
	observerRegistry *observer.Observe
	ServiceName      string
}

func New(name string, observer *observer.Observe) *Service {
	return &Service{
		Name:             name,
		ServiceName:      "service",
		observerRegistry: observer,
	}
}
func (s *Service) SomeFunction() {
	fmt.Println("Some Name Assigned")
	s.Name = "Some Name"
	s.observerRegistry.Notify(s.ServiceName, "some:event", s)
}
