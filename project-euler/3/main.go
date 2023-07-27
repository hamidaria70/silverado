package main

import "fmt"

func main() {
	n := 600851475143
	count := 0
	var primeFactors []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			for j := 2; j < i/2; j++ {
				if i%j == 0 {
					count++
					break
				}
			}
			if count == 0 && i != 1 {
				primeFactors = append(primeFactors, i)
				fmt.Println(primeFactors[len(primeFactors)-1])
			}
		}
	}
}
