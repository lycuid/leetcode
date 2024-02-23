// https://leetcode.com/problems/cheapest-flights-within-k-stops/
package main

type Tuple struct{ city, price, stops int }
type PriorityQueue struct{ inner []Tuple }

func MakePQ() PriorityQueue          { return PriorityQueue{make([]Tuple, 1)} }
func (pq PriorityQueue) Empty() bool { return len(pq.inner) == 1 }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq.inner[i].price == pq.inner[j].price {
		return pq.inner[i].stops < pq.inner[j].stops
	}
	return pq.inner[i].price < pq.inner[j].price
}

func (pq *PriorityQueue) Push(item Tuple) {
	pq.inner = append(pq.inner, item)
	for i := len(pq.inner) - 1; i > 1 && pq.Less(i, i/2); i /= 2 {
		pq.inner[i], pq.inner[i/2] = pq.inner[i/2], pq.inner[i]
	}
}

func (pq *PriorityQueue) Pop() Tuple {
	item, n := pq.inner[1], len(pq.inner)-1
	pq.inner[1], pq.inner = pq.inner[n], pq.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && pq.Less(j+1, j) {
			j++
		}
		if pq.Less(i, j) {
			break
		}
		pq.inner[i], pq.inner[j] = pq.inner[j], pq.inner[i]
	}
	return item
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adj, pq, stops := make([][]Tuple, n+1), MakePQ(), make([]int, n+1)
	for i := range stops {
		stops[i] = k + 1
	}
	for _, flight := range flights {
		adj[flight[0]] = append(adj[flight[0]], Tuple{flight[1], flight[2], 1})
	}
	flight := Tuple{src, 0, 0}
	for pq.Push(flight); !pq.Empty(); {
		if flight = pq.Pop(); flight.city == dst {
			return flight.price
		}
		if flight.stops < stops[flight.city] && flight.stops <= k {
			stops[flight.city] = flight.stops + 1
			for _, new_flight := range adj[flight.city] {
				pq.Push(Tuple{
					new_flight.city,
					flight.price + new_flight.price,
					stops[flight.city],
				})
			}
		}
	}
	return -1
}

func main() {}
