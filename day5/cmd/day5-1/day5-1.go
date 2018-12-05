package main

import (
	"advent-of-code-2018/day5"
	"advent-of-code-2018/utils"
	"fmt"
)

func main() {
	input := []rune(utils.GetInput("../input")[0])

	remaining := day5.ReduceString(input)

	fmt.Printf("Remaining Units: %d\n", len(remaining))
}
