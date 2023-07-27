package main

import "fmt"

func main() {
	n := 100
	var sum1, sum2, power, diff int

	for i := 1; i <= n; i++ {
		sum1 = sum1 + (i * i)
		sum2 = sum2 + i
	}
	power = sum2 * sum2
	diff = power - sum1
	fmt.Println(diff)

}
