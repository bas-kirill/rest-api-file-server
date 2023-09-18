package pg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"rest-api-file-server/config"
	"time"
)

type DB struct {
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

func NewPgDatabase(c *config.PostgresConfig) *DB {
	db := Open(c)
	configJSON, _ := json.Marshal(c)
	fmt.Printf("ping DB `%s`\n", string(configJSON))
	var err error
	for i := 0; i < 6; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Printf("ping postgres database `%s`\n", c.Database)
	}
	if err != nil {
		panic(fmt.Errorf("fail to ping `%s`", string(configJSON)))
	}

	fmt.Printf("set up connection to `%s`\n", c.Database)
	return &DB{config: c, Db: db}
}
