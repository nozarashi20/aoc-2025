package d09

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// See https://en.wikipedia.org/wiki/Flood_fill (using BFS here)
// https://en.wikipedia.org/wiki/Prefix_sum
// https://www.geeksforgeeks.org/dsa/coordinate-compression

type Point struct {
	x, y int
}

// Cell represents a cell in the compressed grid,
// addressed by indices into the xs/ys coordinate slices.
type Cell struct {
	i, j int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func area(a, b Point) int {
	return (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// buildCompressedCoords constructs a compressed coordinate system from all
// polygon corner points.
//
// It collects distinct x and y values from corners, then extends the range
// with a one-cell "frame" around the data (min-1 and max+1 in each axis)
// so flood fill can start from a guaranteed outside cell.
//
// It returns:
//   - xs, ys: sorted unique x and y coordinates
//   - xIndex, yIndex: maps from original x/y to compressed indices
func buildCompressedCoords(corners []Point) (xs, ys []int, xIndex, yIndex map[int]int) {
	xSet := make(map[int]struct{})
	ySet := make(map[int]struct{})

	// Track min/max over input corners.
	minX, maxX := 0, 0
	minY, maxY := 0, 0
	first := true

	for _, p := range corners {
		xSet[p.x] = struct{}{}
		ySet[p.y] = struct{}{}

		if first {
			minX, maxX = p.x, p.x
			minY, maxY = p.y, p.y
			first = false
		} else {
			if p.x < minX {
				minX = p.x
			}
			if p.x > maxX {
				maxX = p.x
			}
			if p.y < minY {
				minY = p.y
			}
			if p.y > maxY {
				maxY = p.y
			}
		}
	}

	// Add one "frame" coordinate around the data so flood fill can start outside.
	xSet[minX-1] = struct{}{}
	xSet[maxX+1] = struct{}{}
	ySet[minY-1] = struct{}{}
	ySet[maxY+1] = struct{}{}

	// Convert sets to sorted slices.
	for x := range xSet {
		xs = append(xs, x)
	}
	for y := range ySet {
		ys = append(ys, y)
	}
	sort.Ints(xs)
	sort.Ints(ys)

	xIndex = make(map[int]int, len(xs))
	for i, x := range xs {
		xIndex[x] = i
	}
	yIndex = make(map[int]int, len(ys))
	for j, y := range ys {
		yIndex[y] = j
	}

	return xs, ys, xIndex, yIndex
}

// buildWallsCompressed converts polygon edges from original coordinates into
// wall cells in the compressed grid.
//
// Using the polygon corners in order, it draws each horizontal/vertical edge
// between consecutive corners and marks every cell along those edges as a wall
// in compressed coordinates.
func buildWallsCompressed(corners []Point, xIndex, yIndex map[int]int) map[Cell]bool {
	walls := make(map[Cell]bool)
	n := len(corners)
	if n == 0 {
		return walls
	}

	for k := 0; k < n; k++ {
		a := corners[k]

		// closed polygon
		// https://leetcode.com/problems/maximum-sum-circular-subarray/description
		// A circular array means the end of the array connects to the beginning of the array.
		// Formally, the next element of nums[i] is nums[(i + 1) % n]
		// and the previous element of nums[i] is nums[(i - 1 + n) % n].
		b := corners[(k+1)%n]

		ia := xIndex[a.x]
		ja := yIndex[a.y]
		ib := xIndex[b.x]
		jb := yIndex[b.y]

		if ia == ib {
			// Vertical segment in compressed grid.
			step := 1
			if jb < ja {
				step = -1
			}
			for j := ja; ; j += step {
				walls[Cell{ia, j}] = true
				if j == jb {
					break
				}
			}
		} else if ja == jb {
			// Horizontal segment in compressed grid.
			step := 1
			if ib < ia {
				step = -1
			}
			for i := ia; ; i += step {
				walls[Cell{i, ja}] = true
				if i == ib {
					break
				}
			}
		} else {
			panic(fmt.Sprintf("non axis-aligned segment between %v and %v", a, b))
		}
	}

	return walls
}

// floodOutsideCompressed performs a BFS flood fill on the compressed grid,
// starting from the start (a known outside cell).
//
// The flood cannot pass through cells marked as walls. It returns a map
// where a key c has value true if and only if c is reachable from the start
// without crossing a wall (i.e., c is outside the polygon).
func floodOutsideCompressed(walls map[Cell]bool, w, h int, start Cell) map[Cell]bool {
	outside := make(map[Cell]bool)
	queue := []Cell{start}
	outside[start] = true

	dirs := []Cell{
		{1, 0}, {-1, 0},
		{0, 1}, {0, -1},
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			nxt := Cell{cur.i + d.i, cur.j + d.j}

			// Out of bounds?
			if nxt.i < 0 || nxt.i >= w || nxt.j < 0 || nxt.j >= h {
				continue
			}

			// Wall blocks the flood.
			if walls[nxt] {
				continue
			}

			// Already known to be outside.
			if outside[nxt] {
				continue
			}

			outside[nxt] = true
			queue = append(queue, nxt)
		}
	}

	return outside
}

// buildBadGridAndPrefix builds a 2D grid bad and a 2D prefix sum ps from the
// set of outside cells.
//
//   - bad[j][i] = 1 if cell (i,j) is outside, 0 otherwise.
//   - ps has dimensions (h+1) x (w+1) and satisfies:
//     ps[y][x] = sum of bad[0..y-1][0..x-1].
func buildBadGridAndPrefix(outside map[Cell]bool, w, h int) (bad [][]int, ps [][]int) {
	// bad[j][i] = 1 if outside, 0 otherwise.
	bad = make([][]int, h)
	for j := 0; j < h; j++ {
		bad[j] = make([]int, w)
		for i := 0; i < w; i++ {
			if outside[Cell{i, j}] {
				bad[j][i] = 1
			}
		}
	}

	// ps has one extra row and column for simpler math.
	ps = make([][]int, h+1)
	for j := 0; j <= h; j++ {
		ps[j] = make([]int, w+1)
	}

	// Standard 2D prefix sum:
	// ps[j+1][i+1] = bad[j][i] + ps[j][i+1] + ps[j+1][i] - ps[j][i].
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			ps[j+1][i+1] = bad[j][i] +
				ps[j][i+1] +
				ps[j+1][i] -
				ps[j][i]
		}
	}

	return bad, ps
}

