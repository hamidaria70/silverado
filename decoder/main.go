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
	values := []string{}

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

	for _, value := range val {
		if strings.Contains(value, "Authorization") && !strings.Contains(value, "Authorization=-") {
			values = append(values, value)
		}
	}

	fmt.Printf("The length of val slice is %d.\n", len(val))
	fmt.Printf("The length of values slice is %d.\n", len(values))
	for _, element := range values {
		tokenSlice = append(tokenSlice, strings.Trim(strings.TrimSpace(strings.Split(element, "Bearer")[1]), "\"}"))
	}

	dup_map := duplicateCount(tokenSlice)
	for tokenString, count := range dup_map {
		jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenString), nil
		})
		fmt.Println(tokenString)
		for key, val := range claims {
			fmt.Printf("Key: %v , value: %v\n", key, val)
		}
		fmt.Printf("key: count , value: %v\n", count)
	}

}

func duplicateCount(tokenSlice []string) map[string]int {

	tokenCount := make(map[string]int)
	for _, item := range tokenSlice {
		_, exist := tokenCount[item]

		if exist {
			tokenCount[item] += 1
		} else {
			tokenCount[item] = 1
		}
	}
	return tokenCount
}
