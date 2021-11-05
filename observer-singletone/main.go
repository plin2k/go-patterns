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

	obs.Register("service", observerNotify, "some:event")

	srv.SomeFunction()

	obs.Unregister("service", observerNotify, "some:event")

	obs.Register("service", observerNotify, observer.AllEventsIndex)

	srv.SomeFunction()
}
