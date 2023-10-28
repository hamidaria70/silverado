package main

import "fmt"

func main() {
	i := 1
	count := 0
	for n := 1; n <= 10; n++ {
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				count++
				break
			}
		}
		fmt.Println(count)
	}
}
