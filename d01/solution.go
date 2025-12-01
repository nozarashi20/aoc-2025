package d01

import (
	"strconv"
)

func PartOne(lines []string) int {
	password := 0
	pos := 50
	for _, line := range lines {
		direction, rest := line[:1], line[1:]
		n, err := strconv.Atoi(rest)
		if err != nil {
			panic(err)
		}
		pos = rotate(direction, pos, n)
		if pos == 0 {
			password++
		}
	}
	return password
}

func PartTwo(lines []string) int {
	password := 0
	pos := 50
	for _, line := range lines {
		direction, rest := line[:1], line[1:]
		n, err := strconv.Atoi(rest)
		if err != nil {
			panic(err)
		}
		start := pos
		pos = rotate(direction, pos, n)
		crossingCount := zeroCrossingCount(direction, start, n)
		password += crossingCount
	}
	return password
}

func rotate(direction string, pos, moves int) int {
	if direction == "L" {
		moves = -1 * moves
	}

	return ((pos+moves)%100 + 100) % 100
}

func zeroCrossingCount(direction string, pos, moves int) int {
	if moves == 0 {
		return 0
	}

	if direction == "R" {
		return (pos + moves) / 100
	}

	if moves < pos {
		return 0
	}

	if pos == 0 {
		return moves / 100
	}

	return 1 + (moves-pos)/100
}
