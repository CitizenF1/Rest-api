package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// POSTGRES
func Initdb() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error load .env file")
	}
	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES"))
	if err != nil {
		log.Printf("WARNING %v", err)
	} else {
		fmt.Println("Database ready")
	}
	// connection check
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return db
}

//REDIS
func RedisConnect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // or localhost
		Password: "",           // no password set
		DB:       0,            // use default DB
	})
}
