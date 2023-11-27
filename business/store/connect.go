package store

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net"
	"net/url"
	"time"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func Connect(cfg Config, sslRequired bool) (*sqlx.DB, error) {
	sslMode := "require"
	if !sslRequired {
		sslMode = "disable"
	}

	connStr := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.Username, cfg.Password),
		Host:     net.JoinHostPort(cfg.Host, cfg.Port),
		Path:     cfg.Database,
		RawQuery: fmt.Sprintf("sslmode=%stimezone=Europe/Istanbul", sslMode),
	}

	conn, err := sqlx.Connect("potgres", connStr.String())
	if err != nil {
		return nil, err
	}

	if err := Ping(conn); err != nil {
		return nil, err
	}

	return conn, nil
}

func Ping(conn *sqlx.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := conn.QueryContext(ctx, "SELECT now()")

	return err
}
