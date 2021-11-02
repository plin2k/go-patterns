package service

import (
	"log"
	"sync"
)

var once sync.Once

// Service - some structure ...
type Service struct {
	Hello int
}

// singleInstance - private variable for save instance
var singleInstance *Service

// GetInstance ...
func GetInstance() *Service {
	if singleInstance == nil {
		once.Do(func() {
			singleInstance = &Service{}
		})
	}
	return singleInstance
}

// Greeter ...
func (s *Service) Greeter() {
	log.Println("Hello world!")
	s.Hello++
}
