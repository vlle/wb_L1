package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
)

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
