package repository

import (
	"github.com/go-redis/redis"
)

type RedisAdapter struct {
	Client *redis.Client
}

func (redisAdapter RedisAdapter) GetRedisClient() (*redis.Client, error) {
	client := redisAdapter.Client
	err := checkRedisConnection(client)

	return client, err
}

func (redisAdapter RedisAdapter) CloseRedisConnection(client *redis.Client) {
	if client != nil {
		client.Close()
	}
}

func checkRedisConnection(client *redis.Client) error {
	if err := client.Ping().Err(); err != nil {
		return err
	} else {
		return nil
	}
}
