package main

import (
	"github.com/plin2k/go-patterns/adapter/grpc"
	"github.com/plin2k/go-patterns/adapter/mysql"
	user_service "github.com/plin2k/go-patterns/adapter/user-service"
)

func main() {

	client := &user_service.UserService{}
	db := &mysql.Mysql{}

	client.SendUserIntoMysqlDatabase(db)

	grpcSender := &grpc.Grpc{}
	grpcSenderAdapter := &user_service.GrpcAdapter{
		Grpc: grpcSender,
	}

	client.SendUserIntoMysqlDatabase(grpcSenderAdapter)
}
