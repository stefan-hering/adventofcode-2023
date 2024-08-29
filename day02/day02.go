package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day02/input")
	if err != nil {
		panic(err)
	}

	content := string(file)
	part1(content)
	part2(content)
}

func part1(content string) {
	total := 0

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		possible := true

		bag := strings.Split(line, ":")[1]
		gameNumber, _ := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		groups := strings.Split(bag, ";")

		for _, group := range groups {
			cubes := strings.Split(group, ",")
			for _, cube := range cubes {
				colour, n := splitColour(cube)
				if colour == "red" && n > 12 {
					possible = false
				}
				if colour == "green" && n > 13 {
					possible = false
				}
				if colour == "blue" && n > 14 {
					possible = false
				}
			}
		}

		if possible {
			total += gameNumber
		}
	}
	println(total)
}

func part2(content string) {
	total := 0

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		maxBlue := 0
		maxGreen := 0
		maxRed := 0

		bag := strings.Split(line, ":")[1]
		groups := strings.Split(bag, ";")

		for _, group := range groups {
			cubes := strings.Split(group, ",")
			for _, cube := range cubes {
				colour, n := splitColour(cube)
				if colour == "red" {
					maxRed = max(n, maxRed)
				}
				if colour == "green" {
					maxGreen = max(n, maxGreen)
				}
				if colour == "blue" {
					maxBlue = max(n, maxBlue)
				}
			}
		}

		total += maxBlue * maxGreen * maxRed
	}
	println(total)
}

func splitColour(colour string) (string, int) {
	var split = strings.Split(strings.TrimSpace(colour), " ")
	var n, _ = strconv.Atoi(split[0])
	return split[1], n
}
