// https://leetcode.com/problems/fruit-into-baskets/
package main

func totalFruit(fruits []int) (ret int) {
	freq, count, sum := make(map[int]int), 0, 0
	for i, j := 0, 0; j < len(fruits); j++ {
		if freq[fruits[j]]++; freq[fruits[j]] == 1 {
			count++
		}
		for sum++; count > 2; sum, i = sum-1, i+1 {
			if freq[fruits[i]]--; freq[fruits[i]] == 0 {
				count--
			}
		}
		if ret < sum {
			ret = sum
		}
	}
	return ret
}

func main() {}
