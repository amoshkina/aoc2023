package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"unicode"
// )

// func readLines(path string) ([]string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	return lines, scanner.Err()
// }

// func main() {
// 	lines, err := readLines("input.txt")

// 	if err != nil {
// 		fmt.Println("Unable to read the input file: ", err)
// 	}

// 	var result int

// 	for _, line := range lines {
// 		first := -1
// 		var current int
// 		for _, ch := range line {
// 			if unicode.IsDigit(ch) {
// 				if first == -1 {
// 					first = int(ch - '0')
// 				}
// 				current = int(ch - '0')
// 			}
// 		}

// 		result += first*10 + current
// 	}
// 	fmt.Println("Part 1: ", result)
// }
