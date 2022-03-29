package main

import (
    "fmt"
)

type Stage interface {
    Exec()
}

func ExecStage(s Stage) {
    s.Exec()
}

type FooStage struct {
    someValues []string
}

func (fs *FooStage) Exec() {
    for _, value := range fs.someValues {
        fmt.Println(value)
    }
}

type MultiStage []Stage

func (ms MultiStage) Exec() {
    for _, s := range ms {
        s.Exec()
    }
}

func main() {
    fooStage1 := FooStage{[]string{"aaa", "bbb"}}
    fooStage2 := FooStage{[]string{"ccc", "ddd"}}
    multiStage := MultiStage{&fooStage1, &fooStage2}

    ExecStage(&fooStage1)
    fmt.Println("")
    ExecStage(multiStage)
}
