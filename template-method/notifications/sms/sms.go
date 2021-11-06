package sms

import (
	"fmt"
	"github.com/plin2k/go-patterns/template-method/notifications"
)

type SMS struct {
	notifications.Notification
	receiver string
}

func (s *SMS) SetReceiver(receiver string) {
	s.receiver = receiver
	fmt.Printf("SMS: set receiver %s\n", receiver)
}

func (s *SMS) GetMessage(login string) string {
	return "SMS Message for login: " + login
}

func (s *SMS) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms to %s message : %s\n", s.receiver, message)
	return nil
}
