package main

import (
	"fmt"
	"github.com/plin2k/go-patterns/template-method/notifications"
	"github.com/plin2k/go-patterns/template-method/notifications/email"
	"github.com/plin2k/go-patterns/template-method/notifications/sms"
)

func main() {

	smsNotify := &sms.SMS{}
	o := notifications.Notification{
		Method: smsNotify,
	}
	err := o.Send("79271234567", "admin")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("")

	emailNotify := &email.Email{}
	o = notifications.Notification{
		Method: emailNotify,
	}
	err = o.Send("aleksandr@kalink.in", "moderator")
	if err != nil {
		fmt.Println(err)
	}

}
