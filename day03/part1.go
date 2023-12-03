package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func part1(lines []string) int {
	result := 0
	xMax := len(lines) - 1
	yMax := len(lines[0]) - 1
	// fmt.Println("last char: ", lines[0][yMax] == '\n', yMax)
	fmt.Println("xMax = ", xMax, "yMax = ", yMax)
	for i, line := range lines {
		number := ""
		is_part_number := false
		for j, ch := range line {
			if unicode.IsDigit(ch) {
				number += string(ch)
				// check all neighbours including oneself, which doesn't matter
				for _, x := range []int{i - 1, i, i + 1} {
					for _, y := range []int{j - 1, j, j + 1} {
						if 0 <= x && x <= xMax && 0 <= y && y <= yMax {
							if string(lines[x][y]) != "." && !unicode.IsDigit(rune(lines[x][y])) {
								is_part_number = true
							}
						}
					}
				}
			} else {
				if number != "" && is_part_number {
					int_number, err := strconv.Atoi(number)
					if err != nil {
						panic(err)
					}
					// fmt.Println(int_number)
					result += int_number
				}
				number = ""
				is_part_number = false
			}
		}
		if number != "" && is_part_number {
			int_number, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			// fmt.Println(int_number)
			result += int_number
		}
	}
	return result
}
