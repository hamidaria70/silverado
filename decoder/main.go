package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
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
}
