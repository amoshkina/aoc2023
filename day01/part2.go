package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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

func replaceWords(line string) string {
	words := []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine",
	}

	if !unicode.IsDigit(rune(line[0])) {
		// replacing the first word
		replaced := false
		for idx := 0; idx < len(line); idx++ {
			for num, word := range words {
				// fmt.Println(line, word, idx)
				if strings.HasPrefix(line[idx:], word) {
					suffix := strings.Replace(line[idx:], word, fmt.Sprint(num+1), 1)
					line = line[:idx] + suffix
					replaced = true
					break
				}
			}
			if replaced {
				break
			}
		}
	}

	if !unicode.IsDigit(rune(line[len(line)-1])) {
		// replacing the last word
		replaced := false

		for idx := len(line) - 1; idx >= 0; idx-- {
			for num, word := range words {
				fmt.Println(line, word, idx, line[idx:])
				if strings.HasPrefix(line[idx:], word) {
					suffix := strings.Replace(line[idx:], word, fmt.Sprint(num+1), 1)
					line = line[:idx] + suffix
					replaced = true
					break
				}
			}
			if replaced {
				break
			}
		}
	}

	return line
}

func main() {
	lines, err := readLines("input.txt")

	if err != nil {
		fmt.Println("Unable to read the input file: ", err)
	}

	var newLines []string

	for _, line := range lines {
		newLines = append(newLines, replaceWords(line))
	}

	var result int

	fmt.Println(newLines)

	for _, line := range newLines {
		first := -1
		var current int
		for _, ch := range line {
			if unicode.IsDigit(ch) {
				if first == -1 {
					first = int(ch - '0')
				}
				current = int(ch - '0')
			}
		}

		result += first*10 + current
	}
	fmt.Println("Part 2: ", result)
}
