package main

import "fmt"

func main() {
	l := []int{1, 2, 3}
	var k []*int
	for _, i := range l {
		n := &i
		m := i
        k = append(k, &i)
		k = append(k, n)
		k = append(k, &m)
	}
	fmt.Println(l)
	fmt.Println(k)
	for _, i := range k {
		fmt.Println(*i)
	}
}
