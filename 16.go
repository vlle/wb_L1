package main

import (
	"fmt"
	"math/rand"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quicksort(array []int) []int {
	// Выбираем элемент для опоры
	if len(array) <= 1 {
		return array
	}
	elem := array[rand.Intn(len(array))]
	// Делим массив
	lower := []int{}
	equal_or_higher := []int{}
	for i := range array {
		if array[i] < elem {
			lower = append(lower, array[i])
		} else {
			equal_or_higher = append(equal_or_higher, array[i])
		}
	}
	// Вызываем сортировку для поделенных частей массива
	return append(quicksort(lower), quicksort(equal_or_higher)...)
}

func t16() {
	arr := []int{4, 5, 2, 6, 10, 15, 0, -114}
	fmt.Println("Before sort:", arr)
	fmt.Println("After sort:", quicksort(arr))
}
