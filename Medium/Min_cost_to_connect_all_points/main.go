// https://leetcode.com/problems/min-cost-to-connect-all-points/
package main

import "math"

type Point struct {
	x, y int
}

func NewPoint(x []int) Point {
	return Point{x[0], x[1]}
}

func (p1 Point) Edge(p2 Point) int {
	return Abs(p1.x-p2.x) + Abs(p1.y-p2.y)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetMin(edges map[Point]int) (p Point, ok bool) {
	min := math.MaxInt
	for point, edge := range edges {
		if edge < min {
			min, p, ok = edge, point, true
		}
	}
	return p, ok
}

func minCostConnectPoints(points [][]int) (sum int) {
	l := len(points)
	edges := make(map[Point]int)
	if l > 0 {
		edges[NewPoint(points[0])] = 0
	}
	for i := 1; i < l; i++ {
		edges[NewPoint(points[i])] = math.MaxInt
	}
	for old_point, ok := GetMin(edges); ok; old_point, ok = GetMin(edges) {
		sum += edges[old_point]
		delete(edges, old_point)
		for point, edge := range edges {
			edges[point] = Min(edge, point.Edge(old_point))
		}
	}
	return sum
}

func main() {}
