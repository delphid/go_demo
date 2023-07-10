package main

import (
	"fmt"
)

type A map[string]string

type C []string

type S struct {
	K1 map[string]string
	K2 A
	K3 []string
	K4 C
}

func main() {
	a := A{"a": "aaa"}
	b := map[string]string{"b": "bbb"}
	c := C{"c"}
	d := []string{"d"}
	fmt.Println(S{K1: a, K2: b, K3: c, K4: d})
}
