package creator

import (
	"fmt"
	"strings"
)

func ContainToken(keyValues []string) []string {

	authValues := []string{}

	for _, value := range keyValues {
		if strings.Contains(value, "Authorization") && !strings.Contains(value, "Authorization=-") {
			authValues = append(authValues, value)
		}
	}

	fmt.Printf("The length of values slice is %d.\n\n", len(authValues))
	return authValues
}

func TokenCatcher(authValues []string) []string {

	tokenSlice := []string{}

	for _, element := range authValues {
		tokenSlice = append(tokenSlice, strings.Trim(strings.TrimSpace(strings.Split(element, "Bearer")[1]), "\"}"))
	}
	return tokenSlice
}

func SimilarCount(tokenSlice []string) map[string]int {

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
