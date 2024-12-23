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

	connectedMap := make(map[string][]string)

	for _, line := range lines {
		split := strings.Split(line, "-")
		connectedMap[split[0]] = append(connectedMap[split[0]], split[1])
		connectedMap[split[1]] = append(connectedMap[split[1]], split[0])
	}

	//for _, e := range connectedMap {
	//	slices.Sort(e)
	//}

	foundClusters := make(map[string][]string)
	for k0, connected := range connectedMap {
		if !strings.HasPrefix(k0, "t") {
			continue
		}
		for _, k1 := range connected {
			if k1 == k0 {
				continue
			}
			for _, k2 := range connectedMap[k1] {
				if k1 == k2 || k2 == k0 {
					continue
				}
				if slices.Contains(connectedMap[k2], k0) {
					k := []string{k0, k1, k2}
					sort.Strings(k)
					foundClusters[strings.Join(k, ",")] = k
				}
			}

		}
	}
	ans1 = len(foundClusters)

	foundClusters = make(map[string][]string)
	for k0, connected := range connectedMap {
		for _, k1 := range connected {
			if k1 == k0 {
				continue
			}
			for _, k2 := range connectedMap[k1] {
				if k1 == k2 || k2 == k0 {
					continue
				}
				if slices.Contains(connectedMap[k2], k0) {
					k := []string{k0, k1, k2}
					sort.Strings(k)
					foundClusters[strings.Join(k, ",")] = k
				}
			}

		}
	}
	for key, _ := range connectedMap {
		for _, cluster := range foundClusters {
			if slices.Contains(cluster, key) {
				continue
			}
			match := true
			for _, node := range cluster {
				if !slices.Contains(connectedMap[node], key) {
					match = false
					break
				}
			}
			if match {

				newCluster := append(slices.Clone(cluster), key)
				sort.Strings(newCluster)
				foundClusters[strings.Join(newCluster, ",")] = newCluster
			}
		}
	}

	maxLen := 0
	for s, _ := range foundClusters {
		if len(s) > maxLen {
			ans2 = s
			maxLen = len(s)
		}
	}
	return ans1, ans2
}
