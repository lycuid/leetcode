// https://leetcode.com/problems/keys-and-rooms/
package main

func canVisitAllRooms(rooms [][]int) bool {
	ret, stack, visited := true, []int{0}, make([]bool, len(rooms))
	for visited[0] = true; len(stack) > 0; {
		for _, room := range rooms[stack[0]] {
			if !visited[room] {
				stack, visited[room] = append(stack, room), true
			}
		}
		stack = stack[1:]
	}
	for i := range visited {
		ret = ret && visited[i]
	}
	return ret
}

func main() {}
