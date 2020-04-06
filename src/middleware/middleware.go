package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/juniormp/metric-api/src/infrastructure/repository"
)

func RedisHandler(c *gin.Context) {
	redisServer := os.Getenv("REDIS_SERVER")
	client := redis.NewClient(&redis.Options{Addr: redisServer})
	redisAdapter := repository.RedisAdapter{client}

	if client, err := redisAdapter.GetRedisClient(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.Set("redis", client)
		c.Next()
		redisAdapter.CloseRedisConnection(client)
	}
}
