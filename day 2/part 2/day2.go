package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filePath := "day 2/input.in"

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	sum := 0
	for _, line := range lines {
		numberCubes := map[string]int{"red": 0, "blue": 0, "green": 0}
		newLine := regexp.MustCompile("; |, |: | |\n").Split(line, -1)
		for i := 3; i < len(newLine); i += 2 {
			value, err := strconv.Atoi(newLine[i-1])
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return
			}
			numberCubes[newLine[i]] = max(numberCubes[newLine[i]], value)
		}

		product := 1
		for _, value := range numberCubes {
			product *= value
		}
		sum += product
	}

	fmt.Println(sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
