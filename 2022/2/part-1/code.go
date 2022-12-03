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

func myMove(oppMove rune, result rune) rune {
	if result == 'Z' {
		if oppMove == 'A' {
			return 'Y'
		}
		if oppMove == 'B' {
			return 'Z'
		}

		return 'X'
	} else if result == 'X' {
		if oppMove == 'A' {
			return 'Z'
		}
		if oppMove == 'B' {
			return 'X'
		}

		return 'Y'
	} else {
		if oppMove == 'A' {
			return 'X'
		}
		if oppMove == 'B' {
			return 'Y'
		}

		return 'Z'
	}

	return oppMove
}

func main() {
	bytesRead, _ := ioutil.ReadFile("input.txt")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	var score int
	for _, line := range lines {
		first, res := rune(line[0]), rune(line[2])
		second := myMove(first, res)
		score += win(first, second) + scoreMove(second)
		// fmt.Println(score)
	}

	fmt.Println(score)
}
