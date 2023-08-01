package main

import (
  "fmt"
  "time"
)

// Написать программу, которая конкурентно рассчитает 
// значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
func t2() {
  // вразнобой
  arr := []int{2,4,6,8,10}
  for _, val := range arr {
    go func(v int) {
      fmt.Println(v*v)
    }(val)
  }
  time.Sleep(2 * time.Second)
}
