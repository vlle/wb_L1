package main

import (
	"fmt"
	"os/exec"
)

// Реализовать собственную функцию sleep

// Реализую функцию через системные вызовы

func sleep(duration int) {
  fmt.Printf("I will sleep for %d second\n", duration)
  c := exec.Command("sleep", fmt.Sprint(duration))
  c.Run()
}

func t25() {
  for i := 0; i < 100; i++ {
    go func() {
      sleep(4)
      fmt.Println("I have awakened")
    }()
  }
  sleep(5)
}
