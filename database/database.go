package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tumeraltunbas/go-blog/constants/errors"
)

var Pool *pgxpool.Pool

func Connect(connectionString string, connectionTimeout int) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(connectionTimeout)*time.Millisecond)
	defer cancel()

	pool, err := pgxpool.New(ctx, connectionString)

	if err != nil {
		log.Println("Database - Connect - New", err)
		panic(errors.ProcessFailureError)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Println("Database - Connect - Ping", err)
		panic(errors.ProcessFailureError)
	}

	Pool = pool
	return pool
}
