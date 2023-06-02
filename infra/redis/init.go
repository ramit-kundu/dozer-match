package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func Initialize() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Redis:", pong)
	return client
}

//Design Decision: not creating global variables cause they create trouble during api tests.
