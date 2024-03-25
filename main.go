package main

import (
	"context"
	"log"
	"task/internal/config"
	"task/internal/service"
	"time"
)

func test(flood *service.FloodService, cntx context.Context, cfg *config.Config) {
	for i := 0; i < int(cfg.N)*10; i++ {
		var userID int64 = 1
		result, err := flood.Check(cntx, userID)
		time.Sleep(4 * time.Second)
		if err != nil {
			log.Fatal("Error checking flood control: ", err)
			return
		}

		log.Printf("Flood control for user with id: %d, check result: %t", userID, result)
	}
}

func main() {
	cntx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error while parsing file:  ")
	}
	redisServer := service.MakeRedisService(cfg.Port)
	floody := service.NewService(redisServer, cfg)

	test(floody, cntx, cfg)
	//test2(floody, cntx, cfg)
}
