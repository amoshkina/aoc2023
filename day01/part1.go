package main

import (
	"unicode"
)

func part1(lines []string) int {
	var result int

	for _, line := range lines {
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
	return result
}
