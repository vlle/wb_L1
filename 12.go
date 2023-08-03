package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func t12() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	unique_vals := make(map[string]bool)
	unique_slice := make([]string, 0)
	for i := range strings {
		// Добавляем уникальные значения в мапу если их там нет и сразу в итоговый слайс
		if _, ok := unique_vals[strings[i]]; !ok {
			unique_vals[strings[i]] = true
			unique_slice = append(unique_slice, strings[i])
		}
	}
	for i := range unique_slice {
		fmt.Println(unique_slice[i])
	}
}
