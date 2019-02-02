// https://leetcode.com/problems/palindrome-number/

func reverse_number(num int) int {
  rnum := 0
  for num > 0 {
    rnum = rnum * 10 + (num % 10)
    num /= 10
  }
  return rnum
}

func isPalindrome(num int) bool {
  reverse := reverse_number(num)
  if reverse == num { return true }
  return false
}

