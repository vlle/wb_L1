package main

type Android interface {
  Dream([]int)
}

func (a *Android) Dream([]int) {

}

type Sheep interface {
  Live() []string
}

type AndroidAdapter struct {
  a Android
  s Sheep
}

func (adpt *AndroidAdapter) SheepToDream() {
  v := adpt.s.Live()
  sheep_to_byte := []int{}
  for _, v := range v {
    if v == "Meeee" {
      sheep_to_byte = append(sheep_to_byte, 1)
    } else {
      sheep_to_byte = append(sheep_to_byte, 0)
    }
  }
  adpt.a.Dream(sheep_to_byte)
}


func t21() {

}
