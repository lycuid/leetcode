// https://leetcode.com/problems/number-of-senior-citizens/
package main

import "strconv"

func countSeniors(details []string) (count int) {
	for _, detail := range details {
		if age, _ := strconv.Atoi(detail[11:13]); age > 60 {
			count++
		}
	}
	return count
}

func main() {}
