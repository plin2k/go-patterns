package types

type Plainer interface {
	GetMessage() string
	GetHeader() string
}

type Plain struct {
	Message string
	Header  string
}

func (p *Plain) GetMessage() string {
	return p.Message
}

func (p *Plain) GetHeader() string {
	return p.Header
}
