package main

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
)

func main() {
	claims := jwt.MapClaims{}

	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "192.168.1.90:6379",
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	val, err := client.LRange("test", -1, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	tokenString := strings.Trim(strings.TrimSpace(strings.Split(val[0], "Bearer")[1]), "\"}")

	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenString), nil
	})

	for key, val := range claims {
		fmt.Printf("Key: %v , value: %v\n", key, val)
	}

}
