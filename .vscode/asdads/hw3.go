package main

import (
	"fmt"
)

func maidn() {
	var n int
	fmt.Scan(&n)
	fmt.Print("Enter the number of rows: ")

	for i := 0; i < n+1; i++ {
		for j := n; j > i; j-- {
			fmt.Print("*")
		}

		fmt.Println()
	}

}
