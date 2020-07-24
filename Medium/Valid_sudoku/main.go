// https://leetcode.com/problems/valid-sudoku/

package main

import "strconv"

func check_rows(s [][]byte) bool {
  visited := make(map[int]bool)
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      if num, err := strconv.Atoi(string(s[i][j])); err == nil {
        if _, ok := visited[num]; ok {
          return false
        }
        visited[num] = true
      }
    }
    for k := range visited {
      delete(visited, k)
    }
  }
  return true
}

func check_columns(s [][]byte) bool {
  visited := make(map[int]bool)
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      if num, err := strconv.Atoi(string(s[j][i])); err == nil {
        if _, ok := visited[num]; ok {
          return false
        }
        visited[num] = true
      }
    }
    for k := range visited {
      delete(visited, k)
    }
  }
  return true
}

func check_boxes(s [][]byte) bool {
  visited := make(map[int]bool)
  for x := 0; x < 9; x+=3 {
    for y := 0; y < 9; y+=3 {
      for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
          if num, err := strconv.Atoi(string(s[j+y][i+x])); err == nil {
            if _, ok := visited[num]; ok {
              return false
            }
            visited[num] = true
          }
        }
      }
      for k := range visited {
        delete(visited, k)
      }
    }
  }
  return true
}

func isValidSudoku(s [][]byte) bool {
  return check_rows(s[:]) && check_columns(s[:]) && check_boxes(s[:])
}

func main() { }

