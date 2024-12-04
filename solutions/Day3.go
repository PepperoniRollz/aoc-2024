package aoc

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pepperonirollz/aoc-2024/utils"
)

func Day3Part1(path string) uint64 {
	data, _ := utils.Parse(path)
	re, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var sum uint64 = 0

	for i := range data {
		matches := re.FindAllStringSubmatch(data[i], -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			product := num1 * num2
			sum += uint64(product)
		}
	}

	return sum
}

func Day3Part2(path string) uint64 {
	data, _ := utils.Parse(path)
	var sum uint64 = 0
	isDoSection := true
	for _, line := range data {
		re := regexp.MustCompile(`(?:mul\((\d+),(\d+)\))|(?:do\(\))|(?:don't\(\))`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			fullMatch := match[0]
			if fullMatch == "do()" {
				isDoSection = true
			} else if fullMatch == "don't()" {
				isDoSection = false

			} else if isDoSection && strings.HasPrefix(fullMatch, "mul") {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				product := num1 * num2
				sum += uint64(product)
			}
		}
	}
	return sum
}
