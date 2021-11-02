package mysql

import "fmt"

type Database interface {
	InsertUserIntoMysql()
}

type Mysql struct {
}

func (m *Mysql) InsertUserIntoMysql() {
	fmt.Println("Success insert user into mysql database")
}
