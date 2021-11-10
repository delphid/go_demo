package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)


func Hello(name string) (string, error) {
    if name == "" {
        return name, errors.New("empty name")
    }

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)
    for index, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
        fmt.Println(index, name, message)
    }
    return messages, nil
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
        "Alha %v!",
        "Waha %v!",
        "Oho %v!",
    }

    return formats[rand.Intn(len(formats))]
}
