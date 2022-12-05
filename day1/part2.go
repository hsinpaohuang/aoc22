package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sums []int
	current := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
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

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	fmt.Println(sums[0] + sums[1] + sums[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
