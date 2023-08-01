package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

// Убираем знаки после запятой и последнюю цифру. То есть округляем до шага в 10 градусов.
func round(n float32) int {
  n_int := int(n)
  return n_int - n_int % 10 
}

func t10() {
  temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
  groups := make(map[int][]float32)
  for i := range temperatures {
    category := round(temperatures[i])
    // Если категории нет — создаем слайс
    if _, ok := groups[category]; !ok {
      groups[category] = make([]float32, 0)
    }
    groups[category] = append(groups[category], temperatures[i])
  }
  fmt.Println(groups)
}
