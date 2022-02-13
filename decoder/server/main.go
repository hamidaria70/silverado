package server

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func GetValues(client *redis.Client, redisKey string) []string {

	keyValues, err := client.LRange(redisKey, 0, -1).Result()
	if err != nil {
		fmt.Println("OPS!!!", err)
	}
	fmt.Printf("\nThe length of keyValues slice is %d.", len(keyValues))
	return keyValues

}

func RedisConnection(redisIp string, redisPort int) *redis.Client {

	redisAddress := fmt.Sprintf("%v:%v", redisIp, redisPort)

	fmt.Print("Checking Redis Connection: PING --> ")

	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	pong, err := client.Ping().Result()
	fmt.Print(pong)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return client
}