// queryOutsideCount returns how many "outside" cells lie inside the inclusive
// rectangle [minI..maxI] x [minJ..maxJ] in compressed indices, using the
// 2D prefix-sum array ps.
func queryOutsideCount(ps [][]int, minI, maxI, minJ, maxJ int) int {
	// Convert inclusive rectangle to prefix-sum coordinates.
	// Offsets of +1 are used because ps is (h+1) x (w+1).
	x1 := minI
	y1 := minJ
	x2 := maxI
	y2 := maxJ

	return ps[y2+1][x2+1] -
		ps[y1][x2+1] -
		ps[y2+1][x1] +
		ps[y1][x1]
}

// isRectangleValidCompressed reports whether the axis-aligned rectangle with
// corners a and b contains no "outside" cells in the compressed grid.
//
// The check works as follows:
//   - a and b are original coordinates (polygon corners).
//   - xIndex and yIndex map original coordinates to compressed indices.
//   - the resulting rectangle in compressed indices is constructed.
//   - queryOutsideCount is used to count outside cells inside that rectangle,
//     including its borders.
func isRectangleValidCompressed(a, b Point, xIndex, yIndex map[int]int, ps [][]int) bool {
	if a.x == b.x && a.y == b.y {
		return false
	}

	ia, ja := xIndex[a.x], yIndex[a.y]
	ib, jb := xIndex[b.x], yIndex[b.y]

	// Compressed rectangle bounds.
	minI, maxI := ia, ib
	if minI > maxI {
		minI, maxI = maxI, minI
	}
	minJ, maxJ := ja, jb
	if minJ > maxJ {
		minJ, maxJ = maxJ, minJ
	}

	// Count how many "outside" cells lie inside this compressed rectangle.
	sumOutside := queryOutsideCount(ps, minI, maxI, minJ, maxJ)

	return sumOutside == 0
}

// findBestRectangleCompressed iterates over all pairs of polygon corners and
// returns the largest-area axis-aligned rectangle (in original coordinates)
// that is fully valid, according to isRectangleValidCompressed.
//
// It returns the best area and the two corner points that achieve it.
func findBestRectangleCompressed(corners []Point, xIndex, yIndex map[int]int, ps [][]int) (bestArea int, bestA, bestB Point) {
	bestArea = 0
	var ba, bb Point

	for i := 0; i < len(corners); i++ {
		for j := i + 1; j < len(corners); j++ {
			a := corners[i]
			b := corners[j]

			if !isRectangleValidCompressed(a, b, xIndex, yIndex, ps) {
				continue
			}

			currentArea := area(a, b)

			if currentArea > bestArea {
				bestArea = currentArea
				ba, bb = a, b
			}
		}
	}

	return bestArea, ba, bb
}

func ParsePoint(line string) Point {
	data := strings.Split(line, ",")
	x, _ := strconv.Atoi(data[0])
	y, _ := strconv.Atoi(data[1])
	return Point{x, y}
}

func PartOne(lines []string) int {
	var corners []Point
	for _, line := range lines {
		corners = append(corners, ParsePoint(line))
	}

	var bestArea int
	for i := 0; i < len(corners); i++ {
		for j := i + 1; j < len(corners); j++ {
			currentArea := area(corners[i], corners[j])
			if currentArea > bestArea {
				bestArea = currentArea
			}
		}
	}
	return bestArea
}

func PartTwo(lines []string) int {
	var corners []Point
	for _, line := range lines {
		corners = append(corners, ParsePoint(line))
	}

	// 1) Build compressed coordinates.
	xs, ys, xIndex, yIndex := buildCompressedCoords(corners)
	w, h := len(xs), len(ys)

	// 2) Build walls in the compressed grid.
	walls := buildWallsCompressed(corners, xIndex, yIndex)

	// 3) Flood fill the outside, starting from a guaranteed outside cell.
	start := Cell{
		i: xIndex[xs[0]], // xs[0] is (minX-1).
		j: yIndex[ys[0]], // ys[0] is (minY-1).
	}
	outside := floodOutsideCompressed(walls, w, h, start)

	// 4) Build a 2D prefix sum over outside cells.
	_, ps := buildBadGridAndPrefix(outside, w, h)

	// 5) Find the best rectangle using O(1) rectangle checks.
	bestArea, _, _ := findBestRectangleCompressed(corners, xIndex, yIndex, ps)

	return bestArea
}
