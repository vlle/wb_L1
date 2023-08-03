package main

import "fmt"

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func t20(s string) {
	fmt.Println("20: Before\n", s)
	words := []string{}
	start_word := 0
	// Разбиваем строчку на массив слов
	for i := range s {
		if s[i] == ' ' || i == len(s)-1 {
			if i == len(s)-1 {
				i = len(s)
			}
			words = append(words, (s[start_word:i]))
			start_word = i + 1
		}
	}
	p2 := len(words) - 1
	j := 0
	// Переворачиваем слова
	for j < p2 {
		words[j], words[p2] = words[p2], words[j]
		j++
		p2--
	}
	res := ""
	// Собираем слова в строчку
	for i, word := range words {
		res += word
		if i != len(words)-1 {
			res += " "
		}
	}
	fmt.Println("20: After\n", res)
}
