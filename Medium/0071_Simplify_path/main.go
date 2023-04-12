// https://leetcode.com/problems/simplify-path/
package main

type PathParser struct{ buffer string }

func (this PathParser) Done() bool { return len(this.buffer) == 0 }

func (this *PathParser) Next() (ret string) {
	i, n := 0, len(this.buffer)-1
	for this.buffer = this.buffer[1:]; i < n && this.buffer[i] != '/'; i++ {
	}
	ret, this.buffer = this.buffer[:i], this.buffer[i:]
	return ret
}

func (this PathParser) Join(stack []string) (ret string) {
	for _, blob := range stack {
		ret += "/" + blob
	}
	if len(ret) == 0 {
		ret = "/"
	}
	return ret
}

func simplifyPath(path string) string {
	path_parser, stack := PathParser{path}, []string{}
	for !path_parser.Done() {
		switch token := path_parser.Next(); token {
		case ".":
			break
		case "..":
			if n := len(stack); n > 0 {
				stack = stack[:n-1]
			}
		default:
			if len(token) > 0 {
				stack = append(stack, token)
			}
		}
	}
	return path_parser.Join(stack)
}

func main() {}
