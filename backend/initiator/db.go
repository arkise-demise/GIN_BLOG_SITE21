package initiator

import (
	"context"
	"fmt"
	"time"

	"ride-plus-v2/platform/logger"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func InitDB(url string, log logger.Logger) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to connect to database: %v", err))
	}

	idleConnTimeout := viper.GetDuration("database.idle_conn_timeout")
	if idleConnTimeout == 0 {
		idleConnTimeout = 4 * time.Minute
	}

	config.ConnConfig.Logger = log.Named("pgx")
	config.MaxConnIdleTime = idleConnTimeout
	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to connect to database: %v", err))
	}

	if _, err := conn.Exec(context.Background(), "show tables"); err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to ping database: %v", err))
	}

	return conn
}
