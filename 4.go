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
	// Создаем пул воркеров в размере N
	for i := 0; i < n; i++ {
		go func(i int) {
			for {
				select {
				// Слушаем: есть ли сигнал закрытого канала?
				case <-done:
					return
				case v := <-ch:
					fmt.Printf("Goroutine number %d, text: %s\n", i, string(v))
				default:
					continue
				}
			}
		}(i)
	}
	closed := false
	// Запускаем горутину которая будет ожидать команды блокировки и передавать эту блокировку всем горутинам
	go func() {
		<-done
		fmt.Println("blocked.")
		close(done)
		closed = true
	}()
	world := "Hello, world!"
	// Печчатаем данные
	for i := range world {
		if closed {
			break
		}
		for j := 0; j < 1000 && !closed; j++ {
			ch <- rune(world[i])
		}
	}
}
