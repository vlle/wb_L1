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

// Определяем методы двух этих интерфейсов
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

// Определяем метод-адаптер
func (adpt *AndroidAdapter) SheepToDream() {
	sheep_sounds := adpt.Live()
	robo_sheep := []int{}
	for _, sound := range sheep_sounds {
		if sound == "Meeee" {
			robo_sheep = append(robo_sheep, 1)
		} else {
			robo_sheep = append(robo_sheep, 0)
		}
	}
	adpt.Dream(robo_sheep)
}

func t21() {
	example := new(AndroidAdapter)
	// adapter function call
	example.SheepToDream()
}
