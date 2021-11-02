package user_service

import (
	"fmt"
	"github.com/plin2k/go-patterns/adapter/grpc"
)

type GrpcAdapter struct {
	Grpc *grpc.Grpc
}

func (g *GrpcAdapter) InsertUserIntoMysql() {
	fmt.Println("User Send Into Grpc Adapter")
	g.Grpc.SendUserIntoGrpcService()
}
