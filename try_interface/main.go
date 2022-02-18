// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"

	mapset "github.com/deckarep/golang-set"
)

type Base struct {
	ID int
}

func (b Base) GetID() int {
	return b.ID
}

type S struct {
	Base
}

type MyInterface interface {
	GetID() int
}

func f(l []MyInterface) {
	for _, i := range l {
		fmt.Println(i.GetID())
	}
}

func main() {
	a := mapset.NewSetFromSlice([]interface{}{1, 2, 3})
	b := mapset.NewSetFromSlice([]interface{}{2, 3, 4})
	fmt.Println(a.Difference(b))
	fmt.Println(b.Difference(a))

	c := S{Base: Base{ID: 1}}
	ls := []S{c}
	li := make([]MyInterface, len(ls))
	for i := range ls {
		li[i] = ls[i]
		fmt.Println(reflect.TypeOf(ls[i]))
		fmt.Println(reflect.TypeOf(li[i]))
	}
	fmt.Println(reflect.TypeOf(ls))
	fmt.Println(reflect.TypeOf(li))
	f(li)

}


// Set{1}
// Set{4}
// main.S
// main.S
// []main.S
// []main.MyInterface
// 1
