// https://leetcode.com/problems/excel-sheet-column-title/
package main

func convertToTitle(num int) (ret string) {
	for ; num > 0; num /= 26 {
		num--
		ret = string(num%26+'A') + ret
	}
	return ret
}

func main() {}
