package main

import (
	"github.com/plin2k/go-patterns/observer/observer"
	"github.com/plin2k/go-patterns/observer/observer/notify"
	"github.com/plin2k/go-patterns/observer/service"
)

func main() {

	srv := service.New("Some Name For Service")

	observerNotify := notify.New()

	srv.Register(observerNotify, "some:event")

	srv.SomeFunction()

	srv.Unregister(observerNotify, "some:event")
	srv.Register(observerNotify, observer.AllEventsIndex)

	srv.SomeFunction()

	srv.Unregister(observerNotify, observer.AllEventsIndex)
}
