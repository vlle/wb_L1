package main

import (
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//
// Например:
// abcd — true
// abCdefAaf — false
// 	aabcd — false

func t26(s string) bool {
	v := strings.ToLower(s)
	m := make(map[rune]bool)
	for _, val := range v {
		if _, ok := m[val]; ok == true {
			return false
		}
		m[val] = true
	}
	return true
}
