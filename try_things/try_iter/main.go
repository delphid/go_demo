package main

import (
	"fmt"
)

type Iter interface {
	Exec()
	Residual() float64
}

func CalcIter(iter Iter, goal float64, max int) {
	if max<=0 {
		max = 100
	}
	iterNum := 0
	for {
		iterNum += 1
		iter.Exec()
		if iter.Residual()<goal || iterNum>max {
			break
		}
	}
}

type Item struct {
	Value float64
}

func (i *Item) Exec() {
	i.Value = i.Value / 10
}

func (i *Item) Residual() float64 {
	return i.Value
}

func main() {
	item := Item{150}
	CalcIter(&item, 0.01, 0)
	fmt.Println(item)
}