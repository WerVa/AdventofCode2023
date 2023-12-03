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
	possibleCubes := map[string]int{"red": 12, "blue": 14, "green": 13}

	for _, line := range lines {
		game := regexp.MustCompile("; |, |: | |\n").Split(line, -1)
		for i := 3; i < len(game); i += 2 {
			value, err := strconv.Atoi(game[i-1])
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return
			}

			if value > possibleCubes[game[i]] {
				sum += atoi(game[1])
				break
			}
		}
	}

	totalSum := int(len(lines) * (len(lines) + 1) / 2)
	result := totalSum - sum

	fmt.Println(result)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting to integer:", err)
	}
	return i
}
