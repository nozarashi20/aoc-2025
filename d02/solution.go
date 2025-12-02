package d02

import (
	"strconv"
	"strings"
)

func PartOne(lines []string) int {
	idRanges := ParseRanges(lines)
	sum := 0
	for _, idRange := range idRanges {
		invalidIds := FindInvalidIds(idRange)
		for _, id := range invalidIds {
			sum += id
		}
	}
	return sum
}

func PartTwo(lines []string) int {
	idRanges := ParseRanges(lines)
	sum := 0
	for _, idRange := range idRanges {
		invalidIds := FindInvalidIds2(idRange)
		for _, id := range invalidIds {
			sum += id
		}
	}
	return sum
}

func ParseRanges(lines []string) [][2]int {
	input := lines[0]
	parts := strings.Split(input, ",")
	var ranges [][2]int
	for _, part := range parts {
		rangeParts := strings.Split(part, "-")
		from, _ := strconv.Atoi(rangeParts[0])
		to, _ := strconv.Atoi(rangeParts[1])
		ranges = append(ranges, [2]int{from, to})
	}

	return ranges
}

func FindInvalidIds(idRange [2]int) []int {
	var invalidIds []int
	for id := idRange[0]; id <= idRange[1]; id++ {
		if IsInvalid(id) {
			invalidIds = append(invalidIds, id)
		}
	}
	return invalidIds
}

func IsInvalid(id int) bool {
	s := strconv.Itoa(id)
	if len(s)%2 != 0 {
		return false
	}

	mid := len(s) / 2
	left := s[:mid]
	right := s[mid:]

	return left == right
}

func FindInvalidIds2(idRange [2]int) []int {
	var invalidIds []int
	for id := idRange[0]; id <= idRange[1]; id++ {
		if IsInvalid2(id) {
			invalidIds = append(invalidIds, id)
		}
	}
	return invalidIds
}

func IsInvalid2(id int) bool {
	s := strconv.Itoa(id)
	n := len(s)

	for divisor := 1; divisor <= n/2; divisor++ {
		if n%divisor == 0 {
			pattern := s[:divisor]
			repeats := n / divisor

			if strings.Repeat(pattern, repeats) == s {
				return true
			}
		}
	}

	return false
}
