package main

import "fmt"

type R interface{}

func f(r R) {
	fmt.Println(r == nil)
	fmt.Println(r)
}

func main() {
	f(nil)
	var a *string
	a = nil
	f(a)
	fmt.Println(a == nil)
	var r R
	r = a
	fmt.Println(r == nil)
	t, _ := r.(*string)
	fmt.Println(t == nil)
}
