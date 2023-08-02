package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(4 + 16 + 36….) с использованием конкурентных вычислений.
func t3() {
	arr := []int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, val := range arr {
		go func(v int) {
			sum += v * v
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println("T3 =", sum)
}
