package email

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

type Email struct {
	Smtp string
}

func (e *Email) Html() types.Htmler {
	return &Html{
		Html: types.Html{
			Body:  "HTML Email Body",
			Title: "HTML Email Title",
		},
	}
}

func (e *Email) Plain() types.Plainer {
	return &Plain{
		Plain: types.Plain{
			Message: "Plain Email Message",
			Header:  "Plain Email Header",
		},
	}
}

func (e *Email) Send() {
	fmt.Println("Email is send")
}
