// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
package main

func twoSum(numbers []int, target int) []int {
	index := []int{1, len(numbers)}
	for index[0] < index[1] {
		sum := numbers[index[0]-1] + numbers[index[1]-1]
		if sum < target {
			index[0]++
		} else if sum > target {
			index[1]--
		} else {
			break
		}
	}
	return index
}

func main() {}
