package model

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Database struct {
	postgres *pgx.Conn
	schema   string
}

func newDatabase(uri, schema string) (*Database, error) {
	pg, err := pgx.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}
	return &Database{
		postgres: pg,
		schema:   schema,
	}, nil
}

func (db *Database) getTableName(tableName string) string {
	return db.schema + `.` + tableName
}
