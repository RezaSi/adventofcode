package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func intersect(slices ...string) rune {
	for _, char := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		notFoundFlag := false
		for i := range slices {
			if !strings.Contains(slices[i], string(char)) {
				notFoundFlag = true
				break
			}
		}
		if !notFoundFlag {
			return char
		}
	}

	return '0'
}

func getVal(r rune) int {
	if unicode.IsUpper(r) {
		return 26 + int(r-rune('A')+1)
	} else {
		return int(r - rune('a') + 1)
	}
}

func main() {
	bytesRead, _ := ioutil.ReadFile("input.txt")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	var total int
	for current := 0; current < len(lines); current += 3 {
		first, second, third := string(lines[current]), string(lines[current+1]), string(lines[current+2])
		common := intersect(first, second, third)

		total += getVal(common)
	}

	fmt.Println(total)
}
