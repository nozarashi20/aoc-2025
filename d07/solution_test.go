package d07

import (
	"fmt"

	"github.com/nozarashi20/aoc-2025/helpers"
)

func ExamplePartOne() {
	lines, _ := helpers.ReadFile("data/input.txt")
	res := PartOne(lines)
	fmt.Println(res)
	// Output: 1535
}

func ExamplePartTwo() {
	lines, _ := helpers.ReadFile("data/input.txt")
	res := PartTwo(lines)
	fmt.Println(res)
	// Output: 4404709551015
}
