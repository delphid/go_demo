package main

import (
    "fmt"
)

var a string

func init() {
    fmt.Println("init in server.go")
    a = "aaa"
}

func main() {
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
