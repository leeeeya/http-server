package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var DB *sql.DB

func ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err := DB.PingContext(ctx)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

func InitDB(connStr string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	DB = db
	ping(ctx)
}

func Close() {
	DB.Close()
	DB = nil
}
