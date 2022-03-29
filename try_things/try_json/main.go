package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type S struct {
    A *string `json:"a, omitempty"`
    B string `json:"b"`
}

func main() {
    s := S{B: "bbb"}
    fmt.Println(s)
    b, err := json.Marshal(s)
    if err != nil {
        fmt.Println("error:", err)
    }
    os.Stderr.Write(b)
}
