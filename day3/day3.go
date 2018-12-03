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

	var claims []Claim


	for _, entry := range data {
		fmt.Printf("Parse: '%s'\n", entry)
		c := NewClaim(entry)
		c.MarkGrid()

		claims = append(claims, c)
	}

	fmt.Println()
	fmt.Printf("Number of Cells Over Claimed: %d\n", G.GetOverclaimed())

	for _, c := range claims {
		if !G.HasOverlappedTile(c) {
			fmt.Printf("Claim #%d has no overlaps", c.ID)
			break
		}
	}
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
			switch G[x][y] {
			case Unclaimed:
				G[x][y] = Claimed
				break
			case Claimed:
				G[x][y] = OverClaimed
				break
			default:
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

func (g *Grid) HasOverlappedTile(c Claim) bool {
	xMax := c.X + c.Width
	yMax := c.Y + c.Height

	for x := c.X; x < xMax; x++ {
		for y := c.Y; y < yMax; y++ {
			switch G[x][y] {
			case OverClaimed:
				return true
			}
		}
	}

	return false
}