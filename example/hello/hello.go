package main

import (
    "fmt"
    "log"

    "example.com/greetings"

    "example.com/hello/play"
)


func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    names := []string{"Ruby", "Weise", "Blake"}
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
    play.Drift()
}
