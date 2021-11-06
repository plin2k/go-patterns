package theme

import (
	"fmt"
	"github.com/plin2k/go-patterns/composite/interfaces"
)

type Theme struct {
	components []interfaces.Component
	Name       string
}

func (t *Theme) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in theme %s\n", keyword, t.Name)
	for _, composite := range t.components {
		composite.Search(keyword)
	}
}

func (t *Theme) Add(c interfaces.Component) {
	t.components = append(t.components, c)
}
