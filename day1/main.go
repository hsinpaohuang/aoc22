package main

import (
	"fmt"

	"path"
	"runtime"
	"sort"
	"strconv"

	"github.com/hsinpaohuang/aoc22/utils"
)

func part1(inputFilePath string) int {
	max := 0
	current := 0

	callback := func(line string) {
		if line == "" {
			if max < current {
				max = current
			}
			current = 0
		} else {
			new, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			current += new
		}
	}

	utils.Readline(inputFilePath, callback)

	return max
}

func part2(inputFilePath string) int {
	var sums []int
	current := 0

	callback := func(line string) {
		if line == "" {
			sums = append(sums, current)
			current = 0
		} else {
			new, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			current += new
		}
	}

	utils.Readline(inputFilePath, callback)

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	return sums[0] + sums[1] + sums[2]
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
