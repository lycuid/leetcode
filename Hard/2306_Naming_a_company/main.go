// https://leetcode.com/problems/naming-a-company/
package main

func distinctNames(ideas []string) (ret int64) {
	var head [26]map[string]bool
	for i := range head {
		head[i] = make(map[string]bool)
	}
	for i := range ideas {
		head[ideas[i][0]-'a'][ideas[i][1:]] = true
	}
	for i := range head {
		for j := i + 1; j < len(head); j++ {
			var count1, count2 int
			for taili := range head[i] {
				if !head[j][taili] {
					count1++
				}
			}
			for tailj := range head[j] {
				if !head[i][tailj] {
					count2++
				}
			}
			ret += int64(count1 * count2)
		}
	}
	return ret * 2
}

func main() {}
