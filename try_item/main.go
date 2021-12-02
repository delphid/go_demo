package main

import (
    "fmt"
    "os"
)

type A struct {
    K   string
}

func main() {
    for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
    env:= os.Getenv("env")
    defer fmt.Println("the defer")
    fmt.Printf("env: %s\n", env)
    fmt.Printf("is '': %t\n", env == "")
    sl := []A{}
    fmt.Println(sl)
    for _, a := range sl {
        fmt.Println("a.K: ", a.K)
    }
}
