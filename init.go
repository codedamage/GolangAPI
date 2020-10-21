package main

import (
	redisCache "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/vmihailenco/go-tinylfu"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func dbInit() {
	var connection = mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        os.Getenv("DB_DSN"),
	})

	gDb, err := gorm.Open(connection, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = gDb
}

func cacheInit() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})

	cache = redisCache.New(&redisCache.Options{
		Redis:      ring,
		LocalCache: tinylfu.NewSync(10000, 100000),
	})
}

func dotEnvInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
