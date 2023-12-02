package main

import (
	"fmt"
	"strings"
	"unicode"
)

var words []string = []string{
	"one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine",
}

func replaceOnce(line string, start int, inc int) string {
	if !unicode.IsDigit(rune(line[start])) {
		replaced := false
		for idx := start; idx < len(line) && idx >= 0; idx += inc {
			for num, word := range words {
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

func replaceWords(line string) string {

	line = replaceOnce(line, 0, 1)
	line = replaceOnce(line, len(line)-1, -1)

	return line
}

func part2(lines []string) int {
	var newLines []string

	for _, line := range lines {
		newLines = append(newLines, replaceWords(line))
	}

	return part1(newLines)
}
