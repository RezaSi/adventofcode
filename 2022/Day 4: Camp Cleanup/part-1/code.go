package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Period struct {
	start int
	end   int
}

func NewPeriod(start, end int) Period {
	return Period{start: start, end: end}
}

func (p Period) Contains(other Period) bool {
	return p.start >= other.start && p.end <= other.end
}

func main() {
	bytesRead, _ := ioutil.ReadFile("input.txt")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	var result int

	for _, line := range lines {
		elves := strings.Split(line, ",")
		var periods []Period
		for _, elf := range elves {
			points := strings.Split(elf, "-")
			start, _ := strconv.Atoi(points[0])
			end, _ := strconv.Atoi(points[1])
			periods = append(periods, NewPeriod(start, end))
		}

		if periods[0].Contains(periods[1]) || periods[1].Contains(periods[0]) {
			result += 1
		}
	}

	fmt.Println(result)
}
