package healthcheck

type Healthcheck struct {
}

func (h *Healthcheck) HandleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok GET"
	}

	if url == "/app/status" && method == "POST" {
		return 201, "Ok POST"
	}
	return 404, "Not Ok"
}
