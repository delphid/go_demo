package main

import (
    "fmt"
)


type Animaler interface {
    Eat()
    Move()
}

type SuperAnimals struct {
    locomotion string
}

type Animals struct {
    SuperAnimals
    food    string
}

type NormalAnimals struct {
    food    string
}

func (x Animals) Eat() {
    fmt.Println(x.food)
}

func (x Animals) Move() {
    fmt.Println(x.locomotion)
}

func (x *NormalAnimals) NormalEat(mood string) {
    fmt.Println("eat", x.food, mood)
}

func main() {
    var cow Animaler
    cow = &Animals{SuperAnimals{"walk"}, "grass"}
    cow.Eat()
    cow.Move()
    someAnimals := []Animals{
        {food: "sworm"},
        Animals{SuperAnimals{"dive"}, "sand"},
    }
    someAnimals[0].Eat()
    someAnimals[0].Move()
    someAnimals[1].Eat()
    someAnimals[1].Move()
    //fmt.Println(cow.food)

    bull := &NormalAnimals{
        food:   "green grass",
    }
    bull.NormalEat(`happily`)
}
