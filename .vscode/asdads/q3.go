package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maisn() {
	var input, unit string
	fmt.Scan(&input)
	fmt.Scan(&unit)

	numStr := strings.TrimRight(input, "mcm")
	origUnit := strings.TrimLeft(input, "-0123456789.")

	num, _ := strconv.ParseFloat(numStr, 64)

	var result float64

	switch origUnit {
	case "cm":
		switch unit {
		case "m":
			result = num / 100
		case "mm":
			result = num * 10
		default:
			result = num
		}
	case "mm":
		switch unit {
		case "m":
			result = num / 1000
		case "cm":
			result = num / 10
		default:
			result = num
		}
	case "m":
		switch unit {
		case "cm":
			result = num * 100
		case "mm":
			result = num * 1000
		default:
			result = num
		}
	}

	if result == float64(int(result)) {
		fmt.Printf("Enter value in mm, cm, and m: Enter unit to convert in mm, cm, m: Value after unit conversion is %.1f%s\n", result, unit)
	} else {
		fmt.Printf("Enter value in mm, cm, and m: Enter unit to convert in mm, cm, m: Value after unit conversion is %g%s\n", result, unit)
	}
}
