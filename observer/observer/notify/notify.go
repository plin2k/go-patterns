package notify

import (
	"fmt"
	"github.com/plin2k/go-patterns/observer/service"
)

type Notify struct {
	Id string
}

func New() *Notify {
	return &Notify{
		Id: "notify",
	}
}

func (n *Notify) Update(srv interface{}, event string) {
	fmt.Printf("Sending notify to observer '%s' for item '%s' from event '%s'\n", n.Id, srv.(*service.Service).Name, event)
}

func (n *Notify) GetId() string {
	return n.Id
}
