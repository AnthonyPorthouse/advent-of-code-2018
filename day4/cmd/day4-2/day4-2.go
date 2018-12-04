package main

import (
	"advent-of-code-2018/utils"
	"advent-of-code-2018/day4"
	"sort"
)

func main() {
	input := utils.GetInput("../input")
	sort.Strings(input)

	day4.Strategy2(input)
}