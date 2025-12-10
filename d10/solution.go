package d10

import (
	"regexp"
	"strconv"
	"strings"
)

type Vec uint64

func VecFromPattern(pattern string) Vec {
	var v Vec
	for i, light := range pattern {
		if light == '#' {
			v |= 1 << i
		}
	}

	return v
}

func ButtonFromIndices(indices []int) Vec {
	var button Vec
	for _, i := range indices {
		button |= 1 << i
	}
	return button
}

type State struct {
	mask    Vec
	presses int
}

type Machine struct {
	target  Vec
	buttons []Vec
}

func ParseMachine(line string) Machine {
	m := Machine{}
	regex := regexp.MustCompile(`^\[([.#]+)]\s(.+)\s(.+)$`)
	if !regex.MatchString(line) {
		panic("invalid input")
	}

	match := regex.FindStringSubmatch(line)
	pattern := match[1]
	var buttons []Vec

	for _, wiring := range strings.Split(match[2], " ") {
		var indices []int
		for _, index := range strings.Split(wiring[1:len(wiring)-1], ",") {
			n, err := strconv.Atoi(index)
			if err != nil {
				panic(err)
			}
			indices = append(indices, n)
		}
		buttons = append(buttons, ButtonFromIndices(indices))
	}

	m.target = VecFromPattern(pattern)
	m.buttons = buttons

	return m
}

func minPresses(m Machine) int {
	if m.target == 0 {
		return 0
	}

	visited := make(map[Vec]int)
	queue := []State{{0, 0}}
	visited[0] = 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, b := range m.buttons {
			next := cur.mask ^ b
			if _, seen := visited[next]; seen {
				continue
			}
			p := cur.presses + 1
			if next == m.target {
				return p
			}
			visited[next] = p
			queue = append(queue, State{mask: next, presses: p})
		}
	}

	return -1
}

func PartOne(lines []string) int {
	total := 0
	for _, line := range lines {
		m := ParseMachine(line)
		total += minPresses(m)
	}
	return total
}
