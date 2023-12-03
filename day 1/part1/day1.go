package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "day 1/input.in"
	sum := 0

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			numbers := []rune{}
			for _, letter := range word {
				if letter >= '0' && letter <= '9' {
					numbers = append(numbers, letter)
				}
			}
			if len(numbers) > 0 {
				extNum, err := strconv.Atoi(string(numbers[0]) + string(numbers[len(numbers)-1]))
				if err != nil {
					fmt.Println("Error converting number:", err)
					return
				}
				sum += extNum
			}
		}
	}

	fmt.Println(sum)
}
