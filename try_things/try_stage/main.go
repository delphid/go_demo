package main

import (
    "fmt"
)

type Stage interface {
    Exec()
}

type StageExecutor struct {
    stage Stage
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
    fmt.Println("")

    se := StageExecutor{&fooStage1}
    se.stage.Exec()
    fmt.Println("")
    se = StageExecutor{multiStage}
    se.stage.Exec()
    fmt.Println("")
}
