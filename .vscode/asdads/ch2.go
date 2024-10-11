package main

import (
	"fmt"
)

var cache = make(map[int]int)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if val, exists := cache[n]; exists {
		return val
	}
	cache[n] = fib(n-1) + fib(n-2)
	return cache[n]
}
func maiห() {
	var n int

	fmt.Scan(&n)
	fmt.Print("Input number of Fibonacci number: ")

	fmt.Printf("0")
	for i := 1; i <= n; i++ {
		fmt.Printf(", %d", fib(i))
	}
	fmt.Println()
}
