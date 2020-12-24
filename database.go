package main

import "github.com/jackc/pgx/pgxpool"
import "log"
import "context"

var dbpool *pgxpool.Pool

func DBpool() *pgxpool.Pool {
	return dbpool
}

func ConnectToDatabase() {
	config, err := pgxpool.ParseConfig("postgres://postgres:tajnehaslo@localhost:5432/ideas")
	if err != nil {
		log.Fatalf("%s", err)
	}
	dbpool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
