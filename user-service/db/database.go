package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"path/filepath"
)

func ConnectDB(connString string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	return dbpool
}

func RunMigrations(dbUrl string) error {
	db, err := goose.OpenDBWithDriver("postgres", dbUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Println(dir)
	migrationsPath := filepath.Join(dir, "migrations")

	return goose.Up(db, migrationsPath)
}

func ConnectRedis() *redis.Client {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_ADDR"))
	rdb := redis.NewClient(opt)

	return rdb
}
