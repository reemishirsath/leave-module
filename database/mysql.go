package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//MySQLClientConn defines GORM instance for Connection.
type MySQLClientConn struct {
	clientConn *gorm.DB
}

//NewMysqlConnection for client connection
func NewMysqlConnection(connectionStr string) *gorm.DB {

	//Open Connection for GORM instance
	client, err := gorm.Open("mysql", connectionStr)

	if err != nil {
		fmt.Println("Error in Create client connection", err)
		panic("Error In Create Client Connection")
	}
	return client

}
