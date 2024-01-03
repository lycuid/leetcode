// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/
package main

func numberOfBeams(bank []string) (count int) {
	var prev int
	for i := range bank {
		var curr int
		for _, ch := range bank[i] {
			curr += int(ch - '0')
		}
		if count += prev * curr; curr != 0 {
			prev = curr
		}
	}
	return count
}

func main() {}
