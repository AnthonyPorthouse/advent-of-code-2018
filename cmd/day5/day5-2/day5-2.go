package main

import (
	"advent-of-code-2018/internal/day5"
	"advent-of-code-2018/internal/utils"
	"fmt"
)

func main() {
	input := []rune(utils.GetInput("../input")[0])

	remaining := day5.FindSmallest(input)

	fmt.Printf("Smallest Possible Units: %d\n", remaining)
}
