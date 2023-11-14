// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
package main

func countPalindromicSubsequence(s string) (ret int) {
	var left, right, used [26]int
	for i := 0; i < len(s); i++ {
		right[s[i]-'a']++
	}
	for index := 0; index < len(s); index++ {
		i := s[index] - 'a'
		right[i]--
		for j := 0; j < 26; j++ {
			if left[j] > 0 && right[j] > 0 && used[i]&(1<<j) == 0 {
				ret, used[i] = ret+1, used[i]|(1<<j)
			}
		}
		left[i]++
	}
	return ret
}

func main() {}
