package server

import (
	"decoder/errors"
	"fmt"

	"github.com/go-redis/redis"
)

func GetValues(client *redis.Client, redisKey string) []string {

	keyValues, err := client.LRange(redisKey, 0, -1).Result()
	errors.ConnectionError(err)
	return keyValues

}

func RedisConnection(redisIp string, redisPort int) *redis.Client {
	
	redisAddress := fmt.Sprintf("%v:%v", redisIp, redisPort)
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	pong, err := client.Ping().Result()
	if pong != "pong" {
	errors.ConnectionError(err)
	}

	return client
}
