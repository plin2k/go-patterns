package main

import (
	"github.com/plin2k/go-patterns/observer-singletone/observer"
	"github.com/plin2k/go-patterns/observer-singletone/observer/notify"
	"github.com/plin2k/go-patterns/observer-singletone/service"
)

func main() {
	obs := observer.New()
	srv := service.New("Some Name For Service", obs)

	observerNotify := notify.New()

	srv.SomeFunction()

	obs.Register(srv.ServiceName, observerNotify, "some:event")

	srv.SomeFunction()

	obs.Unregister(srv.ServiceName, observerNotify, "some:event")

	obs.Register(srv.ServiceName, observerNotify, observer.AllEventsIndex)

	srv.SomeFunction()
}
