package main

import (
	"fmt"
	"path"
	"regexp"
	"runtime"
	"strconv"

	"github.com/hsinpaohuang/aoc22/utils"
)

type Assignment struct {
	min int
	max int
}

// ex: 31-33,32-48
var exp = regexp.MustCompile(`(?P<min1>\d+)-(?P<max1>\d+),(?P<min2>\d+)-(?P<max2>\d+)`)

func parseAssignments(pair string) (assignment1 Assignment, assignment2 Assignment) {
	var err error
	match := exp.FindStringSubmatch(pair)
	result := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	assignment1.min, err = strconv.Atoi(result["min1"])
	if err != nil {
		panic("Failed to parse min1")
	}
	assignment1.max, err = strconv.Atoi(result["max1"])
	if err != nil {
		panic("Failed to parse max1")
	}

	assignment2.min, err = strconv.Atoi(result["min2"])
	if err != nil {
		panic("Failed to parse min2")
	}
	assignment2.max, err = strconv.Atoi(result["max2"])
	if err != nil {
		panic("Failed to parse max2")
	}

	return
}

func isContained(a1 Assignment, a2 Assignment) bool {
	// when a1.min <= a2.min <= a2.max <= a1max and vice versa, we know a2 is contained by a1,
	return (a1.min >= a2.min && a1.max <= a2.max) || (a2.min >= a1.min && a2.max <= a1.max)
}

func hasOverlap(a1 Assignment, a2 Assignment) bool {
	// when a1.min <= a2.min <= a1.max and vice versa, we know a1 and a2 has overlap
	return (a1.min <= a2.min && a2.min <= a1.max) || (a2.min <= a1.min && a1.min <= a2.max)
}

func part1(inputFilePath string) int {
	containedCount := 0
	callback := func(pair string) {
		assignment1, assignment2 := parseAssignments(pair)
		if isContained(assignment1, assignment2) {
			containedCount++
		}
	}

	utils.Readline(inputFilePath, callback)

	return containedCount
}

func part2(inputFilePath string) int {
	overlapCount := 0
	callback := func(pair string) {
		assignment1, assignment2 := parseAssignments(pair)
		if hasOverlap(assignment1, assignment2) {
			overlapCount++
		}
	}

	utils.Readline(inputFilePath, callback)

	return overlapCount
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	inputFilePath := path.Join(path.Dir(filename), "input.txt")

	fmt.Println(part1(inputFilePath))
	fmt.Println(part2(inputFilePath))
}
