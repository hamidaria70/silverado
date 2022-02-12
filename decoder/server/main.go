package server

import (
	"fmt"

	"github.com/go-redis/redis"
)

func GetValues(client *redis.Client) []string {

	keyValues, err := client.LRange("test", 0, -1).Result()
	if err != nil {
		fmt.Println("OPS!!!", err)
	}
	fmt.Printf("The length of keyValues slice is %d.\n", len(keyValues))
	return keyValues

}

func RedisConnection() *redis.Client {

	fmt.Print("Checking Redis Connection: PING --> ")

	client := redis.NewClient(&redis.Options{
		Addr: "192.168.1.90:6379",
	})

	pong, err := client.Ping().Result()
	fmt.Print(pong)

	if err != nil {
		fmt.Println(err)
	}
	return client
}
