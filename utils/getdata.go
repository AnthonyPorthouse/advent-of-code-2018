package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func GetInput(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF || line == "" {
			break
		}

		if err != nil {
			panic(err)
		}

		lines = append(lines, strings.TrimSpace(line))
	}

	return
}
