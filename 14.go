package main

import "fmt"

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
func t14(inc interface{}) {
  if _, ok := inc.(int); ok {
    fmt.Println("T14: This is int")
  } else if _, ok := inc.(string); ok {
    fmt.Println("T14: This is string")
  } else if _, ok := inc.(bool); ok {
    fmt.Println("T14: This is bool")
  } else if _, ok := inc.(chan any); ok {
    fmt.Println("T14: This is channel")
  } else {
    fmt.Println("T14: This is unknown type")
  }
}

func t14_test() {
  t14(1)
  t14("Any")
  t14(true)
  ch:=make(chan interface{})
  t14(ch)
}
