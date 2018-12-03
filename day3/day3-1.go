package main

import (
	"advent-of-code-2018/utils"
	"fmt"
	"regexp"
	"strconv"
)

var G Grid
var ClaimRegex *regexp.Regexp

func init() {
	G = make(Grid)
	ClaimRegex = regexp.MustCompile("^#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)$")
}

func main() {
	data := utils.GetInput("./input")

	for _, entry := range data {
		fmt.Printf("Parse: '%s'\n", entry)
		c := NewClaim(entry)
		c.MarkGrid()
	}

	fmt.Println(G.String())

	fmt.Printf("Number of Cells Over Claimed: %d\n", G.GetOverclaimed())
}

type Grid map[int]GridColumn
type GridColumn map[int]int

const (
	Unclaimed = iota
	Claimed
	OverClaimed
)

type Claim struct {
	ID int
	X int
	Y int
	Width int
	Height int
}

func NewClaim(data string) Claim {
	strings := ClaimRegex.FindStringSubmatch(data)

	ID, _ := strconv.Atoi(strings[1])
	X, _ := strconv.Atoi(strings[2])
	Y, _ := strconv.Atoi(strings[3])
	Width, _ := strconv.Atoi(strings[4])
	Height, _ := strconv.Atoi(strings[5])

	c := Claim{
		ID,
		X,
		Y,
		Width,
		Height,
	}

	fmt.Printf("Parsed: %+v\n", c)
	return c
}

func (c *Claim) MarkGrid() {
	xMax := c.X + c.Width
	yMax := c.Y + c.Height

	fmt.Printf("Going from %d,%d to %d,%d\n", c.X, c.Y, xMax, yMax)

	for x := c.X; x < xMax; x++ {
		if G[x] == nil {
			G[x] = make(GridColumn)
		}

		for y := c.Y; y < yMax; y++ {
			//fmt.Printf("Marking %d,%d\n", x, y)

			switch G[x][y] {
			case Unclaimed:
				//fmt.Printf("%d,%d is Unclaimed, Claiming\n", x, y)
				G[x][y] = Claimed
				break
			case Claimed:
				//fmt.Printf("%d,%d was Claimed, now OverClaimed\n", x, y)
				G[x][y] = OverClaimed
				break
			default:
				//fmt.Printf("%d,%d is OverClaimed\n", x, y)
				break
			}
		}
	}
}

func (g *Grid) String() string {
	out := ""

	for _, row := range G {
		for _, col := range row {
			switch col {
			case Unclaimed:
				out += "."
				break
			case Claimed:
				out += "#"
				break
			case OverClaimed:
				out += "X"
				break
			}
		}

		out += "\n"
	}

	return out
}

func (g *Grid) GetOverclaimed() int {
	overClaimed := 0

	for _, row := range G {
		for _, col := range row {
			switch col {
			case OverClaimed:
				overClaimed++
				break
			}
		}
	}

	return overClaimed
}