package middleware

type Middleware interface {
	HandleRequest(string, string) (int, string)
}
