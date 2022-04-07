package main

import (
    "errors"
    "fmt"
)

func F(a int) (b int, err error) {
    if a == 0 {
        return a + 1, nil
    } else {
        return a, errors.New("just an error")
    }
}

func G() error {
    return errors.New("error g")
}

func H() error {
    err := G()
    return fmt.Errorf("error h caused by: %w", err)
}

func main() {
    a, err := F(1)
    if err != nil {
        fmt.Println("err: ", err)
    }
    fmt.Println("a: ", a)

    fmt.Println("err: ", err)
    b, err := F(0)
    if err != nil {
        fmt.Println("err: ", err)
    }
    fmt.Println("b: ", b)
    fmt.Println(H())
}
