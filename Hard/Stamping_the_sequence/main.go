// https://leetcode.com/problems/stamping-the-sequence/
package main

func Repeat(n int, ch string) (ret string) {
	for i := 0; i < n; i++ {
		ret += ch
	}
	return ret
}

func Substring(str, sub []byte) int {
	if n := len(sub); len(str) >= n {
	OUTTER:
		for i := 0; i <= len(str)-n; i++ {
			for j, k := 0, n-1; j <= k; j, k = j+1, k-1 {
				if str[i+j] != sub[j] || str[i+k] != sub[k] {
					continue OUTTER
				}
			}
			return i
		}
	}
	return -1
}

func movesToStamp(stamp string, _target string) (ret []int) {
	var masks [][]byte
	target := []byte(_target)
	for i, n := 1, len(stamp); i <= n; i++ {
		for j := 0; j+i <= n; j++ {
			mask := Repeat(j, "?") + stamp[j:j+i] + Repeat(n-j-i, "?")
			masks = append(masks, []byte(mask))
		}
	}
	for _, mask := range masks {
		fmt.Println(string(mask))
	}
	for found := true; found; found = !found {
		for _, mask := range masks {
			if index := Substring(target, mask); index != -1 {
				fmt.Println(string(target), string(mask), index)
				ret, found = append(ret, index), !found
				for i := index; i < index+len(mask); i++ {
					target[i] = '?'
				}
				break
			}
		}
	}
	for _, ch := range target {
		if ch != '?' {
			return nil
		}
	}
	if len(ret) >= 10*len(target) {
		return nil
	}
	for i, l := 0, len(ret); i < l/2; i++ {
		ret[i], ret[l-i-1] = ret[l-i-1], ret[i]
	}
	return ret
}

func main() {}
