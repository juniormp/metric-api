package infrastructure_test

import (
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestReturnsARedisClient(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_SERVER")})
	redisAdapter := repository.RedisAdapter{client}

	response, err := redisAdapter.GetRedisClient()

	assert.Empty(t, err)
	assert.Equal(t, client, response, "returns a redis client when test connection is sucess")
}

func TestReturnsAnException(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: "fake"})
	redisAdapter := repository.RedisAdapter{client}

	_, err := redisAdapter.GetRedisClient()

	assert.NotEmpty(t, err, "returns an exception when test connection failed.")
}
