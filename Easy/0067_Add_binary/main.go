// https://leetcode.com/problems/add-binary/
package main

func addBinary(a string, b string) string {
	n, m, i, j, carry := len(a)-1, len(b)-1, 0, 0, false
	ret := make([]byte, n+m+2)
	if n > m {
		n, m, a, b = m, n, b, a
	}
	for num := 0; i <= n; i++ {
		if num = int(a[n-i] + b[m-i] - '0' - '0'); carry {
			num++
		}
		ret[j], j, carry = byte(num%2)+'0', j+1, num >= 2
	}
	for num := 0; i <= m; i++ {
		if num = int(b[m-i] - '0'); carry {
			num++
		}
		ret[j], j, carry = byte(num%2)+'0', j+1, num >= 2
	}
	if carry {
		ret[j], j = '1', j+1
	}
	for i, n = 0, j; i < n/2; i++ {
		ret[i], ret[n-1-i] = ret[n-1-i], ret[i]
	}
	return string(ret[:j])
}

func main() {}
