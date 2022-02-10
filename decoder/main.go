package main

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
)

func main() {
	claims := jwt.MapClaims{}
	tokenSlice := []string{}

	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "192.168.1.90:6379",
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	val, err := client.LRange("test", 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The length of val slice is %d.\n", len(val))
	for _, element := range val {
		tokenSlice = append(tokenSlice, strings.Trim(strings.TrimSpace(strings.Split(element, "Bearer")[1]), "\"}"))
	}

	for _, token := range tokenSlice {
		jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(token), nil
		})
	}

	for key, val := range claims {
		fmt.Printf("Key: %v , value: %v\n", key, val)
	}

}
