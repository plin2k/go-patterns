package main

import (
	"github.com/plin2k/go-patterns/observer-singletone/observer"
	"github.com/plin2k/go-patterns/observer-singletone/observer/notify"
	"github.com/plin2k/go-patterns/observer-singletone/service"
)

func main() {
	obs := observer.GetInstance()
	srv := service.New("Some Name For Service", obs)

	observerNotify := notify.New()

	srv.SomeFunction()

	obs.Register(service.ServiceName, observerNotify, "some:event")

	srv.SomeFunction()

	obs.Unregister(service.ServiceName, observerNotify, "some:event")

	obs.Register(service.ServiceName, observerNotify, observer.AllEventsIndex)

	srv.SomeFunction()
}
