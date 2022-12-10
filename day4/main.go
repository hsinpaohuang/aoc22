package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/hsinpaohuang/aoc22/utils"
)

type Assignment struct {
	min int
	max int
}

func part1(inputFilePath string) int {
	containedCount := 0
	callback := func(pair string) {
		var (
			a1 Assignment
			a2 Assignment
		)

		fmt.Sscanf(pair, "%d-%d,%d-%d", &a1.min, &a1.max, &a2.min, &a2.max)

		// when a1.min <= a2.min <= a1.max and vice versa,
		// we know a1 and a2 has overlap
		if (a1.min >= a2.min && a1.max <= a2.max) ||
			(a2.min >= a1.min && a2.max <= a1.max) {
			containedCount++
		}
	}

	utils.Readline(inputFilePath, callback)

	return containedCount
}

func part2(inputFilePath string) int {
	overlapCount := 0
	callback := func(pair string) {
		var (
			a1 Assignment
			a2 Assignment
		)

		fmt.Sscanf(pair, "%d-%d,%d-%d", &a1.min, &a1.max, &a2.min, &a2.max)

		// when a1.min <= a2.min <= a1.max and vice versa,
		// we know a1 and a2 has overlap
		if (a1.min <= a2.min && a2.min <= a1.max) ||
			(a2.min <= a1.min && a1.min <= a2.max) {
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
