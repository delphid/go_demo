package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type colorGroup struct {
		ID     int `json:",omitempty"`
		Name   string
		Colors []string
	}

	type total struct {
		A *colorGroup `json:"a,omitempty"`
		B string      `json:"b,omitempty"`
		C string      `json:"c,omitempty"`  // non pointer, can omit
		D *string     `json:"d,omitempty"`  // pointer, can omit
		E string      `json:"e, omitempty"` // what matters is the syntax you write "omitempty": no space allowed
	}

	groupWithNilA := total{
		B: "abc",
	}
	b, err := json.Marshal(groupWithNilA)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stderr.Write(b)

	println()

	groupWithPointerToZeroA := total{
		A: &colorGroup{},
		B: "abc",
	}
	b, err = json.Marshal(groupWithPointerToZeroA)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stderr.Write(b)
}
