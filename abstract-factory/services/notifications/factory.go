package notifications

import (
	"fmt"
	"github.com/plin2k/go-patterns/abstract-factory/services/email"
	"github.com/plin2k/go-patterns/abstract-factory/services/notifications/types"
	"github.com/plin2k/go-patterns/abstract-factory/services/telegram"
)

const (
	EmailMethod     = "email"
	TelegramlMethod = "telegram"
)

type NotificationsFactory interface {
	Html() types.Htmler
	Plain() types.Plainer
	Send()
}

func GetNotificationsFactory(method string) (NotificationsFactory, error) {
	switch method {
	case EmailMethod:
		return &email.Email{}, nil
	case TelegramlMethod:
		return &telegram.Telegram{}, nil
	default:
		return nil, fmt.Errorf("Wrong method type passed")
	}
}
