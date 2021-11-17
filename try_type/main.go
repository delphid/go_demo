package main

import (
    "fmt"
)

type Birder interface {
    Echo()
}

type Owl struct {
    K   string
}

func (o Owl) Echo() {
    fmt.Println("Owl: ", o.K)
}

type Chicken struct {
    K   string
}

func (c Chicken) Echo() {
    fmt.Println("Chicken: ", c.K)
}

func BirdEcho(b Birder) {
    b.Echo()
}

func main() {
    a := Owl{"ohoh"}
    b := Chicken{"chichi"}
    BirdEcho(a)
    BirdEcho(b)
}
