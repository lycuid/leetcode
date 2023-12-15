// https://leetcode.com/problems/destination-city/
package main

func destCity(paths [][]string) string {
	cache := make(map[string]bool)
	for _, path := range paths {
		cache[path[0]] = true
	}
	for _, path := range paths {
		if !cache[path[1]] {
			return path[1]
		}
		delete(cache, path[1])
	}
	return ""
}

func main() {}
