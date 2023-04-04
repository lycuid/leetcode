// https://leetcode.com/problems/optimal-partition-of-string/
package main

func partitionString(s string) (snaps int) {
	var found, bit uint32
	for _, ch := range s {
		if bit = 1 << (ch - 'a'); found&bit != 0 {
			found, snaps = 0, snaps+1
		}
		found |= bit
	}
	return snaps + 1
}

func main() {}
