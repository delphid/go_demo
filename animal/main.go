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

func (x Animals) Eat() {
    fmt.Println(x.food)
}

func (x Animals) Move() {
    fmt.Println(x.locomotion)
}

func main() {
    cow := Animals{SuperAnimals{"walk"}, "grass"}
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
    fmt.Println(cow.food)
}
