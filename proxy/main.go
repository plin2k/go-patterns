package main

import (
	"fmt"
	"github.com/plin2k/go-patterns/proxy/healthcheck"
	"github.com/plin2k/go-patterns/proxy/middleware/nginx"
)

func main() {
	healthChk := &healthcheck.Healthcheck{}
	nginxServer := nginx.New(healthChk)
	appStatusURL := "/app/status"
	errorURL := "/app/statuserror"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(errorURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(errorURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
