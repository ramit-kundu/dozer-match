package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Initialize() *redis.Client {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password
		DB:       0,                // Redis database number
	})

	// Ping the Redis server to check the connection
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Redis:", pong)
	return client
}
