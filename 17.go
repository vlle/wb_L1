package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func binSearch(arr []int, search int) int {
  lo := 0
  hi := len(arr)-1
  for lo <= hi {
    mid := (hi + lo) / 2
    if arr[mid] < search {
      lo = mid + 1
    } else if arr[mid] > search {
      hi = mid - 1
    } else {
      return mid
    }
  }
  return -1
}

func t17() {
  arr := []int{-10,4,5,123,515,10004,144415}
  fmt.Printf("Index of %d in arr: %d", -10, binSearch(arr, -10))
}
