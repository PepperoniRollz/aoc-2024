package aoc

import (
	"math"
	"strconv"
	"strings"

	"github.com/pepperonirollz/aoc-2024/utils"
)

func Day2Part1(path string) int {

	data, _ := utils.Parse(path)
	count := 0
	for _, line := range data {
		split := strings.Split(line, " ")
		report := make([]int, len(split))
		for i, s := range split {
			report[i], _ = strconv.Atoi(s)
		}
		if isSafe(report) {
			count++
		}
	}
	return count
}

func Day2Part2(path string) int {
	data, _ := utils.Parse(path)
	count := 0

	for _, line := range data {
		split := strings.Split(line, " ")
		report := make([]int, len(split))
		for i, s := range split {
			report[i], _ = strconv.Atoi(s)
		}

		if isSafe(report) {
			count++
			continue
		}

		for i := range report {
			temp := make([]int, 0, len(report)-1)
			temp = append(temp, report[:i]...)
			temp = append(temp, report[i+1:]...)

			if isSafe(temp) {
				count++
				break
			}
		}
	}
	return count
}

func isSafe(report []int) bool {
	var isSafeAscending bool
	var isSafeDescending bool
	for i := 0; i < len(report)-1; i++ {
		diff := math.Abs(float64(report[i+1] - report[i]))
		diffCondition := diff >= 1 && diff <= 3
		if report[i+1] > report[i] && diffCondition {
			isSafeAscending = true
		} else {
			isSafeAscending = false
			break
		}
	}
	for i := 0; i < len(report)-1; i++ {
		diff := math.Abs(float64(report[i+1] - report[i]))
		diffCondition := diff >= 1 && diff <= 3
		if report[i+1] < report[i] && diffCondition {
			isSafeDescending = true
		} else {
			isSafeDescending = false
			break
		}
	}
	return isSafeAscending || isSafeDescending
}
