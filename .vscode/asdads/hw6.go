package main

import (
	"fmt"
)

func mainss() {
	var n, i, x, a int
	var strr string
	fmt.Scan(&n)
	fmt.Scan(&strr)
	fmt.Print("Enter the number of rows: Enter print symbol: ")
	x = n
	i = 1
	for n > 0 {
		i += 2
		n--
	}
	a = 0
	for x > 0 {
		for z := 0; z < a; z++ {
			fmt.Print(" ")
		}
		for j := 0; j < i-2; j++ {
			fmt.Print(strr)
		}
		fmt.Println()
		a++
		i -= 2
		x--
	}

}
