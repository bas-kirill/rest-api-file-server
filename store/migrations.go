package store

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"rest-api-file-server/config"
)

type PgMigrator struct {
	logger   *zap.Logger
	pgConfig *config.PostgresConfig
}

func NewPgMigrator(logger *zap.Logger, pgConfig *config.PostgresConfig) *PgMigrator {
	return &PgMigrator{logger: logger, pgConfig: pgConfig}
}

func (m *PgMigrator) RunMigrations() {
	migration, err := migrate.New(m.pgConfig.MigrationsUrl, m.pgConfig.DSN)
	if err != nil {
		panic(fmt.Errorf("fail to apply migration: %v", err))
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(fmt.Errorf("fail to run migrate up: %v", err))
	}

	m.logger.Info("postgres migrated successfully")
}
