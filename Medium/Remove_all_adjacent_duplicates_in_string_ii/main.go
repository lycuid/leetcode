// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
package main

type StateMachine struct {
	state [1e5 + 1]struct {
		value byte
		count int
	}
	cursor int
}

func (sm *StateMachine) NextOffset(ch byte, k int) int {
	if sm.cursor > 0 && sm.state[sm.cursor-1].value == ch {
		// update previous state.
		sm.state[sm.cursor-1].count++
		if sm.state[sm.cursor-1].count == k {
			sm.cursor--
			return 1 - k
		}
	} else {
		// Add new state.
		sm.state[sm.cursor].value = ch
		sm.state[sm.cursor].count = 1
		sm.cursor++
	}
	return 1
}

func removeDuplicates(s string, k int) string {
	var (
		cursor int
		sm     StateMachine
	)
	buffer := make([]byte, len(s))
	for _, ch := range []byte(s) {
		buffer[cursor] = ch
		cursor += sm.NextOffset(ch, k)
	}
	return string(buffer[:cursor])
}

func main() {}
