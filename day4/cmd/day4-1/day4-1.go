package main

import (
	"advent-of-code-2018/day4"
	"advent-of-code-2018/utils"
	"sort"
)

func main() {
	input := utils.GetInput("../input")
	sort.Strings(input)

	day4.Strategy1(input)
}


