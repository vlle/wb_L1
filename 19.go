package main

import (
  "fmt"
)

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func t19(s string) {
  res := []rune(s)
  p1 := 0
  p2 := len(res)-1
  for p1 < p2 {
    res[p1], res[p2] = res[p2], res[p1]
    p1++
    p2--
  }
  fmt.Println(string(res))
}
