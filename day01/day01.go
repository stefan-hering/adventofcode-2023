package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day01/input")
	if err != nil {
		panic(err)
	}

	content := string(file)
	part1(content)
	part2(content)
}

func part1(content string) {
	sum := 0
	for _, line := range strings.Split(content, "\n") {
		first, last := 'a', 'a'
		for _, char := range line {
			if char >= 48 && char <= 57 {
				if first == 'a' {
					first = char
				}
				last = char
			}
		}
		num, _ := strconv.Atoi(string(first) + string(last))
		sum += num
	}
	println(sum)
}

func part2(content string) {
	result := strings.Replace(content, "one", "o1ne", -1)
	result = strings.Replace(result, "two", "t2wo", -1)
	result = strings.Replace(result, "three", "t3hree", -1)
	result = strings.Replace(result, "four", "f4our", -1)
	result = strings.Replace(result, "five", "f5ive", -1)
	result = strings.Replace(result, "six", "s6ix", -1)
	result = strings.Replace(result, "seven", "s7even", -1)
	result = strings.Replace(result, "eight", "e8ight", -1)
	result = strings.Replace(result, "nine", "n9ine", -1)
	part1(result) // 53652
}
