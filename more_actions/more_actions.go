package more_actions

import "fmt"


func Clap(name string) string {
    message := fmt.Sprintf("Hi, %v. clap!", name)
    return message
}
