package db

import (
	"context"
	"fmt"
	//"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeConnectionToDB() (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		"postgres", //os.Getenv("DB_USER"),
		"postgres", //os.Getenv("DB_PASSWORD"),
		"db", //os.Getenv("DB_HOST"),
		"5432", //os.Getenv("DB_PORT"),
		"postgresdb", //os.Getenv("DB_NAME"),
	)

	var pool *pgxpool.Pool
	var err error
	
	// Try connecting to the db up to 10 times
	for i := 1; i <= 10; i++ {
		pool, err = pgxpool.New(context.Background(), dsn)
		if err == nil {
			err = pool.Ping(context.Background())
			if err == nil {
				fmt.Println("Connected to DB")
				return pool, nil
			}
		}

		fmt.Printf("DB not ready (attempt %d/10), retrying...\n", i)
		time.Sleep(2 * time.Second) // wait 2 seconds to retry
	}

	return nil, fmt.Errorf("could not connect to DB after retries. Error: %w", err)
}

