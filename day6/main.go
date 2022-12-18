package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/hsinpaohuang/aoc22/utils"
)

/*
 ** find duplicate character in string, and returns the index of the first occurance
 ** if no duplicate characters are found, returns -1
 */
func findDuplicateIndex(input string) int {
	set := make(map[rune]int)
	for index, char := range input {
		if val, ok := set[char]; ok {
			return val
		}

		set[char] = index
	}

	return -1
}

func parseStart(input string, markerLen int) int {
	marker := 0
	for i := 0; i < len(input)-markerLen; i++ {
		duplicateIndex := findDuplicateIndex(input[i : i+markerLen])
		if duplicateIndex != -1 {
			// skip duplicateIndex indexes because there will be duplicate in those substrings
			i += duplicateIndex
		} else {
			marker = i
			break
		}
	}

	return marker + markerLen
}

func part1(inputFilePath string) int {
	marker := 0
	callback := func(input string) {
		marker = parseStart(input, 4)
	}

	utils.Readline(inputFilePath, callback)

	return marker
}

func part2(inputFilePath string) int {
	marker := 0
	callback := func(input string) {
		marker = parseStart(input, 14)
	}

	utils.Readline(inputFilePath, callback)

	return marker
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
