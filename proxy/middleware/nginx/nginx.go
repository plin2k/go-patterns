package nginx

import "github.com/plin2k/go-patterns/proxy/middleware"

type Nginx struct {
	Application       middleware.Middleware
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func New(mdw middleware.Middleware) *Nginx {
	return &Nginx{
		Application:       mdw,
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.Application.HandleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
