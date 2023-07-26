package main

import "fmt"

func main() {
	n := 252050
	m := 10
	var okSlice []string
	for i := 1; i <= m; i++ {
		if n%i == 0 && len(okSlice) < m {
			okSlice = append(okSlice, "ok")
			if len(okSlice) == m {
				fmt.Println(n, "is the samllest number")
			}
		} else {
			okSlice = nil
		}
	}
}