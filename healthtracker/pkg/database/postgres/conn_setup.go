package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func NewConn(cfg Config) (*sql.DB, error) {
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}
	if cfg.Port == 0 {
		cfg.Port = 5431
	}
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "open sql conn")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "postgres ping error : (%v)")
	}
	return db, nil
}
