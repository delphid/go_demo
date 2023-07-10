package cmd

import (
	"fmt"
)

func init() {
	fmt.Println("t_inject init")
	flags.name = "injected name"
}
