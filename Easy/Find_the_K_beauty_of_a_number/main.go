// https://leetcode.com/problems/find-the-k-beauty-of-a-number/
package main

func ToArray(num int) (arr []int) {
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}
	return arr
}

func ToNumber(start, k int, arr []int) (ret int) {
	for i := 0; i < k; i++ {
		ret *= 10
		ret += arr[i+start]
	}
	return ret
}

func divisorSubstrings(num int, k int) (count int) {
	arr := ToArray(num)
	for i := 0; i < len(arr)-k+1; i++ {
		n := ToNumber(i, k, arr)
		if n > 0 && num%n == 0 {
			count++
		}
	}
	return count
}

func main() {}
