package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
)

func main() {
	tokenString := ""
	claims := jwt.MapClaims{}

	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenString), nil
	})

	for key, val := range claims {
		fmt.Printf("Key: %v , value: %v\n", key, val)
	}

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

	runes := []rune(val[1])
	safeSubstring := string(runes[1:3])

	fmt.Println(safeSubstring)
}
