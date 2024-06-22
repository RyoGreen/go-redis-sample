package postgres

import (
	"database/sql"
	"go-redis/config"

	"github.com/RyoGreen/pgdsn"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect(cfg *config.Config) error {
	c := &pgdsn.Config{
		User:     cfg.Database.User,
		DbName:   cfg.Database.Dbname,
		Password: cfg.Database.Password,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		SslMode:  pgdsn.Disable,
	}
	var err error
	db, err = sql.Open("postgres", c.FormatDSN())
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

func Close() error {
	return db.Close()
}
