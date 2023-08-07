package main

// Удалить i-ый элемент из слайса

// 2 способа:
// 1) изменить оригинальный слайс
// 2) вернуть новый слайс

func t23(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}
	if i == len(slice)-1 {
		return slice[:i]
	}
	return append(slice[:i], slice[i+1:]...)
}

func t23_1(slice *[]int, i int) {
	if i < 0 || i >= len(*slice) {
		return
	}
	if i == len(*slice)-1 {
		*slice = (*slice)[:i]
		return
	}
	*slice = append((*slice)[:i], ((*slice)[i+1:])...)
}
