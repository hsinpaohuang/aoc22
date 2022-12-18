package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/hsinpaohuang/aoc22/utils"
)

type Directory struct {
	name           string
	subdirectories map[string]*Directory
	sumFileSize    int // file name not needed
	parent         *Directory
}

func (d *Directory) makeDirectory(name string) {
	// create subdirectories map if it doesn't exist
	if d.subdirectories == nil {
		d.subdirectories = make(map[string]*Directory)
	}

	d.subdirectories[name] = &Directory{
		name:        name,
		sumFileSize: 0,
		parent:      d,
	}
}

func (d *Directory) updateFileSize(size int) {
	d.sumFileSize += size
}

func changeDirectory(dir *Directory, name string) *Directory {
	if name == ".." {
		return dir.parent
	} else {
		return dir.subdirectories[name]
	}
}

// recursivly calculate the sum of size of the given directory and its subdirectories
func getDirectorySize(dir Directory) int {
	subDirSize := 0
	for _, d := range dir.subdirectories {
		subDirSize += getDirectorySize(*d)
	}

	return dir.sumFileSize + subDirSize
}

// recursively calculate the total size of directories where the size is less than 100000
func calcTotalSize(cwd Directory) int {
	total := 0

	subDirSize := getDirectorySize(cwd)

	if subDirSize <= 100000 {
		total += subDirSize
	}

	if cwd.subdirectories != nil {
		for _, dir := range cwd.subdirectories {
			total += calcTotalSize(*dir)
		}
	}

	return total
}

func parseLine(cwd *Directory, input string) *Directory {
	if input == "$ cd /" || input == "$ ls" {
		// no action needed
	} else if strings.HasPrefix(input, "$ cd") {
		var dirName string
		fmt.Sscanf(input, "$ cd %s", &dirName)
		return changeDirectory(cwd, dirName)
	} else if strings.HasPrefix(input, "dir") {
		var dirName string
		fmt.Sscanf(input, "dir %s", &dirName)
		cwd.makeDirectory(dirName)
	} else {
		// [fileSize] [fileName]
		var fileSize int
		fmt.Sscanf(input, "%d %s", &fileSize)
		cwd.updateFileSize(fileSize)
	}

	return cwd
}

// recursively find directories where size is greater than minSpace
func findDeleteCandidates(dir Directory, minSpace int) map[string]int {
	candidates := make(map[string]int)

	dirSize := getDirectorySize(dir)
	if dirSize > minSpace {
		candidates[dir.name] = dirSize
	}

	if dir.subdirectories != nil {
		for _, d := range dir.subdirectories {
			subCandidates := findDeleteCandidates(*d, minSpace)
			for name, size := range subCandidates {
				candidates[name] = size
			}
		}
	}

	return candidates
}

// find min size of candidates
func findMin(candidates map[string]int) int {
	var min int
	isFirst := true

	for _, size := range candidates {
		if isFirst {
			isFirst = false
			min = size
		} else if size < min {
			min = size
		}
	}

	return min
}

func part1(inputFilePath string) int {
	rootDir := &Directory{
		name:        "/",
		sumFileSize: 0,
	}
	cwd := rootDir

	callback := func(input string) {
		cwd = parseLine(cwd, input)
	}

	utils.Readline(inputFilePath, callback)

	return calcTotalSize(*rootDir)
}

func part2(inputFilePath string) int {
	rootDir := &Directory{
		name:        "/",
		sumFileSize: 0,
	}
	cwd := rootDir

	callback := func(input string) {
		cwd = parseLine(cwd, input)
	}

	utils.Readline(inputFilePath, callback)

	minSpace := 30000000 - (70000000 - getDirectorySize(*rootDir))
	candidates := findDeleteCandidates(*rootDir, minSpace)

	return findMin(candidates)
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
