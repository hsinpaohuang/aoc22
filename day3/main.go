package main

import (
	"fmt"

	"path"
	"runtime"

	"github.com/hsinpaohuang/aoc22/utils"
)

func getRucksackCommonItem(compartment1 string, compartment2 string) rune {
	set := map[rune]bool{}
	for _, item := range compartment1 {
		set[item] = true
	}

	for _, item := range compartment2 {
		if set[item] {
			return item
		}
	}

	panic("Common item not found")
}

func calcPriority(commonItem rune) int {
	if commonItem >= 65 && commonItem <= 90 {
		// A ~ Z, subtract by 38 to make it 27 ~ 52
		return int(commonItem) - 38
	} else if commonItem >= 97 && commonItem <= 122 {
		// a ~ z, subtract by 96 to make it 1 ~ 26
		return int(commonItem) - 96
	}

	panic("Invalid commonItem")
}

func getGroupCommonItem(group []string) rune {
	sets := []map[rune]bool{{}, {}}
	for _, item := range group[0] {
		sets[0][item] = true
	}

	/*
		Alternate between two sets to store common items.

		At the start of this loop, sets[0] contains all items from group[0]
		the first iteration will add common items between set[0] and group[1] into set[1].

		Now that all common items are stored in set[1], set[0] is no longer needed,
		in the next iteration it is recycled
		to store common items between set[1] (which contains common items of group[0:2]) and group[2].

		At the end of the second iteration,
		all common items are stored in set[0], set[1] is no longer needed.

		So set[1] will be recycled in the next iteration
		to store common items between set[0] (which contains common items of group[0:3]) and group[3]
		(should there be more rucksacks in a group).

		This way only 2 maps are required for any group length.

		Because overdesigning is fun ;)
	*/
	for index, rucksack := range group[1:] {
		emptyIndex := (index + 1) % 2 // set that will be storing new commonItems
		currentIndex := index % 2     // set that currently holds commonItems
		for _, item := range rucksack {
			if sets[currentIndex][item] {
				if index+2 == len(group) {
					// common item amongst all rucksacks found
					return item
				} else {
					// add common item amongst all previous sets and current set to empty set
					sets[emptyIndex][item] = true
				}
			}
		}

		sets[currentIndex] = nil // reset to free unused memory
	}

	panic("Common item not found")
}

func part1(inputFilePath string) int {
	priority := 0
	callback := func(rucksack string) {
		middle := len(rucksack) / 2
		compartment1 := rucksack[0:middle]
		compartment2 := rucksack[middle:]
		commonItem := getRucksackCommonItem(compartment1, compartment2)
		priority += calcPriority(commonItem)
	}

	utils.Readline(inputFilePath, callback)

	return priority
}

func part2(inputFilePath string) int {
	priority := 0
	var group []string
	callback := func(rucksack string) {
		group = append(group, rucksack)
		if len(group) == 3 {
			commonItem := getGroupCommonItem(group)
			priority += calcPriority(commonItem)
			group = nil // reset for next group
		}
	}

	utils.Readline(inputFilePath, callback)

	return priority
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
