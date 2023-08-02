package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств
func t11(f1 []int, f2 []int) {
  // Создаем мапу для того за О(1) в дальнейшем проверять пересечение
  intersection := make(map[int]bool) 
  for i := range f1 {
    intersection[f1[i]] = true
  }
  res := make([]int ,0)
  for i := range f2 {
    // Проверяем наличие пересечения и в случае успеха добавляем в результирующий слайс
    if (intersection[f2[i]] == true) {
      res = append(res, f2[i])
    }
  }
  for i := range res {
    fmt.Println(res[i])
  }
}
