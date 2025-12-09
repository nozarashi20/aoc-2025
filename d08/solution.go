package d08

import (
	"sort"
	"strconv"
	"strings"
)

// For a better implementation https://fr.wikipedia.org/wiki/Union-find

type Point struct {
	x, y, z, circuit int
}

type Pair struct {
	i, j     int
	distance float64
}

func ParsePoints(lines []string) []Point {
	var points []Point
	for i, line := range lines {
		data := strings.Split(line, ",")
		x, _ := strconv.Atoi(data[0])
		y, _ := strconv.Atoi(data[1])
		z, _ := strconv.Atoi(data[2])
		points = append(points, Point{x: x, y: y, z: z, circuit: i})
	}

	return points
}

func Distance(p, q Point) float64 {
	a := p.x - q.x
	b := p.y - q.y
	c := p.z - q.z
	return float64(a*a + b*b + c*c)
}

func PointCountInCircuit(points []Point, circuit int) int {
	count := 0
	for _, point := range points {
		if point.circuit == circuit {
			count++
		}
	}
	return count
}

func MergeCircuits(points []Point, fromCircuit, toCircuit int) {
	if fromCircuit == toCircuit {
		return
	}
	for i := range points {
		if points[i].circuit == fromCircuit {
			points[i].circuit = toCircuit
		}
	}
}

func CombineCircuits(points []Point, pIdx int, qIdx int) {
	p := points[pIdx]
	q := points[qIdx]
	circuitP := p.circuit
	circuitQ := q.circuit
	if circuitP == circuitQ {
		return
	}

	countP := PointCountInCircuit(points, circuitP)
	countQ := PointCountInCircuit(points, circuitQ)

	if countP < countQ {
		MergeCircuits(points, circuitP, circuitQ)
	} else if countP > countQ {
		MergeCircuits(points, circuitQ, circuitP)
	} else {
		circuit := min(circuitP, circuitQ)
		MergeCircuits(points, circuitP, circuit)
		MergeCircuits(points, circuitQ, circuit)
	}
}

func SortedPairsByDistance(points []Point) []Pair {
	var pairs []Pair
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, Pair{i, j, Distance(points[i], points[j])})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	return pairs
}

func CircuitSizes(points []Point) []int {
	counts := make(map[int]int)
	for i := range points {
		counts[points[i].circuit]++
	}

	sizes := make([]int, 0, len(counts))
	for _, c := range counts {
		sizes = append(sizes, c)
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	return sizes
}

func PartOne(lines []string) int {
	points := ParsePoints(lines)
	pairs := SortedPairsByDistance(points)
	for idx := 0; idx < 1000; idx++ {
		pair := pairs[idx]
		CombineCircuits(points, pair.i, pair.j)
	}

	total := 1
	sizes := CircuitSizes(points)
	for _, size := range sizes[:3] {
		total *= size
	}

	return total
}

func PartTwo(lines []string) int {
	points := ParsePoints(lines)
	pairs := SortedPairsByDistance(points)
	for idx := 0; idx < len(pairs); idx++ {
		pair := pairs[idx]
		CombineCircuits(points, pair.i, pair.j)
		sizes := CircuitSizes(points)
		if len(sizes) == 1 {
			return points[pair.i].x * points[pair.j].x
		}
	}

	return 0
}
