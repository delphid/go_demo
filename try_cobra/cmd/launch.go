package cmd

import (
    "fmt"

    "github.com/spf13/cobra"

    p "try_cobra/prepare"
)

func init() {
    rootCmd.AddCommand(launchProc)
}

var launchProc = &cobra.Command{
    Use:    "launch_alala",
    Short:  "the launch proc",
    Long:   "use it to launch a rocket",
    Run:    func(cmd *cobra.Command, args []string) {
        p.GetRocketReady()
        fmt.Println("now launching")
    },
}
