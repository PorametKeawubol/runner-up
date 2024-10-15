package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func mains() {
	counts := make(map[rune]int)
	scanner := bufio.NewScanner(os.Stdin)
	inputCount := 0

	for {
		fmt.Print("Enter string: ")
		inputCount++
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "end" {
			break
		}

		for _, r := range strings.ToLower(input) {
			if unicode.IsLetter(r) {
				counts[r]++
			}
		}
	}

	keys := make([]rune, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	fmt.Println("******************************")
	fmt.Println("*     Alphabet Counting      *")
	fmt.Println("******************************")
	for _, k := range keys {
		fmt.Printf("%c %d\n", k, counts[k])
	}

	fmt.Println("******************************")
}
