package main

import (
    "fmt"
    "reflect"
)

type Stater interface {
    Exec() Stater
}

type S1 struct {
}

func (s *S1) Exec() Stater {
    return &S2{}
}

type S2 struct {
}

func (s *S2) Exec() Stater {
    return &S3{}
}

type S3 struct {
}

func (s *S3) Exec() Stater {
    return &S3{}
}

func main() {
    var nowS Stater
    nowS = &S1{}
    for {
        fmt.Println(reflect.TypeOf(nowS))
        if _, ok := nowS.(*S3); ok {
            break
        }
        nextS := nowS.Exec()
        nowS = nextS
    }
    fmt.Println("finish")
}
