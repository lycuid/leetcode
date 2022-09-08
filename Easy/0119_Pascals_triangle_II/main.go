// https://leetcode.com/problems/pascals-triangle-ii/

func getRow(n int) []int {
  if n == 0 {
    return []int{1}
  }
  previous := getRow(n - 1)
  result := make([]int, n + 1)
  result[0] = 1
  result[n] = 1
  for i := 1; i < n; i++ {
    result[i] = previous[i] + previous[i - 1]
  }
  return result
}

func main() { }

