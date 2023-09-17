package pg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"rest-api-file-server/config"
	"time"
)

type DB struct {
	logger *zap.Logger
	config *config.PostgresConfig
	Db     *sql.DB
}

func Open(c *config.PostgresConfig) *sql.DB {
	db, err := sql.Open("postgres", c.DSN)
	if err != nil {
		panic(fmt.Errorf("fail to open postgres conn=%s", c.DSN))
	}
	db.SetMaxIdleConns(c.MaxIdleConn)
	db.SetMaxOpenConns(c.MaxOpenConn)
	db.SetConnMaxLifetime(c.ConnMaxLifetime)
	db.SetConnMaxIdleTime(c.ConnMaxIdleTime)
	return db
}

func NewPgDatabase(l *zap.Logger, c *config.PostgresConfig) *DB {
	db := Open(c)
	var err error
	for i := 0; i < 6; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
		l.Info("ping postgres database", zap.String("database", c.Database), zap.Int("cnt", i))
	}
	if err != nil {
		configJSON, _ := json.Marshal(c)
		l.Panic("fail to ping DB", zap.String("config", string(configJSON)), zap.Error(err))
	}

	l.Info("ping postgres database successfully", zap.String("database", c.Database))
	l.Info("set up connection to postgres")
	return &DB{logger: l, config: c, Db: db}
}
