package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	getTotal()
}

func getTotal() {
	file, err := os.Open("../input")
	if err != nil {
		panic(err)
	}

	var total int

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		total += val
		fmt.Printf("Read Line: %s\n", line)
		fmt.Printf("Current Total: %d\n", total)
		fmt.Println()
	}

	fmt.Printf("Final Total: %d\n", total)
}