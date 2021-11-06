package main

import (
	"fmt"
	"github.com/plin2k/go-patterns/prototype/course"
	"github.com/plin2k/go-patterns/prototype/prototype"
	"github.com/plin2k/go-patterns/prototype/theme"
	"github.com/plin2k/go-patterns/prototype/topic"
)

func main() {

	topic1 := &topic.Topic{
		Name: "Hello World",
	}
	topic2 := &topic.Topic{
		Name: "Channels",
	}
	topic3 := &topic.Topic{
		Name: "HTTP",
	}
	topic4 := &topic.Topic{
		Name: "TCP",
	}

	theme1 := &theme.Theme{
		Name:     "First Program",
		Children: []prototype.Prototype{topic1},
	}
	theme2 := &theme.Theme{
		Name:     "Goroutines",
		Children: []prototype.Prototype{topic2},
	}
	theme3 := &theme.Theme{
		Name:     "NET",
		Children: []prototype.Prototype{topic3, topic4},
	}

	course1 := &course.Course{
		Name:     "Golang Junior",
		Children: []prototype.Prototype{theme1, theme2, theme3},
	}

	fmt.Println("\nPrinting hierarchy for cource1")
	course1.Print("  ")

	cloneCource := course1.Clone()
	fmt.Println("\nPrinting hierarchy for clone cource1")
	cloneCource.Print("  ")
}
