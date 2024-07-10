// https://leetcode.com/problems/crawler-log-folder/
package main

func minOperations(logs []string) (count int) {
	for _, log := range logs {
		switch log {
		case "./":
			break
		case "../":
			count = max(0, count-1)
		default:
			count++
		}
	}
	return count
}

func main() {}
