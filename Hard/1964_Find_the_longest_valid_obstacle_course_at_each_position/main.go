// https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/
package main

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	cache, ret := []int{0}, make([]int, len(obstacles))
	for i, num := range obstacles {
		if n := len(cache); num >= cache[n-1] {
			cache, ret[i] = append(cache, num), n
		} else {
			l, r := 0, n-1
			for l < r {
				if mid := (l + r) / 2; cache[mid] > num {
					r = mid
				} else {
					l = mid + 1
				}
			}
			cache[l], ret[i] = num, l
		}
	}
	return ret
}

func main() {}
