/*
 * File: X:\go-api\model\Model.go
 * Created At: 2019-09-18 15:42:05
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-21 19:15:54
 * Modified By: Mauhoi WU
 */

package model

import (
	"github.com/go-pg/pg"
	"../configuration"
)

var db	*pg.DB

func OpenDatabase() {
	config := configuration.GetDatabase()
	db = pg.Connect(&pg.Options{
		Addr: config["address"] +":"+ config["port"],
		User: config["username"],
		Password: config["password"],
		Database: config["database"],
	})
}

func CloseDatabase() {
	db.Close()
}
