package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsh",
	Short: "A simple SSH connection manager",
	Long:  `LSH (Lazy SSH Human) is a CLI tool to manage and simplify SSH connections, designed for lazy humans like me.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
