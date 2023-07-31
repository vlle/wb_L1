package main

import (
  "fmt"
  "time"
  "sync"
)

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {}

func (h *Human) SayHi() {
  fmt.Println("Hi")
}

type Action struct {
  Human
}

func t1() { 
  a := Action{}
  a.SayHi()
}

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

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(4 + 16 + 36….) с использованием конкурентных вычислений.
func t3() {
  arr := []int{2,4,6,8,10}
  sum := 0
  var wg sync.WaitGroup
  wg.Add(len(arr))
  for _, val := range arr {
    go func(v int) {
      sum += v*v
      wg.Done()
    }(val)
  }
  wg.Wait()
  fmt.Println("T3 =", sum)
}

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

// n = workers amount
func t4(n int) {
  ch := make(chan string)
  go func() {
    v := <-ch
    fmt.Println(v)
  }()
  ch <- "Hello, world!"
}


func main() {
  t1()
  t2()
  t3()
  t4(4)
}
