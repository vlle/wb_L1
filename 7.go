package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать конкурентную запись данных в map.
func t7() {
	type concurr_map struct {
		mu sync.Mutex
		m  map[int]string
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
