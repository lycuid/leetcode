// https://leetcode.com/problems/three-consecutive-odds/
package main

func threeConsecutiveOdds(arr []int) bool {
	for i := 2; i < len(arr); i++ {
		if arr[i-2]&1+arr[i-1]&1+arr[i]&1 == 3 {
			return true
		}
	}
	return false
}

func main() {}
