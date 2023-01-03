// https://leetcode.com/problems/delete-columns-to-make-sorted/
package main

func minDeletionSize(strs []string) (ret int) {
	for x := 0; x < len(strs[0]); x++ {
		for y := 1; y < len(strs); y++ {
			if strs[y][x] < strs[y-1][x] {
				ret++
				break
			}
		}
	}
	return ret
}

func main() {}
