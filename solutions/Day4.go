package aoc

import (
	"github.com/pepperonirollz/aoc-2024/utils"
)

func Day4Part1(path string) int {
	data, _ := utils.Parse(path)
	sum := 0
	rows := len(data)
	cols := len(data[0])
	for row := range data {
		for col := range data[row] {
			sum += search(row, col, rows, cols, data)
		}
	}

	return sum
}

func Day4Part2(path string) int {
	data, _ := utils.Parse(path)
	sum := 0
	rows := len(data)
	cols := len(data[0])
	for row := range data {
		for col := range data[row] {
			sum += searchXmas(row, col, rows, cols, data)
		}
	}

	return sum
}

func search(x int, y int, rows int, cols int, grid []string) int {
	inUpperRowBounds := x-3 >= 0
	inLowerRowBounds := x+3 <= rows-1
	inUpperColBounds := y+3 <= cols-1
	inLowerColBounds := y-3 >= 0
	count := 0
	if grid[x][y] == 'X' {
		//up
		if inUpperRowBounds && grid[x-1][y] == 'M' && grid[x-2][y] == 'A' && grid[x-3][y] == 'S' {
			count++
		}
		//down
		if inLowerRowBounds && grid[x+1][y] == 'M' && grid[x+2][y] == 'A' && grid[x+3][y] == 'S' {
			count++
		}
		//left
		if inLowerColBounds && grid[x][y-1] == 'M' && grid[x][y-2] == 'A' && grid[x][y-3] == 'S' {
			count++
		}
		//right
		if inUpperColBounds && grid[x][y+1] == 'M' && grid[x][y+2] == 'A' && grid[x][y+3] == 'S' {
			count++
		}
		//topleft
		if inUpperRowBounds && inLowerColBounds && grid[x-1][y-1] == 'M' && grid[x-2][y-2] == 'A' && grid[x-3][y-3] == 'S' {
			count++
		}
		//toright
		if inUpperRowBounds && inUpperColBounds && grid[x-1][y+1] == 'M' && grid[x-2][y+2] == 'A' && grid[x-3][y+3] == 'S' {
			count++
		}
		//bottomleft
		if inLowerRowBounds && inLowerColBounds && grid[x+1][y-1] == 'M' && grid[x+2][y-2] == 'A' && grid[x+3][y-3] == 'S' {
			count++
		}
		//bottomright
		if inLowerRowBounds && inUpperColBounds && grid[x+1][y+1] == 'M' && grid[x+2][y+2] == 'A' && grid[x+3][y+3] == 'S' {
			count++
		}
	}
	return count
}

func searchXmas(x int, y int, rows int, cols int, grid []string) int {
	inUpperRowBounds := x-1 >= 0
	inLowerRowBounds := x+1 <= rows-1
	inUpperColBounds := y+1 <= cols-1
	inLowerColBounds := y-1 >= 0
	inBounds := inLowerColBounds && inUpperColBounds && inUpperRowBounds && inLowerRowBounds
	count := 0
	if grid[x][y] == 'A' && inBounds {
		if (grid[x-1][y-1] == 'M' && grid[x+1][y+1] == 'S' || grid[x-1][y-1] == 'S' && grid[x+1][y+1] == 'M') && (grid[x-1][y+1] == 'M' && grid[x+1][y-1] == 'S' || grid[x-1][y+1] == 'S' && grid[x+1][y-1] == 'M') {
			count++
		}
	}
	return count
}
