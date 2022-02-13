package main

import (
	"decoder/creator"
	"decoder/server"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func main() {

	claims := jwt.MapClaims{}

	client, err := server.RedisConnection()
	if err != nil {
		os.Exit(1)
	}

	keyValues := server.GetValues(client)
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
