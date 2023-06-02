package cache

import (
	"time"

	"github.com/go-redis/redis"
	rd "github.com/kundu-ramit/dozer_match/infra/redis"
)

type Cache interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	Remove(key string)
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
	err := c.client.Set(key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c cache) Get(key string) (string, error) {
	value, err := c.client.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return value, nil
}

func (c cache) Remove(key string) {
	c.client.Del(key)
}
