package day23

import (
	"slices"
	"sort"
	"strings"
)

func Run(input string) (int, string) {
	ans1 := 0
	ans2 := ""

	lines := strings.Fields(input)
	graph := make(map[string][]string)
	for _, line := range lines {
		split := strings.Split(line, "-")
		graph[split[0]] = append(graph[split[0]], split[1])
		graph[split[1]] = append(graph[split[1]], split[0])
	}

	ans1 = solvePartOne(graph)
	ans2 = solvePartTwo(graph)

	return ans1, ans2
}

func solvePartOne(graph map[string][]string) int {
	found := make(map[string][]string)
	for node, connected := range graph {
		if !strings.HasPrefix(node, "t") {
			continue
		}
		for _, connectedNode := range connected {
			for _, maybeConnectedNode := range graph[connectedNode] {
				if slices.Contains(graph[maybeConnectedNode], node) {
					threeConnectedNodes := []string{node, connectedNode, maybeConnectedNode}
					sort.Strings(threeConnectedNodes)
					found[strings.Join(threeConnectedNodes, ",")] = threeConnectedNodes
				}
			}

		}
	}
	return len(found)
}

func solvePartTwo(graph map[string][]string) string {
	maxFound := make([]string, 0)
	c := make(chan []string)
	for nodeId := range graph {
		go calculateLargestConnectedGroup(&graph, nodeId, c)
	}
	for range graph {
		found := <-c
		if len(found) > len(maxFound) {
			maxFound = found
		}
	}
	sort.Strings(maxFound)
	return strings.Join(maxFound, ",")
}

func calculateLargestConnectedGroup(graph *map[string][]string, node string, c chan []string) {
	connections := (*graph)[node]
	root := TreeNode{0, 0, node, 0}
	tree := make([]TreeNode, 1)
	tree[0] = root
	idx := 1
	for _, connectedNode := range connections {
		for _, branch := range tree {
			allParentsMatch := true
			parent := branch
			for parent != root {
				if !slices.Contains((*graph)[connectedNode], parent.id) {
					allParentsMatch = false
					break
				}
				parent = tree[parent.parent]
			}
			if allParentsMatch {
				tree = append(tree, TreeNode{idx, branch.idx, connectedNode, branch.level + 1})
				idx++
			}
		}
	}
	maxElem := root
	for _, treeNode := range tree {
		if treeNode.level > maxElem.level {
			maxElem = treeNode
		}
	}

	ret := make([]string, maxElem.level+1)
	for maxElem != root {
		ret[maxElem.level] = maxElem.id
		maxElem = tree[maxElem.parent]
	}
	ret[0] = maxElem.id
	c <- ret
}

type TreeNode struct {
	idx    int
	parent int
	id     string
	level  int
}
