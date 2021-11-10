package main

import (
    "fmt"

    "rsc.io/quote"

    "example.com/greetings"
)


func main() {
    message := greetings.Hello("lalala")
    fmt.Println(message)
    fmt.Println(quote.Go())
}
