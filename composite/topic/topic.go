package topic

import "fmt"

type Topic struct {
	Name string
}

func (t *Topic) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in topic %s\n", keyword, t.Name)
}

func (t *Topic) GetName() string {
	return t.Name
}
