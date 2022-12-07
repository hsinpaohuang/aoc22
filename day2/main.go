package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/hsinpaohuang/aoc22/utils"
)

type Choice int8

const (
	Rock     Choice = iota
	Paper           = iota
	Scissors        = iota
)

type Outcome int8

const (
	Lose Outcome = iota
	Draw         = iota
	Win          = iota
)

func calcScoreFromResult(elfChoice Choice, myChoice Choice) int {
	switch myChoice - elfChoice {
	case 0:
		return 3
	case -2: // wraparound
		fallthrough
	case 1:
		return 6
	case 2: //wraparound
		fallthrough
	case -1:
		return 0
	default:
		panic("Invalid result")
	}
}

func calcScoreFromChoice(choice Choice) int {
	switch choice {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		panic("Invalid Choice")
	}
}

func getChoiceFromOutcome(elfChoice Choice, outcome Outcome) Choice {
	switch outcome {
	case Win:
		if elfChoice == Scissors {
			// wraparound
			return Rock
		} else {
			return elfChoice + 1
		}
	case Draw:
		return elfChoice
	case Lose:
		if elfChoice == Rock {
			// wraparound
			return Scissors
		} else {
			return elfChoice - 1
		}
	default:
		panic("Invalid elfChoice")
	}
}

func part1(inputFilePath string) int {
	score := 0

	callback := func(line string) {
		if line == "" {
			return
		}

		// use ascii int to calculate difference
		elfChoice := Choice(int(line[0]) - int('A'))
		myChoice := Choice(int(line[2]) - int('X'))

		score += calcScoreFromResult(elfChoice, myChoice)
		score += calcScoreFromChoice(myChoice)
	}

	utils.Readline(inputFilePath, callback)

	return score
}

func part2(inputFilePath string) int {
	score := 0

	callback := func(line string) {
		if line == "" {
			return
		}

		// use ascii int to calculate difference
		elfChoice := Choice(int(line[0]) - int('A'))
		outcome := Outcome(int(line[2]) - int('X'))
		myChoice := getChoiceFromOutcome(elfChoice, outcome)

		score += calcScoreFromResult(elfChoice, myChoice)
		score += calcScoreFromChoice(myChoice)
	}

	utils.Readline(inputFilePath, callback)

	return score
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
