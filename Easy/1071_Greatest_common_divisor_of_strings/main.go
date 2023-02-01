// https://leetcode.com/problems/greatest-common-divisor-of-strings/
package main

func gcdOfStrings(str1 string, str2 string) (ret string) {
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}
	for i, j, n, m := 1, 0, len(str1), len(str2); i <= len(str1); i++ {
		if str := str1[:i]; (n+m)%i == 0 {
			for j = 0; j+i <= m; j += i {
				if j+i <= n && str1[j:j+i] != str {
					break
				}
				if str2[j:j+i] != str {
					break
				}
			}
			if j == m && i > len(ret) {
				ret = str1[:i]
			}
		}
	}
	return ret
}

func main() {}
