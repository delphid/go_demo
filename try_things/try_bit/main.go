package main

import (
	"fmt"
)


func GetSetBits(i uint) int {
	n := int(i)
	if n == 1 {
		return 1
	}
	temp := 0
	setBits := 1
	for n != 1 {
		temp = n
		n = n >> 1
		if temp - n << 1 == 1{
			setBits ++
		}
	}

	return setBits
}

func main() {
	fmt.Println(GetSetBits(6))
}