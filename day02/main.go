package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("input.txt")

	if err != nil {
		fmt.Println("Unable to read the input file: ", err)
	}

	fmt.Println("Part 1: ", part1(lines))
	// fmt.Println("Part 2: ", part2(lines))
}
