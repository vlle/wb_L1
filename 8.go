package main

import (
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func t8(num int64, i int) {
	// Переворот бита
	and := num ^ (1 << i)
	fmt.Println(and)
}

func t8_2(num int64, i int) {
	// Переключение бита в 1
	and := num | (1 << i)
	fmt.Println(and)
}

func t8_3(num int64, i int) {
	// Переключение бита в 0
	and := num ^ (0 << i)
	fmt.Println(and)
}
