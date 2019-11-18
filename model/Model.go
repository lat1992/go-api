/*
 * File: X:\enigm-mvc\model\Model.go
 * Created At: 2019-09-18 15:42:05
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-17 15:54:36
 * Modified By: Mauhoi WU
 */

package model

import (
	"github.com/go-pg/pg"
	"../configuration"
)

type Model struct {
	db	*pg.DB
}

func NewModel() *Model {
	m := &Model{}
	m.OpenDatabase()
	return m
}

func (m *Model) Destroy() {
	m.CloseDatabase()
}

func (m *Model) OpenDatabase() {
	config := configuration.GetDatabase()
	m.db = pg.Connect(&pg.Options{
		Addr: config["address"] +":"+ config["port"],
		User: config["username"],
		Password: config["password"],
		Database: config["database"],
	})
}

func (m *Model) CloseDatabase() {
	m.db.Close()
}
