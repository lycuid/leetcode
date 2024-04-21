// https://leetcode.com/problems/find-if-path-exists-in-graph/
package main

type Queue struct {
	inner      []int
	start, end int
}

func MakeQ(capacity int) Queue    { return Queue{inner: make([]int, capacity+1)} }
func (q Queue) Empty() bool       { return q.start == q.end }
func (q *Queue) Enqueue(item int) { q.inner[q.end], q.end = item, (q.end+1)%len(q.inner) }
func (q *Queue) Dequeue() (item int) {
	item, q.start = q.inner[q.start], (q.start+1)%len(q.inner)
	return item
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	queue, done := MakeQ(n), make([]bool, n)
	done[source] = true
	for queue.Enqueue(source); !queue.Empty(); {
		node := queue.Dequeue()
		for _, child := range adj[node] {
			if child == destination {
				return true
			}
			if !done[child] {
				queue.Enqueue(child)
				done[child] = true
			}
		}
	}
	return false
}

func main() {}
