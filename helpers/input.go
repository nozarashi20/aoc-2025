package helpers

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ReadFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadInput(f)
}

func Transpose[T any](grid [][]T) [][]T {
	if len(grid) == 0 {
		return nil
	}

	rows := len(grid)
	cols := len(grid[0])

	res := make([][]T, cols)
	for j := 0; j < cols; j++ {
		res[j] = make([]T, rows)
		for i := 0; i < rows; i++ {
			res[j][i] = grid[i][j]
		}
	}
	return res
}

func PrintQuotedSlice[T any](s []T) {
	fmt.Println("[")
	for _, v := range s {
		fmt.Printf("  %q,\n", fmt.Sprint(v))
	}
	fmt.Println("]")
}

func PrintQuotedMatrix[T any](grid [][]T) {
	fmt.Println("[")
	for _, row := range grid {
		fmt.Print("  [")
		for j, v := range row {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%q", fmt.Sprint(v))
		}
		fmt.Println("],")
	}
	fmt.Println("]")
}

// TransposeRagged transposes a possibly ragged 2D slice.
// Missing elements are filled with the zero value of T.
func TransposeRagged[T any](grid [][]T) [][]T {
	if len(grid) == 0 {
		return nil
	}

	maxCols := 0
	for _, row := range grid {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	res := make([][]T, maxCols)
	for j := 0; j < maxCols; j++ {
		res[j] = make([]T, len(grid))
		for i, row := range grid {
			if j < len(row) {
				res[j][i] = row[j]
			}
		}
	}

	return res
}
