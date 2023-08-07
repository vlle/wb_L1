package main

import "fmt"

// Поменять местами два числа без создания временной переменной
func t13(first, second int) {
	fmt.Println("Before:", first, second)
	// XOR: исключающее или
	// example:
	// initial state in bits:
	// first: 101
	// second: 011
	first ^= second
	// first: 110
	// second: 011
	second ^= first
	// first: 110
	// second: 101
	first ^= second
	// first: 110
	// second: 011
	fmt.Println("After:", first, second)

	// Или второй способ:
	first, second = second, first
	fmt.Println("Swap again:", first, second)
}
