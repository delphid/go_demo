package main

import "fmt"

type A struct {
	Ka string
    Kc string
}

func (a A) getKc() {
    fmt.Println(a.Kc)
}

type B struct {
	A
	Kb string
    Kc string
}

// if have this func, then `b.getKc()` will yield `bc`
// func (b B) getKc() {
//     fmt.Println(b.Kc)
// }

func main() {
	a := A{"a", "ac"}
	b := B{A: A{"a", "ac"}, Kb: "b", Kc: "bc"}
	fmt.Println(a, b)
	fmt.Printf("b: %+v\n", b)
	fmt.Println(b.Ka, b.Kc, b.A)  // a bc {a ac}
    b.getKc()  // ac
}
