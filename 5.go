package main

import (
  "fmt"
  "time"
)

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

