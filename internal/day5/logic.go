package day5

import (
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
