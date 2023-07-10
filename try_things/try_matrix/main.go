package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func main() {
	data := make([]float64, 6)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	const appCount = 3
	data1 := make([]float64, 9)
	proportions := mat.NewDense(3, 3, data1)
	proportions.SetRow(0, []float64{0.75, 0.2, 0.3})
	proportions.SetRow(1, []float64{0.1, 0.55, 0.35})
	proportions.SetRow(2, []float64{0.15, 0.25, 0.35})
	data2 := make([]float64, 3)
	totals := mat.NewDense(3, 1, data2)
	totals.SetCol(0, []float64{100, 200, 300})
	var result mat.Dense
	result.Mul(proportions, totals)
	fmt.Println(result)
}