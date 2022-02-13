package main

import (
	"decoder/creator"
	"decoder/server"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {

	redisIp := "192.168.1.90"
	redisPort := "6379"
	redisKey := "test"

	claims := jwt.MapClaims{}

	client := server.RedisConnection(redisIp, redisPort)
	keyValues := server.GetValues(client, redisKey)
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
