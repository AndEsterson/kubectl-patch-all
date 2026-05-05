package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "kubectl-patch-all",
    Short: "kubectl-patch-all is a kubectl plugin for batching kubectl patch operations",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func Run() int {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "error: '%s'\n", err)
        return 1
    }
    return 0
}
