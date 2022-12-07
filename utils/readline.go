package utils

import (
	"bufio"
	"os"
)

func Readline(file_path string, callback func(line string)) {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		callback(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
