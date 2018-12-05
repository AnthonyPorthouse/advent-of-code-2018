package main

import (
	"advent-of-code-2018/internal/utils"
	"fmt"
)

func main() {
	data := utils.GetInput("../input")

	var twos uint
	var threes uint

	for i, line := range data {
		fmt.Printf("Parsing Line %d: %s\n", i, line)

		if hasNIdenticalRunes(2, line) {
			fmt.Println("Found a double Value")
			twos++
		}

		if hasNIdenticalRunes(3, line) {
			fmt.Println("Found a triple Value")
			threes++
		}
	}

	fmt.Printf("Twos: %d\n", twos)
	fmt.Printf("Threes: %d\n", threes)

	fmt.Printf("Checksum: %d\n", twos * threes)
}

func hasNIdenticalRunes(n uint, line string) bool {
	runes := []rune(line)
	totals := make(map[rune]uint)

	for _, r := range runes {
		totals[r]++
	}

	for _, total := range totals {
		if total == n {
			return true
		}
	}

	return false
}