package telegram

import (
	"fmt"
	"github.com/plin2k/go-patterns/abstract-factory/services/notifications/types"
)

type Html struct {
	types.Html
}

type Plain struct {
	types.Plain
}

type Telegram struct {
	Bot string
}

func (t *Telegram) Html() types.Htmler {
	return &Html{
		Html: types.Html{
			Body:  "HTML Telegram Body",
			Title: "HTML Telegram Title",
		},
	}
}

func (t *Telegram) Plain() types.Plainer {
	return &Plain{
		Plain: types.Plain{
			Message: "Plain Telegram Message",
			Header:  "Plain Telegram Header",
		},
	}
}

func (e *Telegram) Send() {
	fmt.Println("Email is send")
}
