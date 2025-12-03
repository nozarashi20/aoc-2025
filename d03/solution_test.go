package d03

import (
	"fmt"
	"testing"

	"github.com/nozarashi20/aoc-2025/helpers"
)

/*
You'll need to find the largest possible joltage each bank can produce. In the above example:

In 987654321111111, you can make the largest joltage possible, 98, by turning on the first two batteries.
In 811111111111119, you can make the largest joltage possible by turning on the batteries labeled 8 and 9, producing 89 jolts.
In 234234234234278, you can make 78 by turning on the last two batteries (marked 7 and 8).
In 818181911112111, the largest joltage you can produce is 92.
*/

func TestLargestJoltage(t *testing.T) {
	tests := []struct {
		bank     string
		expected int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, tt := range tests {
		res := LargestJoltage(tt.bank)
		if res != tt.expected {
			t.Errorf("Expected %v for %v, got %v", tt.expected, tt.bank, res)
		}
	}
}

/*
Now, the joltages are much larger:

In 987654321111111, the largest joltage can be found by turning on everything except some 1s at the end to produce 987654321111.
In the digit sequence 811111111111119, the largest joltage can be found by turning on everything except some 1s, producing 811111111119.
In 234234234234278, the largest joltage can be found by turning on everything except a 2 battery, a 3 battery, and another 2 battery near the start to produce 434234234278.
In 818181911112111, the joltage 888911112111 is produced by turning on everything except some 1s near the front.
*/
func TestLargestJoltage2(t *testing.T) {
	tests := []struct {
		bank     string
		expected int64
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, tt := range tests {
		res := LargestJoltage2(tt.bank)
		if res != tt.expected {
			t.Errorf("Expected %v for %v, got %v", tt.expected, tt.bank, res)
		}
	}
}

func ExamplePartOne() {
	lines, _ := helpers.ReadFile("data/input.txt")
	res := PartOne(lines)
	fmt.Println(res)
	// Output: 17179
}

func ExamplePartTwo() {
	lines, _ := helpers.ReadFile("data/input.txt")
	res := PartTwo(lines)
	fmt.Println(res)
	// Output: 170025781683941
}
