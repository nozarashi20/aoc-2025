package d11

import (
	"strings"
)

type Node struct {
	name    string
	outputs []*Node
}

type Graph map[string]*Node

func BuildGraph(lines []string) Graph {
	g := make(Graph)

	getNode := func(name string) *Node {
		if node, ok := g[name]; ok {
			return node
		}
		n := &Node{name: name}
		g[name] = n
		return n
	}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		sourceNode := getNode(strings.TrimSpace(parts[0]))
		for _, target := range strings.Fields(parts[1]) {
			targetNode := getNode(strings.TrimSpace(target))
			sourceNode.outputs = append(sourceNode.outputs, targetNode)
		}
	}

	return g
}

func CountAllPaths(g Graph, startName, targetName string) int {
	start, ok := g[startName]
	if !ok {
		return 0
	}

	var res int
	var dfs func(n *Node)

	dfs = func(n *Node) {
		if n.name == targetName {
			res++
		} else {
			for _, out := range n.outputs {
				dfs(out)
			}
		}
	}

	dfs(start)
	return res
}

func CountAllPathsVisiting(g Graph, startName, targetName string, mustNames []string) int64 {
	start, ok := g[startName]
	if !ok {
		return 0
	}
	target, ok := g[targetName]
	if !ok {
		return 0
	}

	idx := make(map[string]int, len(mustNames))
	for i, name := range mustNames {
		idx[name] = i
	}

	allMask := 0
	if len(mustNames) > 0 {
		allMask = (1 << len(mustNames)) - 1
	}

	// memo[node][mask] = number of paths from node to target
	memo := make(map[*Node]map[int]int64)

	var dfs func(n *Node, mask int) int64
	dfs = func(n *Node, mask int) int64 {
		if bit, ok := idx[n.name]; ok {
			mask |= 1 << bit
		}

		if m, ok := memo[n]; ok {
			if v, ok2 := m[mask]; ok2 {
				return v
			}
		} else {
			memo[n] = make(map[int]int64)
		}

		// base case
		if n == target {
			if mask == allMask {
				memo[n][mask] = 1
			} else {
				memo[n][mask] = 0
			}
			return memo[n][mask]
		}

		var total int64
		for _, out := range n.outputs {
			total += dfs(out, mask)
		}

		memo[n][mask] = total
		return total
	}

	return dfs(start, 0)
}

func PartOne(lines []string) int {
	g := BuildGraph(lines)
	return CountAllPaths(g, "you", "out")
}

func PartTwo(lines []string) int64 {
	g := BuildGraph(lines)
	return CountAllPathsVisiting(g, "svr", "out", []string{"dac", "fft"})
}
