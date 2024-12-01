package aoc

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/pepperonirollz/aoc-2024/utils"
)

func Part1(path string) int64 {
	data, _ := utils.Parse(path)
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range data {
		split := strings.Fields(line)
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])
		left = append(left, l)
		right = append(right, r)

	}
	sort.Ints(left)
	sort.Ints(right)
	var sum int64
	for i := range left {
		diff := int64(math.Abs(float64(right[i] - left[i])))
		sum += diff
	}
	return sum
}

func Part2(path string) int64 {
	data, _ := utils.Parse(path)
	m := make(map[int64]int64)
	right := make([]int, 0)

	for _, line := range data {
		split := strings.Fields(line)
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])
		m[int64(l)] = 0
		right = append(right, r)
	}

	for _, r := range right {
		if _, exists := m[int64(r)]; exists {
			m[int64(r)]++
		}
	}
	var sum int64 = 0
	for k, v := range m {
		sum += k * v
	}
	return sum
}
