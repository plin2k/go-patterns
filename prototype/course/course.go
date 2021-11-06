package course

import (
	"fmt"
	"github.com/plin2k/go-patterns/prototype/prototype"
)

type Course struct {
	Children []prototype.Prototype
	Name     string
}

func (c *Course) Print(indentation string) {
	fmt.Println(indentation + c.Name)
	for _, i := range c.Children {
		i.Print(indentation + indentation)
	}
}

func (c *Course) Clone() prototype.Prototype {
	cloneFolder := &Course{Name: c.Name + "_clone"}
	var tempChildren []prototype.Prototype
	for _, i := range c.Children {
		tempChildren = append(tempChildren, i.Clone())
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
