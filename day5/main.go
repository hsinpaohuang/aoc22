package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/hsinpaohuang/aoc22/utils"
)

type Procedure struct {
	count int
	from  int
	to    int
}

type Stack struct {
	crates []byte
}

// JS style Array.prototype.pop
func (s *Stack) Pop() byte {
	lastIndex := len(s.crates) - 1
	last := s.crates[lastIndex]
	s.crates = s.crates[:lastIndex]
	return last
}

func (s Stack) LastCrate() byte {
	return s.crates[len(s.crates)-1]
}

// end Stack

/*
Stacks are lined up such that every fourth char of each line,
starting from index 1 is the crate.
So all we need to do is filter out the blank ones and prepend it to stack
*/
func parseStartingStacks(input string, stacks []Stack) (hasMoreStacks bool) {
	for i := 1; i < len(input); i += 4 {
		if input == " 1   2   3   4   5   6   7   8   9 " {
			// end of stacks
			return false
		}

		if input[i] != ' ' {
			// prepends stack in place
			stacks[i/4].crates = append([]byte{input[i]}, stacks[i/4].crates...)
		}
	}

	return true
}

// move crates one by one (in place)
func execProcedure(procedure Procedure, stacks []Stack) {
	for i := 0; i < procedure.count; i++ {
		crate := stacks[procedure.from-1].Pop()
		stacks[procedure.to-1].crates = append(stacks[procedure.to-1].crates, crate)
	}
}

// move all crates at once (in place)
func execProcedureMulti(procedure Procedure, stacks []Stack) {
	fromIndex := procedure.from - 1
	toIndex := procedure.to - 1
	// index to split crates to move and crates to stay
	startIndex := len(stacks[fromIndex].crates) - procedure.count

	crates := stacks[fromIndex].crates[startIndex:]
	stacks[fromIndex].crates = stacks[fromIndex].crates[:startIndex]
	stacks[toIndex].crates = append(stacks[toIndex].crates, crates...)
}

// concatenate last crates from stack into a string
func getLastCrates(stacks []Stack) string {
	var lastCrates []byte

	for _, stack := range stacks {
		lastCrates = append(lastCrates, stack.LastCrate())
	}

	return string(lastCrates)
}

func part1(inputFilePath string) string {
	stacks := make([]Stack, 9)
	hasMoreStacks := true
	callback := func(input string) {
		if input == "" {
			return
		}

		if hasMoreStacks {
			hasMoreStacks = parseStartingStacks(input, stacks)
		} else {
			var procedure Procedure
			fmt.Sscanf(
				input,
				"move %d from %d to %d",
				&procedure.count,
				&procedure.from,
				&procedure.to,
			)

			execProcedure(procedure, stacks)
		}
	}

	utils.Readline(inputFilePath, callback)

	return getLastCrates(stacks)
}

func part2(inputFilePath string) string {
	stacks := make([]Stack, 9)
	hasMoreStacks := true
	callback := func(input string) {
		if input == "" {
			return
		}

		if hasMoreStacks {
			hasMoreStacks = parseStartingStacks(input, stacks)
		} else {
			var procedure Procedure
			fmt.Sscanf(
				input,
				"move %d from %d to %d",
				&procedure.count,
				&procedure.from,
				&procedure.to,
			)

			execProcedureMulti(procedure, stacks)
		}
	}

	utils.Readline(inputFilePath, callback)

	return getLastCrates(stacks)
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
