// https://leetcode.com/problems/process-string-with-special-operations-ii/
package main

func processStr(s string, k int64) byte {
	var size int64
	for i := range s {
		switch ch := s[i]; ch {
		case '*':
			size = max(0, size-1)
		case '#':
			size += size
		case '%':
		default:
			size++
		}
	}
	if k < size {
		for i := len(s) - 1; i >= 0; i-- {
			switch ch := s[i]; ch {
			case '*':
				size++
			case '#':
				size /= 2
				k %= size
			case '%':
				k = size - 1 - k
			default:
				if size-1 == k {
					return ch
				}
				size--
			}
		}
	}
	return '.'
}

func main() {}
