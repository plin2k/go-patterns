package service

import (
	"sync"
)

type Service struct {
	Count int
}

var (
	singleInstance *Service
	once           sync.Once
)

func GetInstance() *Service {
	once.Do(func() {
		singleInstance = &Service{}
	})
	return singleInstance
}

func (s *Service) Greeter() string {
	s.Count++
	return "Hello world!"
}
