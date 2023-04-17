// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ret, max := make([]bool, len(candies)), 0
	for i := range candies {
		if candies[i] > max {
			max = candies[i]
		}
	}
	for i := range candies {
		ret[i] = candies[i]+extraCandies >= max
	}
	return ret
}

func main() {}
