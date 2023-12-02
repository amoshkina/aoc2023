package main

import (
	"strconv"
	"strings"
)

func part2(games []string) int {
	var result int
	for _, game := range games {
		var minColor = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		sets := strings.Split(strings.Split(game, ":")[1], ";")
		for _, set := range sets {
			for _, cubes := range strings.Split(set, ",") {
				data := strings.Split(strings.TrimSpace(cubes), " ")
				color := data[1]
				amount, err := strconv.Atoi(data[0])
				if err != nil {
					panic(err)
				}
				if amount > minColor[color] {
					minColor[color] = amount
				}
			}
		}
		result += minColor["red"] * minColor["green"] * minColor["blue"]

	}
	return result
}
