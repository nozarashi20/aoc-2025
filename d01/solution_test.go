package d01

import (
	"fmt"

	"github.com/nozarashi20/aoc-2025/helpers"
)

func ExamplePartOne() {
	lines, _ := helpers.ReadFile("data/sample.txt")
	res := PartOne(lines)
	fmt.Println(res)
	// Output: 3
}

func ExamplePartTwo() {
	lines, _ := helpers.ReadFile("data/sample.txt")
	res := PartTwo(lines)
	fmt.Println(res)
	// Output: 6
}
