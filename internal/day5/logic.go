package day5

import (
	"math"
	"unicode"
)

func ReduceString(chars []rune) []rune {
	for i := 0; i < len(chars); i++ {
		if i+1 == len(chars) {
			continue
		}

		c := chars[i]
		c2 := chars[i+1]

		if unicode.ToLower(c) == unicode.ToLower(c2) && c != c2 {
			var nextInput []rune

			if i > 0 {
				nextInput = append(nextInput, chars[:i]...)
			}

			nextInput = append(nextInput, chars[i+2:]...)

			return ReduceString(nextInput)
		}
	}

	return chars
}

func removeUnit(char rune, input []rune) []rune {
	var out []rune

	var startRange int

	for i := 0; i < len(input); i++ {
		if unicode.ToLower(input[i]) == char {
			out = append(out, input[startRange:i]...)
			startRange = i + 1
			continue
		}
	}

	out = append(out, input[startRange:]...)

	return out
}

func FindSmallest(input []rune) int {
	shortest := math.MaxInt64

	for i := 'a'; i <= 'z'; i++ {
		chars := removeUnit(i, input)

		l := len(ReduceString(chars))
		if l < shortest {
			shortest = l
		}
	}

	return shortest
}