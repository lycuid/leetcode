// https://leetcode.com/problems/process-string-with-special-operations-i/
package main

func processStr(s string) string {
	var res []byte
	for i := range s {
		switch ch := s[i]; ch {
		case '*':
			res = res[:max(0, len(res)-1)]
		case '#':
			res = append(res, res...)
		case '%':
			for i, n := 0, len(res); i < n/2; i++ {
				res[i], res[n-1-i] = res[n-1-i], res[i]
			}
		default:
			res = append(res, ch)
		}
	}
	return string(res)
}

func main() {}
