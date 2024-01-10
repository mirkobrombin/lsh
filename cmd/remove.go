package cmd

import (
	"fmt"

	"github.com/mirkobrombin/lsh/core"
	"github.com/spf13/cobra"
)

func NewRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove <bookmarkName>",
		Short: "Remove a saved bookmark",
		Args:  cobra.ExactArgs(1),
		Run:   removeCmdRun,
	}
}

func removeCmdRun(cmd *cobra.Command, args []string) {
	bookmarkName := args[0]

	_, err := core.LoadBookmark(bookmarkName)
	if err != nil {
		fmt.Printf("Error loading bookmark: %v\n", err)
		return
	}

	err = core.RemoveBookmark(bookmarkName)
	if err != nil {
		fmt.Printf("Error removing bookmark: %v\n", err)
		return
	}
	fmt.Printf("Bookmark %s removed\n", bookmarkName)
}

func init() {
	rootCmd.AddCommand(NewRemoveCmd())
}
