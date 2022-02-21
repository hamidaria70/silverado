package main

import (
	"decoder/creator"
	"decoder/server"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
)

type dataMapSlice []map[string]interface{}

func main() {
	var client *redis.Client
	keyValues := []string{}
	authValues := []string{}
	tokenSlice := []string{}
	countOfToken := map[string]int{}
	claims := jwt.MapClaims{}
	duration := time.Duration(1) * time.Second

	redisIp := flag.String("a", "", "IP or address of redis server")
	redisPort := flag.Int("p", 6379, "redis port , the default is 6379")
	redisKey := flag.String("k", "test", "key is redis database")
	flag.Parse()
	for {
		client = server.RedisConnection(*redisIp, *redisPort)
		keyValues = server.GetValues(client, *redisKey)
		authValues = creator.ContainToken(keyValues)
		tokenSlice = creator.TokenCatcher(authValues)
		countOfToken = creator.SimilarCount(tokenSlice)

		for tokenString, count := range countOfToken {
			var data dataMapSlice
			var dataMap map[string]interface{}
			jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(tokenString), nil
			})
			jsonString, err := json.Marshal(claims)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(jsonString, &dataMap)
			dataMap["count"] = count
			dataMap["token"] = tokenString

			data = append(data, dataMap)
			fmt.Println()
			fmt.Println(data)
			fmt.Println()
			fmt.Println(dataMap)
			fmt.Println()

			//keys := make([]string, 0, len(dataMap))
			//for key := range dataMap {
			//keys = append(keys, key)
			//}
			//sort.Strings(keys)
			//for _, key := range keys {
			//upperCaseKey := fmt.Sprintf(strings.Title(key))
			//fmt.Printf("\r%v: %v\n", upperCaseKey, dataMap[key])
			//}
		}
		time.Sleep(duration)
	}
}
