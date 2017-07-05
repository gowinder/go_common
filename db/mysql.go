//  gowinder@hotmail.com 2017/7/5 9:27
package db

import (
 	"database/sql"
 _ "github.com/go-sql-driver/mysql"
	"fmt"
)

type MysqlClient struct{
	Client	*sql.DB
	Addr	string
	Port	int
	User	string
	Pwd		string
	Db		string
}

func (self *MysqlClient) Init(ping bool) error {
	con := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", self.User, self.Pwd, self.Addr, self.Port, self.Db)

	var err error
	self.Client, err = sql.Open("mysql", con)
	if err != nil{
		fmt.Println("MysqlClient.Init open db failed:", err)
		return err
	}

	err = self.Client.Ping()
	if err != nil{
		fmt.Println("MysqlClient.Init ping db failed:", err)
		return err
	}

	fmt.Println("MysqlClient.Init ping db ok")

	return err
}
