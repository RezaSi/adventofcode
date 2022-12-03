package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func win(x, y rune) int {
	if x == 'A' && y == 'Y' {
		return 6
	}
	if x == 'B' && y == 'Z' {
		return 6
	}
	if x == 'C' && y == 'X' {
		return 6
	}

	if (x == 'A' && y == 'X') || (x == 'B' && y == 'Y') || (x == 'C' && y == 'Z') {
		return 3
	}

	return 0
}

func scoreMove(x rune) int {
	if x == 'X' {
		return 1
	} else if x == 'Y' {
		return 2
	}

	return 3
}

func main() {
	bytesRead, _ := ioutil.ReadFile("input.txt")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	var score int
	for _, line := range lines {
		first, second := rune(line[0]), rune(line[2])
		score += win(first, second) + scoreMove(second)
	}

	fmt.Println(score)
}
