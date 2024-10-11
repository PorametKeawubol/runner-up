package main

import (
	"fmt"
	"math"
)

func min(M [3][2]int) int {
	minValue := math.MaxInt
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if M[i][j] < minValue {
				minValue = M[i][j]
			}
		}
	}
	return minValue
}

func max(M [3][2]int) int {
	maxValue := math.MinInt
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if M[i][j] > maxValue {
				maxValue = M[i][j]
			}
		}
	}
	return maxValue
}
func inputMatrix(M *[3][2]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&M[i][j])
		}
	}
}
func mainsss() {
	var M [3][2]int
	inputMatrix(&M)
	fmt.Println("M[1][1]: M[1][2]: M[2][1]: M[2][2]: M[3][1]: M[3][2]: ")

	fmt.Println("Matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d\t", M[i][j])
		}
		fmt.Println()
	}
	minValue := min(M)
	maxValue := max(M)
	fmt.Printf("\nMin = %d\n", minValue)
	fmt.Printf("Max = %d\n", maxValue)
}
