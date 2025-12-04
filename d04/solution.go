package d04

import (
	"strings"
)

func ParseGrid(lines []string) [][]string {
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func GridToString(grid [][]string) string {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(strings.Join(row, ""))
		sb.WriteString("\n")
	}
	return sb.String()
}

func EightAdjacents(grid [][]string, x, y int) int {
	count := 0
	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= len(grid) {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if j < 0 || j >= len(grid[i]) {
				continue
			}
			if grid[i][j] == "@" {
				count++
			}
		}
	}
	return count
}

func TraverseGrid(grid [][]string) ([][]string, int) {
	var accessibleCells [][2]int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == "@" && EightAdjacents(grid, x, y) < 4 {
				accessibleCells = append(accessibleCells, [2]int{x, y})
			}
		}
	}

	for _, cell := range accessibleCells {
		grid[cell[0]][cell[1]] = "x"
	}

	return grid, len(accessibleCells)
}

func PartOne(lines []string) int {
	grid := ParseGrid(lines)
	_, count := TraverseGrid(grid)
	return count
}

func PartTwo(lines []string) int {
	grid := ParseGrid(lines)
	totalCount := 0
	newGrid, count := TraverseGrid(grid)
	totalCount += count
	for count > 0 {
		newGrid, count = TraverseGrid(newGrid)
		totalCount += count
	}

	return totalCount
}
