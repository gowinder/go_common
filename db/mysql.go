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

/**
检查表是否存在，不存在就建表
 */
func (self *MysqlClient) CheckToCreateTable(tableName string, createSql string) error{
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := self.Client.Query(sql)
	if err != nil{
		println("checkToCreateMysqlTable check table ", sql, "faield:", err)
		return err
	}

	cols, _ := rows.Columns()
	if cols == nil{
		_, err := self.Client.Exec(createSql)
		if err != nil{
			println("checkToCreateMysqlTable create table ", createSql, "faield:", err)
			return err
		}
	}

	return nil
}