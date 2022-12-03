package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytesRead, _ := ioutil.ReadFile("input.txt")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")

	var calories []int
	var curColery int
	for _, line := range lines {
		if line == "" {
			calories = append(calories, curColery)
			curColery = 0
		} else {
			val, _ := strconv.Atoi(line)
			curColery += val
		}
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	fmt.Println(calories[0] + calories[1] + calories[2])
}
