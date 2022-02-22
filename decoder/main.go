package main

import (
	"decoder/creator"
	"decoder/server"
	"encoding/json"
	"flag"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/dariubs/percent"
	"github.com/dgrijalva/jwt-go"
)

type dataMapSlice []map[string]interface{}
type dataMapToken []map[string]interface{}

func main() {
	claims := jwt.MapClaims{}

	redisIp := flag.String("a", "", "IP or address of redis server")
	redisPort := flag.Int("p", 6379, "redis port , the default is 6379")
	redisKey := flag.String("k", "test", "key is redis database")
	flag.Parse()

	client := server.RedisConnection(*redisIp, *redisPort)
	keyValues := server.GetValues(client, *redisKey)
	authValues := creator.ContainToken(keyValues)
	tokenSlice := creator.TokenCatcher(authValues)
	countOfToken := creator.SimilarCount(tokenSlice)
	duration := time.Duration(1) * time.Second

	var data dataMapSlice
	var token dataMapToken
	var perCent float64
	for tokenString, count := range countOfToken {
		dataMap := make(map[string]interface{})
		dataToken := make(map[string]interface{})
		jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenString), nil
		})
		jsonString, err := json.Marshal(claims)
		if err != nil {
			fmt.Println(err)
		}
		countPercentage := percent.PercentOf(count, len(tokenSlice))
		err = json.Unmarshal(jsonString, &dataMap)
		dataMap["count"] = count
		dataMap["token"] = tokenString
		dataMap["percentage"] = countPercentage
		data = append(data, dataMap)
		fmt.Println()

		dataToken["id"] = dataMap["id"]
		dataToken["token"] = tokenString
		token = append(token, dataToken)
		fmt.Println()

		keys := make([]string, 0, len(dataMap))
		for key := range dataMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			upperCaseKey := fmt.Sprintf(strings.Title(key))
			fmt.Printf("\r%v: %v\n", upperCaseKey, dataMap[key])
		}
		perCent = perCent + countPercentage
	}
	fmt.Println(perCent)
	fmt.Println("******************************************")
	fmt.Println(data)
	fmt.Println("******************************************")
	fmt.Println(token)
	fmt.Println("******************************************")
	time.Sleep(duration)
}
