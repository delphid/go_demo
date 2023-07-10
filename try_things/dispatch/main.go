package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Assigner interface {
	Total() float64
	Increase(float64)
	Proportions() map[Assigner]float64
	Components() map[Assigner]float64
}

type AssignTask struct {
	AssignFrom Assigner
	AssignAmount float64
}

func Dispatch(a AssignTask) {
	total := a.AssignAmount
	proportions := a.AssignFrom.Proportions()
	for item, ratio := range proportions{
		if item == a.AssignFrom {
			continue
		}
		dispatchAmount := total * ratio
		fmt.Println("total: ", total, ", ratio: ", ratio, ", dispatchAmount: ", dispatchAmount)
		if dispatchAmount < 1 {
			continue
		}
		task := AssignTask{
			AssignFrom: item,
			AssignAmount: dispatchAmount,
		}
		item.Increase(dispatchAmount)
		item.Components()[a.AssignFrom] += dispatchAmount
		a.AssignFrom.Increase(-dispatchAmount)
		a.AssignFrom.Components()[item] -= dispatchAmount

		fmt.Println("now: ")
		for item, _ := range proportions {
			fmt.Println(item.Total())
		}
		fmt.Println("")
		Dispatch(task)
	}
	fmt.Println("a.Total(): ", a.AssignFrom.Total())
}

func DispatchAll(as []Assigner) {
	for _, a := range as {
		fmt.Println("dispatch")
		Dispatch(AssignTask{
			AssignFrom: a,
			AssignAmount: a.Total(),
		})
	}
}

// randPartitions returns n random numbers that add up to 1
func randPartitions(n int) []float64 {
	rand.Seed(time.Now().UnixMilli())
	parts := make([]float64, n)
	partitions := make([]float64, n)
	var sum float64
	for i, _ := range parts {
		parts[i] = rand.Float64()
		sum += parts[i]
	}
	fmt.Println(parts)
	var accumulatePartition float64
	for i, _ := range partitions {
		if i != n-1 {
			partition := parts[i] / sum
			accumulatePartition += partition
				partitions[i] = partition
		} else {
			partitions[i] = 1 - accumulatePartition
		}
	}
	return partitions
}

type CostItem struct {
	Money float64
	CostDistribution map[Assigner]float64
	CostContribution map[Assigner]float64
}

func (a *CostItem) Total() float64 {
	return a.Money
}

func (a *CostItem) Increase(v float64) {
	a.Money += v
}

func (a *CostItem) Proportions() map[Assigner]float64 {
	return a.CostDistribution
}

func (a *CostItem) Components() map[Assigner]float64 {
	return a.CostContribution
}

func main() {
	const itemCount = 3
	items := make([]CostItem, itemCount)
	rand.Seed(time.Now().UnixMilli())
	for i, _ := range items {
		items[i].Money = float64(rand.Intn(5000))
		items[i].CostDistribution = make(map[Assigner]float64)
		items[i].CostContribution = make(map[Assigner]float64)
	}
	for i, _ := range items {
		ps := randPartitions(itemCount)
		for j, p := range ps {
			items[i].CostDistribution[&items[j]] = p
			items[i].CostContribution[&items[j]] = p * items[i].Total()
		}
	}
	assigners := make([]Assigner, itemCount)
	for i, _ := range items {
		assigners[i] = &items[i]
	}
	var sum float64
	fmt.Println("before dispatch: ")
	for _, item := range items {
		sum += item.Total()
		fmt.Println(item.Total())
	}
	fmt.Println("sum: ", sum)

	DispatchAll(assigners)
	sum = 0
	fmt.Println("after dispatch: ")
	for _, item := range items {
		sum += item.Total()
		fmt.Println(item.Total())
		fmt.Println(item.Components())
		var componentsSum float64
		for _, v := range item.Components() {
			componentsSum += v
		}
		fmt.Println("componentsSum: ", componentsSum)
	}
	fmt.Println("sum: ", sum)
}


