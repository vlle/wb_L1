package main

import (
	"fmt"
	"strings"
	"time"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

func createHugeString(l int) string {
	var s strings.Builder
	for i := 0; i < l; i++ {
		s.WriteRune('a')
	}
	return s.String()
}

var justString string

// Негативные последствия:
// 1) Глобальная переменная
// 2) Аллокация памяти для переменной v
// 3) Копирование результата в итоговую строку вместо присвоения по указателю

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func fixedSomeFunc(s *string) {
	*s = createHugeString(1 << 10)[:100]
}

func t15() {
	// original implementation
	f := time.Now().UnixMicro()
	someFunc()
	s := time.Now().UnixMicro()
	fmt.Println("original implementation\ntime passed in ms:", s-f)
	// fixed implementation
	var str string
	f = time.Now().UnixMicro()
	fixedSomeFunc(&str)
	s = time.Now().UnixMicro()
	fmt.Println("fixed implementation\ntime passed in ms:", s-f)
}
