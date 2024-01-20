// https://leetcode.com/problems/sum-of-subarray-minimums/
package main

func sumSubarrayMins(arr []int) (sum int) {
	stack, cache := []int{}, make([]int, len(arr))
	for i := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		if n := len(stack) - 1; n >= 0 {
			cache[i] = (i-stack[n])*arr[i] + cache[stack[n]]
		} else {
			cache[i] = (i + 1) * arr[i]
		}
		stack, cache[i] = append(stack, i), cache[i]%(1e9+7)
	}
	for i := range cache {
		sum = (sum + cache[i]) % (1e9 + 7)
	}
	return sum
}

func main() {}
