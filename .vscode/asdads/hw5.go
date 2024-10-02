package main

import (
	"fmt"
)

func main4() {
	var n, r, i int
	var strr string
	fmt.Scan(&n)
	fmt.Scan(&strr)
	fmt.Print("Enter the number of rows: Enter print symbol:")
	i = 1
	for n > 0 {

		r = n
		for x := 0; x < r-1; x++ {

			fmt.Print(" ")

		}
		if i == 1 {
			fmt.Print(" ")

		}
		for j := 0; j < i; j++ {

			fmt.Print(strr)

		}
		i += 2
		fmt.Println()
		n--
	}
}
