package course

import (
	"fmt"
	"github.com/plin2k/go-patterns/composite/interfaces"
)

type Course struct {
	components []interfaces.Component
	Name       string
}

func (c *Course) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in cource %s\n", keyword, c.Name)
	for _, composite := range c.components {
		composite.Search(keyword)
	}
}

func (c *Course) Add(com interfaces.Component) {
	c.components = append(c.components, com)
}
