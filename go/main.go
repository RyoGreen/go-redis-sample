package main

import (
	"go-redis/config"
	"go-redis/infrastructure/postgres"
	"go-redis/infrastructure/redis"

	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
		return
	}

	if err := postgres.Connect(cfg); err != nil {
		log.Println(err)
		return
	}
	defer postgres.Close()
	redis.Connect(cfg)

}
