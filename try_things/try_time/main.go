package main

import (
	"fmt"
	"time"
)

func main() {
	rawDate := "2022-07-10"
	date, err := time.Parse("2006-01-02", rawDate)
	fmt.Println(err)
	fmt.Println(date)
}
	
