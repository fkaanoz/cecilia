package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net"
	"net/url"
	"time"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func Connect(cfg Config) (*redis.Client, error) {
	connStr := url.URL{
		Scheme: "redis",
		User:   url.UserPassword(cfg.Username, cfg.Password),
		Host:   net.JoinHostPort(cfg.Host, cfg.Port),
		Path:   cfg.Database,
	}

	options, err := redis.ParseURL(connStr.String())
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)
	err = Ping(client)

	return client, err
}

func Ping(client *redis.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	status := client.Ping(ctx)

	return status.Err()
}
