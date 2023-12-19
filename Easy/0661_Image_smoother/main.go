// https://leetcode.com/problems/image-smoother/
package main

func imageSmoother(img [][]int) [][]int {
	new_img := make([][]int, len(img))
	for i := range new_img {
		new_img[i] = make([]int, len(img[0]))
		for j := range new_img[i] {
			new_img[i][j] = img[i][j]
		}
	}

	for i := range new_img {
		for j := range new_img[i] {
			cells := 1
			if i > 0 {
				new_img[i][j], cells = new_img[i][j]+img[i-1][j], cells+1
			}
			if i < len(new_img)-1 {
				new_img[i][j], cells = new_img[i][j]+img[i+1][j], cells+1
			}
			if j > 0 {
				new_img[i][j], cells = new_img[i][j]+img[i][j-1], cells+1
			}
			if j < len(new_img[i])-1 {
				new_img[i][j], cells = new_img[i][j]+img[i][j+1], cells+1
			}
			if i > 0 && j > 0 {
				new_img[i][j], cells = new_img[i][j]+img[i-1][j-1], cells+1
			}
			if i > 0 && j < len(new_img[i])-1 {
				new_img[i][j], cells = new_img[i][j]+img[i-1][j+1], cells+1
			}
			if i < len(new_img)-1 && j > 0 {
				new_img[i][j], cells = new_img[i][j]+img[i+1][j-1], cells+1
			}
			if i < len(new_img)-1 && j < len(new_img[i])-1 {
				new_img[i][j], cells = new_img[i][j]+img[i+1][j+1], cells+1
			}
			new_img[i][j] /= cells
		}
	}
	return new_img
}

func main() {}
