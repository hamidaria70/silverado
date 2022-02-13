package main

import (
	"decoder/creator"
	"decoder/server"
	"flag"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {

	redisIp := flag.String("a", "", "IP or address of redis server")
	redisPort := flag.Int("p", 6379, "redis port , the default is 6379")
	redisKey := flag.String("k", "test", "key is redis database")
	flag.Parse()

	claims := jwt.MapClaims{}

	client := server.RedisConnection(*redisIp, *redisPort)
	keyValues := server.GetValues(client, *redisKey)
	authValues := creator.ContainToken(keyValues)
	tokenSlice := creator.TokenCatcher(authValues)
	countOfToken := creator.SimilarCount(tokenSlice)

	for tokenString, count := range countOfToken {
		jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenString), nil
		})
		fmt.Printf("%v\n", tokenString)
		for key, val := range claims {
			fmt.Printf("Key: %v , value: %v\n", key, val)
		}
		fmt.Printf("key: count , value: %v\n\n", count)
	}

}
