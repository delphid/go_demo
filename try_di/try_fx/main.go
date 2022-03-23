package main

import (
    "context"
    "errors"
    "fmt"
    "time"

    "go.uber.org/fx"
)

type Message string

func NewMessage() Message {
    return Message("a message")
}

type Greeter struct {
    Message Message
    Grumpy bool
}

func NewGreeter(m Message) Greeter {
    grumpy := false
    if time.Now().Unix()%2 == 0 {
        grumpy = true
    }
    return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
    if g.Grumpy {
        return Message("Go away!")
    }
    return g.Message
}

type Event struct {
    Greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
    if g.Grumpy {
        return Event{}, errors.New("could not create event: event.greeter is grumpy")
    }
    return Event{Greeter: g}, nil
}

func (e Event) Start() {
    msg := e.Greeter.Greet()
    fmt.Println(msg)
}

func KickOff(_ Event) {
}

var Module = fx.Options(
    fx.Provide(NewMessage),
    fx.Provide(NewGreeter),
)

func main() {
    app := fx.New(
        Module,
        fx.Provide(NewEvent),
        fx.Invoke(KickOff),
    )
    startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
    app.Start(startCtx)

}
