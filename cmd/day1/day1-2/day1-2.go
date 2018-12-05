package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	getFirstRepeat()
}

func getFirstRepeat() {
	var total int
	totals := make(map[int]bool)

	for {
		file, err := os.Open("../input")
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			if line == "" {
				break
			}

			fmt.Printf("Read Line: %s\n", line)

			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			total += val

			if totals[total] == true {
				fmt.Printf("Found First Duplicate Total: %d\n", total)
				return
			}

			fmt.Printf("Current Total: %d\n", total)
			fmt.Println()

			totals[total] = true
		}

		file.Close()
	}
}