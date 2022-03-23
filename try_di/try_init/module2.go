package main

import (
    "fmt"
)

var c string

func init() {
    fmt.Println("init in module2.go")
    c = "ccc"
    b = "bbbc"
    a = "aaac"
}
