package d02

import (
	"fmt"
	"testing"

	"github.com/nozarashi20/aoc-2025/helpers"
)

/*
11-22 has two invalid IDs, 11 and 22.
95-115 has one invalid ID, 99.
998-1012 has one invalid ID, 1010.
1188511880-1188511890 has one invalid ID, 1188511885.
222220-222224 has one invalid ID, 222222.
1698522-1698528 contains no invalid IDs.
446443-446449 has one invalid ID, 446446.
38593856-38593862 has one invalid ID, 38593859.
The rest of the ranges contain no invalid IDs.
Adding up all the invalid IDs in this example produces 1227775554.
*/
func ExampleParseRanges() {
	lines, _ := helpers.ReadFile("data/sample.txt")
	ranges := ParseRanges(lines)
	fmt.Println(ranges)
	// Output: [[11 22] [95 115] [998 1012] [1188511880 1188511890] [222220 222224] [1698522 1698528] [446443 446449] [38593856 38593862] [565653 565659] [824824821 824824827] [2121212118 2121212124]]
}

func TestIsIdValid(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{11, true},
		{13, false},
		{22, true},
		{99, true},
		{1010, true},
		{1188511885, true},
		{222222, true},
		{1698528, false},
		{446446, true},
		{2121212118, false},
	}

	for _, tt := range tests {
		res := IsInvalid(tt.id)
		if res != tt.expected {
			t.Errorf("Expected %v for %v, got %v", tt.expected, tt.id, res)
		}
	}
}

func TestFindInvalidIds(t *testing.T) {
	tests := []struct {
		idRange  [2]int
		expected []int
	}{
		{[2]int{11, 22}, []int{11, 22}},
		{[2]int{95, 115}, []int{99}},
		{[2]int{998, 1012}, []int{1010}},
		{[2]int{1188511880, 1188511890}, []int{1188511885}},
		{[2]int{222220, 222224}, []int{222222}},
		{[2]int{1698522, 1698528}, []int{}},
		{[2]int{446443, 446449}, []int{446446}},
		{[2]int{38593856, 38593862}, []int{38593859}},
		{[2]int{565653, 565659}, []int{}},
		{[2]int{824824821, 824824827}, []int{}},
		{[2]int{2121212118, 2121212124}, []int{}},
	}

	for _, tt := range tests {
		res := FindInvalidIds(tt.idRange)
		if len(res) != len(tt.expected) {
			t.Errorf("Expected %v, got %v", tt.expected, res)
		}
		for i, v := range res {
			if v != tt.expected[i] {
				t.Errorf("Expected %v, got %v", tt.expected, res)
			}
		}
	}
}

func ExamplePartOne() {
	lines, _ := helpers.ReadFile("data/input.txt")
	res := PartOne(lines)
	fmt.Println(res)
	// Output: 18893502033
}

/*
From the same example as before:

11-22 still has two invalid IDs, 11 and 22.
95-115 now has two invalid IDs, 99 and 111.
998-1012 now has two invalid IDs, 999 and 1010.
1188511880-1188511890 still has one invalid ID, 1188511885.
222220-222224 still has one invalid ID, 222222.
1698522-1698528 still contains no invalid IDs.
446443-446449 still has one invalid ID, 446446.
38593856-38593862 still has one invalid ID, 38593859.
565653-565659 now has one invalid ID, 565656.
824824821-824824827 now has one invalid ID, 824824824.
2121212118-2121212124 now has one invalid ID, 2121212121.
Adding up all the invalid IDs in this example produces 4174379265.
*/

func TestIsInvalid2(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{11, true},
		{13, false},
		{22, true},
		{99, true},
		{111, true},
		{999, true},
		{1010, true},
		{1188511885, true},
		{222222, true},
		{1698528, false},
		{446446, true},
		{38593859, true},
		{565656, true},
		{824824824, true},
		{2121212121, true},
		{1212121212, true},
		{1111111, true},
	}

	for _, tt := range tests {
		res := IsInvalid2(tt.id)
		if res != tt.expected {
			t.Errorf("Expected %v for %v, got %v", tt.expected, tt.id, res)
		}
	}
}

func TestFindInvalidIds2(t *testing.T) {
	tests := []struct {
		idRange  [2]int
		expected []int
	}{
		{[2]int{11, 22}, []int{11, 22}},
		{[2]int{95, 115}, []int{99, 111}},
		{[2]int{998, 1012}, []int{999, 1010}},
		{[2]int{1188511880, 1188511890}, []int{1188511885}},
		{[2]int{222220, 222224}, []int{222222}},
		{[2]int{1698522, 1698528}, []int{}},
		{[2]int{446443, 446449}, []int{446446}},
		{[2]int{38593856, 38593862}, []int{38593859}},
		{[2]int{565653, 565659}, []int{565656}},
		{[2]int{824824821, 824824827}, []int{824824824}},
		{[2]int{2121212118, 2121212124}, []int{2121212121}},
	}

	for _, tt := range tests {
		res := FindInvalidIds2(tt.idRange)
		if len(res) != len(tt.expected) {
			t.Errorf("Expected %v, got %v", tt.expected, res)
		}
		for i, v := range res {
			if v != tt.expected[i] {
				t.Errorf("Expected %v, got %v", tt.expected, res)
			}
		}
	}
}

func ExamplePartTwo() {
	lines, _ := helpers.ReadFile("data/input.txt")
	res := PartTwo(lines)
	fmt.Println(res)
	// Output: 26202168557
}
