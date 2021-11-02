package user_service

import (
	"fmt"
	"github.com/plin2k/go-patterns/adapter/mysql"
)

type UserService struct {
}

func (c *UserService) SendUserIntoMysqlDatabase(db mysql.Database) {
	fmt.Println("User Send Into Mysql")
	db.InsertUserIntoMysql()
}
