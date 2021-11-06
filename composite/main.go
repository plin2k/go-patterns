package main

import (
	"github.com/plin2k/go-patterns/composite/course"
	"github.com/plin2k/go-patterns/composite/theme"
	"github.com/plin2k/go-patterns/composite/topic"
)

func main() {

	course1 := &course.Course{
		Name: "Golang Junior",
	}

	theme1 := &theme.Theme{
		Name: "First Program",
	}
	theme2 := &theme.Theme{
		Name: "Goroutines",
	}
	theme3 := &theme.Theme{
		Name: "NET",
	}

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

	theme1.Add(topic1)
	theme2.Add(topic2)
	theme3.Add(topic3)
	theme3.Add(topic4)

	course1.Add(theme1)
	course1.Add(theme2)
	course1.Add(theme3)

	course1.Search("UDP")
}
