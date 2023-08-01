package main

import (
  "fmt"
)

type Human struct {}

func (h *Human) SayHi() {
  fmt.Println("Hi")
}

type Action struct {
  Human
}
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

func t1() { 
  a := Action{}
  a.SayHi()
}

