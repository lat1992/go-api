package model

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Model struct {
	database *Database
	cache    *Cache
	logger   *zap.SugaredLogger
}

func NewModel(logger *zap.SugaredLogger, conf *viper.Viper) *Model {
	db, err := newDatabase(conf.GetString("postgres.uri"), conf.GetString("postgres.schema"))
	if err != nil {
		logger.Fatalf("NewModel: newDatabase: %v", err)
	}
	logger.Infof("Database connected")
	c, err := newCache(viper.GetString("redis.uri"), viper.GetInt("redis.database"))
	if err != nil {
		logger.Fatalf("NewModel: newCache: %v", err)
	}
	logger.Infof("Cache connected")
	return &Model{
		database: db,
		cache:    c,
		logger:   logger,
	}
}

func (m *Model) Close() {
	if err := m.database.postgres.Close(context.Background()); err != nil {
		m.logger.Fatalf("Database closing error: %v", err)
	}
	if err := m.cache.redis.Close(); err != nil {
		m.logger.Fatalf("Cache closing error: %v", err)
	}
}

func (m *Model) BeginTransaction() (pgx.Tx, error) {
	tx, err := m.database.postgres.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (m *Model) CommitTransaction(tx pgx.Tx) error {
	return tx.Commit(context.Background())
}

func (m *Model) RollbackTransaction(tx pgx.Tx) error {
	return tx.Rollback(context.Background())
}
