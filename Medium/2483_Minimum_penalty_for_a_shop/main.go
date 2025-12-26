// https://leetcode.com/problems/minimum-penalty-for-a-shop/
package main

func bestClosingTime(customers string) (hour int) {
	right, left, lowest := make([]int, len(customers)+1), 0, len(customers)
	for i := len(customers) - 1; i >= 0; i-- {
		if right[i] = right[i+1]; customers[i] == 'Y' {
			right[i]++
		}
	}
	customers += " "
	for i, val := range customers {
		if penalty := left + right[i]; penalty < lowest {
			lowest, hour = penalty, i
		}
		if val == 'N' {
			left++
		}
	}
	return hour
}

func main() {}
