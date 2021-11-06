package notifications

type Method interface {
	SetReceiver(string)
	GetMessage(string) string
	SendNotification(string) error
}

type Notification struct {
	Method
}

func (n *Notification) Send(receiver string, login string) error {
	n.Method.SetReceiver(receiver)
	message := n.Method.GetMessage(login)
	err := n.Method.SendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
