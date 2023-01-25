// https://leetcode.com/problems/find-closest-node-to-given-two-nodes/
package main

func closestMeetingNode(edges []int, node0, node1 int) (node int) {
	var dst [2][]int
	dst[0], dst[1], node = make([]int, len(edges)), make([]int, len(edges)), -1
	for i := range edges {
		dst[0][i], dst[1][i] = -1, -1
	}
	for d0, d1 := 0, 0; (node0 != -1 && dst[0][node0] == -1) || (node1 != -1 && dst[1][node1] == -1); {
		if node0 != -1 && dst[0][node0] == -1 {
			dst[0][node0], node0, d0 = d0, edges[node0], d0+1
		}
		if node1 != -1 && dst[1][node1] == -1 {
			dst[1][node1], node1, d1 = d1, edges[node1], d1+1
		}
	}
	for i, max := 0, len(edges)+1; i < len(edges); i++ {
		if max_d := dst[0][i]; max_d != -1 && dst[1][i] != -1 {
			if max_d < dst[1][i] {
				max_d = dst[1][i]
			}
			if max_d < max {
				node, max = i, max_d
			}
		}
	}
	return node
}

func main() {}
