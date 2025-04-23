package api

import (
	"github.com/chocological13/tech-stream/user-service/config"
	"github.com/chocological13/tech-stream/user-service/db"
	"log"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	dbpool := db.ConnectDB(cfg.DatabaseUrl)
	log.Println("Connected to database")
	defer dbpool.Close()

	// Connect to redis
	rdb := db.ConnectRedis()
	log.Println("Connected to redis")
	defer rdb.Close()

	// Run migrations
	err = db.RunMigrations(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
