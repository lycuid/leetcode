// https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
package main

func garbageCollection(garbage []string, travel []int) (ret int) {
	freq := make([][3]int, len(garbage))
	for i := range garbage {
		for j := range garbage[i] {
			switch garbage[i][j] {
			case 'G':
				freq[i][0]++
			case 'P':
				freq[i][1]++
			case 'M':
				freq[i][2]++
			}
		}
	}
	travel = append([]int{0}, travel...)
	for i := 1; i < len(travel); i++ {
		travel[i] += travel[i-1]
	}
	found := [3]bool{}
	for i := len(freq) - 1; i >= 0; i-- {
		if !found[0] && freq[i][0] > 0 {
			found[0], ret = true, ret+travel[i]
		}
		if !found[1] && freq[i][1] > 0 {
			found[1], ret = true, ret+travel[i]
		}
		if !found[2] && freq[i][2] > 0 {
			found[2], ret = true, ret+travel[i]
		}
		for j := 0; j < 3; j++ {
			ret += freq[i][j]
		}
	}
	return ret
}

func main() {}
