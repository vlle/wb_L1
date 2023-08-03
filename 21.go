package main

import (
	"fmt"
	"math/rand"
)

type AndroidJob interface {
  Dream([]int)
}


type SheepJob interface {
  Live() []string
}

type AndroidAdapter struct {
  AndroidJob
  SheepJob
}

func (a *AndroidAdapter) Live() []string {
  v := make([]string, 0)
  actions := []string{"Meeee", "Eating.."}
  for i := 0; i < 10; i++ {
    v = append(v, actions[rand.Intn(len(actions))])
  }
  return v
}

func (a *AndroidAdapter) Dream(v []int) {
  for i := range v {
    if v[i] == 1 {
      fmt.Println("Dreaming..")
    } else if v[i] == 0 {
      fmt.Println("Sheep.. zzz")
    }
  }
}

func (adpt *AndroidAdapter) SheepToDream() {
  v := adpt.Live()
  sheep_to_byte := []int{}
  for _, v := range v {
    if v == "Meeee" {
      sheep_to_byte = append(sheep_to_byte, 1)
    } else {
      sheep_to_byte = append(sheep_to_byte, 0)
    }
  }
  adpt.Dream(sheep_to_byte)
}


func t21() {
  example := new(AndroidAdapter)
  // adapter function call
  example.SheepToDream()
}
