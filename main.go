package main

import (
	"fmt"
)

func main() {
	fmt.Println(solve(5))
}

func solve(x int) int {
	var fact func(n int) int
	fact = func(n int) int {
		if n == 1 {
			return n
		}
		return n * fact(n-1)
	}
	return fact(x)
}
