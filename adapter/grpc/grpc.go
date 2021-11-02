package grpc

import "fmt"

type Grpc struct {
}

func (g *Grpc) SendUserIntoGrpcService() {
	fmt.Println("Success send user into grpc service")
}
