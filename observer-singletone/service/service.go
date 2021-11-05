package service

import (
	"fmt"
	"github.com/plin2k/go-patterns/observer-singletone/observer"
)

type Service struct {
	Name     string
	observer *observer.Observe
}

func New(name string, observer *observer.Observe) *Service {
	return &Service{
		Name:     name,
		observer: observer,
	}
}
func (s *Service) SomeFunction() {
	fmt.Println("Some Name Assigned")
	s.Name = "Some Name"
	s.observer.Notify("service", "some:event", s)
}
