package topic

import (
	"fmt"
	"github.com/plin2k/go-patterns/prototype/prototype"
)

type Topic struct {
	Name string
}

func (t *Topic) Print(indentation string) {
	fmt.Println(indentation + t.Name)
}

func (t *Topic) Clone() prototype.Prototype {
	return &Topic{Name: t.Name + "_clone"}
}
