package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "zero",
    Short: "zero is a cli tool for performing basic mathematical operations",
    Long:  "zero is a cli tool for performing basic mathematical operations - addition, multiplication, division and subtraction.",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Run() int {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
        return 1
    }
    return 0
}
