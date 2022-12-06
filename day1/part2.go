package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/hsinpaohuang/aoc22/utils"
)

func main() {
	var sums []int
	current := 0

	filepath := "./input/day1/input.txt"
	callback := func(line string) {
		if line == "" {
			sums = append(sums, current)
			current = 0
		} else {
			new, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			current += new
		}
	}

	utils.Readline(filepath, callback)

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	fmt.Println(sums[0] + sums[1] + sums[2])
}
