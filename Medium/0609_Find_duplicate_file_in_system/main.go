// https://leetcode.com/problems/find-duplicate-file-in-system/
package main

import "fmt"

func findDuplicate(paths []string) (ret [][]string) {
	cache := make(map[string][]string)
	for _, path := range paths {
		var dir, tmp string
		fmt.Sscanf(path, "%s", &dir)
		path = path[len(dir):]
		dir += "/"
		for i, e := fmt.Sscanf(path, "%s", &tmp); e == nil; i, e = fmt.Sscanf(path, "%s", &tmp) {
			var file, content string
			for i, file = 0, dir; i < len(tmp) && tmp[i] != '('; i++ {
				file += string(tmp[i])
			}
			for i++; i < len(tmp) && tmp[i] != ')'; i++ {
				content += string(tmp[i])
			}
			path, cache[content] = path[len(tmp)+1:], append(cache[content], file)
		}
	}
	for _, files := range cache {
		if len(files) > 1 {
			ret = append(ret, files)
		}
	}
	return ret
}

func main() {}
