package main

import (
	"strconv"
	"strings"
)

// 12 red cubes, 13 green cubes, and 14 blue
var redMax int = 12
var greenMax int = 13
var blueMax int = 14

var maxColor = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1(games []string) int {
	var result int
	for gameId, game := range games {
		gamePossible := true
		sets := strings.Split(strings.Split(game, ":")[1], ";")
		for _, set := range sets {
			if !gamePossible {
				break
			}
			for _, cubes := range strings.Split(set, ",") {
				data := strings.Split(strings.TrimSpace(cubes), " ")
				color := data[1]
				amount, err := strconv.Atoi(data[0])
				if err != nil {
					panic(err)
				}
				if amount > maxColor[color] {
					gamePossible = false
					break
				}
			}
		}
		if gamePossible {
			result += gameId + 1
		}

	}
	return result
}
