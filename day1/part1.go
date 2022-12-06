package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hsinpaohuang/aoc22/utils"
)

func main() {
	max := 0
	current := 0

	filepath := "./input/day1/input.txt"
	callback := func(line string) {
		if line == "" {
			if max < current {
				max = current
			}
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

	fmt.Println(max)
}
