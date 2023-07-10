package main

import (
	"fmt"
	"os"
)

func main() {
	d := []byte("bbb")
	err := os.WriteFile("./a", d, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
