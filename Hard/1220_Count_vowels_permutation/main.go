// https://leetcode.com/problems/count-vowels-permutation/
package main

const MOD = 1e9 + 7
const A, E, I, O, U = 0, 1, 2, 3, 4

func countVowelPermutation(n int) int {
	state, Prev, Cur := [2][5]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}}, 0, 1
	for ; n > 1; n-- {
		Prev, Cur = Cur, Prev
		state[Cur][A] = (state[Prev][E] + state[Prev][I] + state[Prev][U]) % MOD
		state[Cur][E] = (state[Prev][A] + state[Prev][I]) % MOD
		state[Cur][I] = (state[Prev][E] + state[Prev][O]) % MOD
		state[Cur][O] = (state[Prev][I]) % MOD
		state[Cur][U] = (state[Prev][I] + state[Prev][O]) % MOD
	}
	return (state[Cur][A] + state[Cur][E] + state[Cur][I] + state[Cur][O] + state[Cur][U]) % MOD
}

func main() {}
