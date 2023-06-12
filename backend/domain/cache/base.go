package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	rd "github.com/kundu-ramit/dozer_match/infra/redis"
)

type Cache interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	Remove(key string)
	FlushAll()
}

type cache struct {
	client *redis.Client
}

func NewCache() Cache {
	return cache{
		client: rd.Initialize(),
	}
}

func (c cache) Set(key string, value string, expiration time.Duration) error {
	fmt.Println("redis : setting key" + key + "with value" + value)

	err := c.client.Set(key, value, expiration).Err()
	if err != nil {
		fmt.Println("Some error happened while setting key" + key + "with value" + value + err.Error())
		return err
	}
	return nil
}

func (c cache) Get(key string) (string, error) {
	fmt.Println("redis : getting key" + key)

	value, err := c.client.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		fmt.Println("Some error happened while getting redis key" + key + err.Error())
		return "", err
	}

	return value, nil
}

func (c cache) Remove(key string) {
	fmt.Println("redis : deleting key" + key)
	c.client.Del(key)
}

func (c cache) FlushAll() {
	fmt.Println("redis : flushing all keys")
	c.client.FlushAll()
}
