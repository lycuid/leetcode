// https://leetcode.com/problems/pascals-triangle/submissions/
func solve(n int, result [][]int) {
  if n == 0 { return }
  if n == 1 {
    result[0] = []int{1}
    return
  }
  
  solve(n - 1, result)
  tmp := make([]int, n)
  tmp[0] = 1;
  tmp[n - 1] = 1;
  
  previous := result[n - 2]
  for i := 1; i < (n - 1); i++ {
    tmp[i] = previous[i] + previous[i - 1]
  }

  result[n - 1] = tmp
}

func generate(numRows int) [][]int {
  result := make([][]int, numRows)
  solve(numRows, result)
  return result
}

func main() { }
