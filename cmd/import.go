package cmd

import (
	"fmt"

	"github.com/mirkobrombin/lsh/core"
	"github.com/spf13/cobra"
)

func NewImportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "import <importPath>",
		Short: "Import bookmarks",
		Args:  cobra.ExactArgs(1),
		Run:   importCmdRun,
	}
}

func importCmdRun(cmd *cobra.Command, args []string) {
	importPath := args[0]
	err := core.ImportBookmarks(importPath)
	if err != nil {
		fmt.Printf("Error importing bookmarks: %v\n", err)
		return
	}
	fmt.Printf("Bookmarks imported from %s\n", importPath)
}

func init() {
	rootCmd.AddCommand(NewImportCmd())
}
