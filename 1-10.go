package main

import (
  "fmt"
  "time"
  "sync"
	"os"
	"os/signal"
	"syscall"
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

func t4(n int) {
  ch := make(chan rune)
	done := make(chan os.Signal, 1)
  signal.Notify(done, os.Interrupt, syscall.SIGINT)
  for i := 0; i < n; i++ {
    go func() {
      for {
        select {
        case <- done:
          return
        case v := <- ch:
          fmt.Println(string(v))
        default:
          continue
        }
      }
    }()
  }
  world :=  "Hello, world!"
  for i := range world {
    ch <- rune(world[i])
  }
}

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
func t5(n time.Duration) {
  data := "Hello, another world!"
  ch := make(chan string)
  go func() {
    for {
      v := <- ch
      fmt.Println(v)
    }
  }()
  for i := range data {
    ch <- string(data[i])
  }
  time.Sleep(n * time.Second)
}

// Реализовать все возможные способы остановки выполнения горутины. 
func t6() {

}

// Реализовать конкурентную запись данных в map.
func t7() {
  type concurr_map struct {
    mu sync.Mutex
    m map[int]string
  }
  var m concurr_map
  m.m = make(map[int]string)
  for i := 0; i < 10; i++ {
    go func(i int) {
      m.mu.Lock()
      m.m[i] = "hello"
      m.mu.Unlock()
    }(i)
  }
  time.Sleep(4 * time.Second)
  for i := 0; i < 10; i++ {
    fmt.Println(m.m[i])
  }
}


// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func t8() {

}


func main() {
  // t1()
  // t2()
  // t3()
  // t4(10)
  // t5(4)
  // t6()
  t7()
  // t8()
}
