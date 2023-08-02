package main

import (
	"context"
	"fmt"
  // "time"
)

// Реализовать все возможные способы остановки выполнения горутины.
func t6() {
	// Первый способ: горутина заканчивается сама.

	go func() {
		for i := 5; i > 0; i-- {
			fmt.Printf("1: У меня осталось %d действие\n", i)
		}
	}()

	// Второй способ: горутину можно прекратить через подачу сигнала в канал (если горутина его слушает)
	done := make(chan bool)
	go func(done chan bool) {
		for {
			select {
			case <-done:
				fmt.Println("2: Another goroutine was finished")
				break
			default:
				continue
			}
		}
	}(done)
	done <- true
	// Третий способ: Закончить горутину через close канала
	close_channel := make(chan int)
	go func(close_channel chan int) {
		for {
			v, ok := <-close_channel
			if !ok {
				return
			}
			fmt.Printf("3: Hello, world! I will be closed via closed channel in %d\n", v)
		}
	}(close_channel)
	close_channel <- 3
	close_channel <- 2
	close_channel <- 1
	close(close_channel)

	// Четвертый способ: Закрытие через контекст
  print := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context, print chan bool) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("4: I was finished too!")
				return
      case <-print:
        fmt.Println("4: Waiting to be finished")
			default:
				continue
			}
		}
	}(ctx, print)
  print<-true
  print<-true
  print<-true
	cancel()
}
