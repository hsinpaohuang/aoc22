package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	max := 0
	current := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if max < current {
				max = current
			}
			current = 0
		} else {
			new, _ := strconv.Atoi(line)
			current += new
		}
	}

	fmt.Println(max)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
