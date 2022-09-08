// https://leetcode.com/problems/add-binary/

package main

func addBinary(a string, b string) string {
  c := []byte{}

  var carry byte = 0
  var sub byte = 48
  var base byte = 2

  i, j := len(a) - 1, len(b) - 1
  for i >= 0 && j >= 0 {
    sum := a[i] + b[j] - (2 * sub) + carry
    c = append([]byte{sum % base + sub}, c...)
    carry = sum / base
    i--
    j--
  }
  for i >= 0 {
    sum := a[i] - sub + carry
    c = append([]byte{sum % base + sub}, c...)
    carry = sum / base
    i--
  }
  for j >= 0 {
    sum := b[j] - sub + carry
    c = append([]byte{sum % base + sub}, c...)
    carry = sum / base
    j--
  }
  if carry > 0 {
    c = append([]byte{carry + 48}, c...)
  }
  return string(c)
}

func main() { }

