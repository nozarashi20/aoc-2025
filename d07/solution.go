package d07

import (
	"strings"
)

type Position struct {
	row, col int
}

func PartOne(lines []string) int {
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return SendBeam(grid, Position{row: 1, col: len(grid[0]) / 2})
}

func PartTwo(lines []string) int {
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}
	SendBeam(grid, Position{row: 1, col: len(grid[0]) / 2})

	values := make([][]int, len(grid))
	for row := range values {
		values[row] = make([]int, len(grid[row]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "S" {
				values[i][j] = 1
			}
		}
	}

	for i := 0; i < len(grid)-1; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if values[i][j] == 0 {
				continue
			}
			cell := grid[i][j]
			if cell == "S" || cell == "|" {
				switch grid[i+1][j] {
				case "|":
					values[i+1][j] += values[i][j]
					break
				case "^":
					values[i+1][j-1] += values[i][j]
					values[i+1][j+1] += values[i][j]
					break
				}
			}
		}
	}

	lastRow := values[len(values)-1]
	total := 0
	for _, v := range lastRow {
		total += v
	}

	return total
}

func SendBeam(grid [][]string, pos Position) int {
	var splitCount int
	start := pos.row

	if start >= len(grid) || start < 0 {
		return splitCount
	}
	if pos.col >= len(grid[pos.row]) || pos.col < 0 {
		return splitCount
	}
	for start < len(grid) && grid[start][pos.col] == "." {
		grid[start][pos.col] = "|"
		start++
	}

	if start < len(grid) && grid[start][pos.col] == "^" {
		splitCount++
		splitCount += SendBeam(grid, Position{row: start, col: pos.col - 1})
		splitCount += SendBeam(grid, Position{row: start, col: pos.col + 1})
	}

	return splitCount
}
