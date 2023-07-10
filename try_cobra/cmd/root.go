package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	flags struct {
		name string
	}
)

var rootCmd = &cobra.Command{
	Use:   "try_cobra",
	Short: "just try cobra",
	Long:  "we are trying to use cobra",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	fmt.Println("start")
	fmt.Println("flag `name`: ", flags.name)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println("root init")
	d := func() string {
		fmt.Println("providing default name")
		return "default name"
	}
	rootCmd.PersistentFlags().StringVar(&flags.name, "name", d(), "this is the name")
}
