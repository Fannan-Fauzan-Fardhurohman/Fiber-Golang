package database

import (
	"fanxan/web-service-gin/internal/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.DbConfig) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		db = nil
		return
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdleConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxIdletimeConnection) * time.Second)
	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxOpenConnection))
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxIdleConnection))

	return

}
