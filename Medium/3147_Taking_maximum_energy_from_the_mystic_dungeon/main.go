// https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/
package main

func maximumEnergy(energy []int, k int) int {
	var (
		n   = len(energy)
		ret = energy[n-1]
	)
	for i := n - 1; i >= 0; i-- {
		if j := i + k; j < n {
			energy[i] += energy[j]
		}
		ret = max(ret, energy[i])
	}
	return ret
}

func main() {}
