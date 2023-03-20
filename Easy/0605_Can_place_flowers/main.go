// https://leetcode.com/problems/can-place-flowers/
package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append(append([]int{0}, flowerbed...), 0)
	for i := 1; n > 0 && i < len(flowerbed)-1; i++ {
		if flowerbed[i-1]+flowerbed[i]+flowerbed[i+1] == 0 {
			i, n = i+1, n-1
		}
	}
	return n == 0
}

func main() {}
