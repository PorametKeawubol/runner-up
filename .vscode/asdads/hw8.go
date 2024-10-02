package main

import (
	"fmt"
)

func maisn() {
	var n int
	fmt.Scan(&n)
	fmt.Println("Enter number: The number is", n)
	if n%5 == 0 && n%3 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println(("Buzz"))
	}
}
