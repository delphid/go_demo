package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:    "try_cobra",
    Short:  "just try cobra",
    Long:   "we are trying to use cobra",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Print("the root cmd")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
