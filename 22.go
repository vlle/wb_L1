package main

import "fmt"

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.


func t22(a, b int64) {
  fmt.Printf("a = %d\n b = %d\n", a, b)
  fmt.Printf("a - b = %d\n", a - b)
  fmt.Printf("a + b = %d\n", a + b)
  fmt.Printf("a * b = %d\n", a * b)
  fmt.Printf("a / b = %d\n", a / b)
}
