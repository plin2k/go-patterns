package main

import (
	"fmt"
	"github.com/plin2k/go-patterns/abstract-factory/services/notifications"
	"github.com/plin2k/go-patterns/abstract-factory/services/notifications/types"
)

func main() {
	telegramFactory, _ := notifications.GetNotificationsFactory(notifications.TelegramlMethod)
	emailFactory, _ := notifications.GetNotificationsFactory(notifications.EmailMethod)

	telegramPlain := telegramFactory.Plain()
	emailPlain := emailFactory.Plain()

	telegramHtml := telegramFactory.Html()
	emailHtml := emailFactory.Html()

	printPlainDetails(telegramPlain)
	printPlainDetails(emailPlain)

	printHtmlDetails(telegramHtml)
	printHtmlDetails(emailHtml)
}

func printPlainDetails(p types.Plainer) {
	fmt.Printf("Header: %s", p.GetHeader())
	fmt.Println()
	fmt.Printf("Message: %s", p.GetMessage())
	fmt.Println()
}

func printHtmlDetails(h types.Htmler) {
	fmt.Printf("Title: %s", h.GetTitle())
	fmt.Println()
	fmt.Printf("Body: %s", h.GetBody())
	fmt.Println()
}
