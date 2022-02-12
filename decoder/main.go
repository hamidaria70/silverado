package main

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
)

func main() {

	claims := jwt.MapClaims{}

	client := redisConnection()
	keyValues := getValues(client)
	authValues := containToken(keyValues)
	tokenSlice := tokenCatcher(authValues)
	countOfToken := duplicateCount(tokenSlice)

	for tokenString, count := range countOfToken {
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

	countOfToken := make(map[string]int)
	for _, item := range tokenSlice {
		_, exist := countOfToken[item]

		if exist {
			countOfToken[item] += 1
		} else {
			countOfToken[item] = 1
		}
	}
	return countOfToken
}

func redisConnection() *redis.Client {

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

func getValues(client *redis.Client) []string {

	keyValues, err := client.LRange("test", 0, -1).Result()
	if err != nil {
		fmt.Println("OPS!!!", err)
	}
	fmt.Printf("The length of keyValues slice is %d.\n", len(keyValues))
	return keyValues

}

func containToken(keyValues []string) []string {

	authValues := []string{}

	for _, value := range keyValues {
		if strings.Contains(value, "Authorization") && !strings.Contains(value, "Authorization=-") {
			authValues = append(authValues, value)
		}
	}

	fmt.Printf("The length of values slice is %d.\n", len(authValues))
	return authValues
}

func tokenCatcher(authValues []string) []string {

	tokenSlice := []string{}

	for _, element := range authValues {
		tokenSlice = append(tokenSlice, strings.Trim(strings.TrimSpace(strings.Split(element, "Bearer")[1]), "\"}"))
	}
	return tokenSlice
}
