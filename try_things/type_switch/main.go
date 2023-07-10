// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type A string

func (a *A) p() string {
	return string(*a)
}

type B string

func (b *B) p() string {
	return string(*b)
}

type P interface {
	p() string
}

func dop(p P) {
	switch v := p.(type) {
	case *A:
		fmt.Printf("%T\n", v)
		fmt.Println(v.p())
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	s1 := A("aaa")
	dop(&s1)
	s2 := B("bbb")
	dop(&s2)
}
