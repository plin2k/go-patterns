package theme

import (
	"fmt"
	"github.com/plin2k/go-patterns/prototype/prototype"
)

type Theme struct {
	Children []prototype.Prototype
	Name     string
}

func (t *Theme) Print(indentation string) {
	fmt.Println(indentation + t.Name)
	for _, i := range t.Children {
		i.Print(indentation + indentation)
	}
}

func (t *Theme) Clone() prototype.Prototype {
	cloneFolder := &Theme{Name: t.Name + "_clone"}
	var tempChildren []prototype.Prototype
	for _, i := range t.Children {
		tempChildren = append(tempChildren, i.Clone())
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
