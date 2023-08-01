package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в
// первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.


func t9() {
  first_ch := make(chan int)
  second_ch := make(chan int)
  done := make(chan bool)
  nums := []int{1, 2, 3, 4, 5, 6}

  // Первая часть конвеера: принимает число и отдает его в канал удвоенным
  go func(c1, c2 chan int) {
    for {
      v := <- c1
      c2 <- v * 2 
    }
  }(first_ch, second_ch)

  // Вторая часть конвеера: принтует удвоенное число из второго канала
  go func(c2 chan int, l int) {
    for i := 0; i < l; i++ {
      fmt.Println(<-c2)
    }
    done <- true
  }(second_ch, len(nums))

  // Запуск конвеера 
  for i := range nums {
    first_ch <- nums[i]
  }
  // блокируем дальнейшую работу мейна и ждем конца работы второй горутины
  <-done
}
