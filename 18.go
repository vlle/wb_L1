package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

type counter struct {
  mu sync.Mutex
  val int
};

func t18() {
  var wg sync.WaitGroup
  wg.Add(5)
  v := new(counter)
  for j := 0; j < 5; j++ {
    go func(c *counter) {
      for i:=0; i < 100; i++ {
        c.mu.Lock()
        c.val++
        c.mu.Unlock()
      }
      wg.Done()
    }(v)
  }
  wg.Wait()
  fmt.Println("18: counter value =", v.val)
}
