package main

import (
	"advent-of-code-2018/utils"
	"fmt"
	"github.com/xrash/smetrics"
)

func main() {
	data := utils.GetInput("./input")

	var lowestPair []string

	Search:
	for i, current := range data {
		for j, compare := range data {
			if i == j {
				continue
			}

			if hasSimilarEnoughMatch(current, compare) {
				lowestPair = []string{current, compare}
				break Search
			}
		}
	}

	fmt.Printf("Closest Match: %+v\n", lowestPair)

	var result string

	lowestA := []rune(lowestPair[0])
	lowestB := []rune(lowestPair[1])

	for i, c := range lowestA {
		d := lowestB[i]

		if c == d {
			result += string(c)
		}
	}

	fmt.Printf("Result: %s\n", result)
}

func hasSimilarEnoughMatch(a string, b string) bool {
	i, _ := smetrics.Hamming(a, b)

	if i == 1 {
		return true
	}

	return false
}