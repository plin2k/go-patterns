package email

import (
	"fmt"
	"github.com/plin2k/go-patterns/template-method/notifications"
)

type Email struct {
	notifications.Notification
	receiver string
}

func (e *Email) SetReceiver(receiver string) {
	e.receiver = receiver
	fmt.Printf("EMAIL: set receiver %s\n", receiver)
}

func (e *Email) GetMessage(login string) string {
	return "EMAIL Message for login: " + login
}

func (e *Email) SendNotification(message string) error {
	fmt.Printf("EMAIL: sending sms to %s message : %s\n", e.receiver, message)
	return nil
}
