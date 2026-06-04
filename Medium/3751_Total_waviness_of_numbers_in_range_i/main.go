// https://leetcode.com/problems/total-waviness-of-numbers-in-range-i/
package main

func waves(num int) (count int) {
	if num > 100 {
		fst, snd := num%10, (num/10)%10
		for num /= 100; num > 0; num /= 10 {
			trd := num % 10
			if (fst > snd && trd > snd) || (fst < snd && trd < snd) {
				count++
			}
			fst, snd = snd, trd
		}
	}
	return count
}

func totalWaviness(num1 int, num2 int) (count int) {
	for ; num1 <= num2; num1++ {
		count += waves(num1)
	}
	return count
}

func main() {}
