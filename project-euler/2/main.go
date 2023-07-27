package main

import "fmt"

func main() {
	var (
		sum int64
		a   int64
		b   int64
	)

	a, b = 0, 1
	for b < 4000000 {
		a, b = b, b+a

		if a%2 == 1 {
			sum = sum + a
		}
	}
	fmt.Println(sum)
}
