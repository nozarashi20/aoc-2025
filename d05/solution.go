package d05

import (
	"sort"
	"strconv"
	"strings"
)

func ParseDb(lines []string) ([][]int64, []int64) {
	var ranges [][]int64
	var ids []int64
	var separatorLine int64

	for i, line := range lines {
		if line == "" {
			separatorLine = int64(i)
			break
		}
	}

	seen := make(map[string]bool)
	for _, line := range lines[:separatorLine] {
		if seen[line] {
			continue
		}
		seen[line] = true
		rangeParts := strings.Split(line, "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])
		ranges = append(ranges, []int64{int64(start), int64(end)})
	}

	for _, line := range lines[separatorLine+1:] {
		id, _ := strconv.Atoi(line)
		ids = append(ids, int64(id))
	}

	return ranges, ids
}

func PartOne(lines []string) int64 {
	ranges, ids := ParseDb(lines)
	freshCount := make(map[int64]int64)
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				freshCount[id]++
			}
		}
	}
	return int64(len(freshCount))
}

func SortRanges(ranges [][]int64) {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] != ranges[j][0] {
			return ranges[i][0] < ranges[j][0]
		}
		return ranges[i][1] < ranges[j][1]
	})
}

func Overlap(r1, r2 []int64) bool {
	return r1[0] <= r2[1] && r2[0] <= r1[1]
}

func PartTwo(lines []string) int64 {
	ranges, _ := ParseDb(lines)
	SortRanges(ranges)
	mergedRanges := [][]int64{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		last := mergedRanges[len(mergedRanges)-1]
		cur := ranges[i]

		if Overlap(last, cur) {
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else {
			mergedRanges = append(mergedRanges, cur)
		}
	}

	var count int64
	for _, r := range mergedRanges {
		count += r[1] - r[0] + 1
	}
	return count
}
