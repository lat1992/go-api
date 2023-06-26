package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/exp/slog"
)

type Database struct {
	postgres *pgx.Conn
	schema   string
}

func NewDatabase(host, port, database, user, password, schema string) (*Database, error) {
	pg, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s/%s?user=%s&password=%s", host, port, database, user, password))
	if err != nil {
		return nil, fmt.Errorf("NewDatabase: %w", err)
	}
	return &Database{
		postgres: pg,
		schema:   schema,
	}, nil
}

func (db *Database) getTableName(tableName string) string {
	return db.schema + `.` + tableName
}

func (db *Database) Close() {
	if err := db.postgres.Close(context.Background()); err != nil {
		slog.Error("Database closing error", err)
	}
}

func (db *Database) BeginTransaction() (pgx.Tx, error) {
	tx, err := db.postgres.Begin(context.Background())
	if err != nil {
		return nil, fmt.Errorf("BeginTransaction: %w", err)
	}
	return tx, nil
}

func (db *Database) CommitTransaction(tx pgx.Tx) error {
	return tx.Commit(context.Background())
}

func (db *Database) RollbackTransaction(tx pgx.Tx) error {
	return tx.Rollback(context.Background())
}
