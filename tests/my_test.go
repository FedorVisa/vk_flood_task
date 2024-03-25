package tests

import (
	"context"
	"log"
	"math/rand"
	"task/internal/config"
	"task/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setup() *service.FloodService {
	cfg := config.Config{
		Port: ":8080",
		N:    5,
		K:    10 * time.Second,
	}
	rs := service.MakeRedisService(cfg.Port)
	fs := service.NewService(rs, &cfg)
	return fs
}

func TestFloodService(t *testing.T) {
	flood := setup()
	falsecounter := 0
	cntx := context.Background()
	for i := 0; i < int(1)*10; i++ {
		var userID int64 = 1
		result, err := flood.Check(cntx, userID)
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Fatal("Error checking flood control: ", err)
			return
		}

		if result == false {
			falsecounter++
		}
		log.Printf("Flood control for user with id: %d, check result: %t", userID, result)
	}
	assert.Equal(t, 2, falsecounter)
}

func TestDefault(t *testing.T) {
	flood := setup()
	cntx := context.Background()
	for i := 0; i < int(5)*10; i++ {
		userID := rand.Int63n(5)
		result, err := flood.Check(cntx, userID)
		assert.NoError(t, err, nil)

		log.Printf("Flood control for user with id: %d, check result: %t", userID, result)
	}

}

// test that function works correctly
func TestOverflowMessages(t *testing.T) {
	flood := setup()
	bool_test := false
	cntx := context.Background()
	for i := 0; i < int(5); i++ {
		var userID int64 = 1
		result, err := flood.Check(cntx, userID)
		if err != nil {
			log.Fatal("Error checking flood control: ", err)
			return
		}
		if !result {
			bool_test = true
		}
		log.Printf("Flood control for user with id: %d, check result: %t", userID, result)
	}
	assert.True(t, bool_test)
}
