package types

type Htmler interface {
	GetBody() string
	GetTitle() string
}

type Html struct {
	Title string
	Body  string
}

func (h *Html) GetBody() string {
	return h.Body
}

func (h *Html) GetTitle() string {
	return h.Title
}
