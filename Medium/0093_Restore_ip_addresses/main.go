// https://leetcode.com/problems/restore-ip-addresses/
package main

func Valid(s string) bool {
	switch len(s) {
	case 1:
		return true
	case 2:
		return s[0] != '0'
	case 3:
		if s[0] == '2' {
			if s[1] == '5' {
				return s[2] <= '5'
			}
			return s[1] < '5'
		}
		return s[0] == '1'
	default:
		return false
	}
}

func Aux(n int, s string) (ret []string) {
	if n == 1 {
		if Valid(s) {
			ret = append(ret, s)
		}
		goto RETURN
	}
	for i := 0; i <= len(s)-n; i++ {
		if Valid(s[:i+1]) {
			for _, num := range Aux(n-1, s[i+1:]) {
				ret = append(ret, s[:i+1]+"."+num)
			}
		}
	}
RETURN:
	return ret
}

func restoreIpAddresses(s string) []string { return Aux(4, s) }

func main() {}
