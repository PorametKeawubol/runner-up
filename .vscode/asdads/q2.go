package main

import (
	"fmt"
)

func main() {
	var line int
	fmt.Scan(&line)

	if line <= 0 || line >= 10 {
		fmt.Println("Line: Out of Range")
	} else {
		printNumberTriangle(line)
	}
}

func printNumberTriangle(line int) {
	fmt.Print("Line: ")
	for i := 1; i <= line; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for j := i - 1; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
