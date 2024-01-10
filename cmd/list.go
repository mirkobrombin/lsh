package cmd

import (
	"fmt"

	"github.com/mirkobrombin/lsh/core"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all saved bookmarks",
		Run:   listCmdRun,
	}
}

func listCmdRun(cmd *cobra.Command, args []string) {
	bookmarkList, err := core.ListBookmarks()
	if err != nil {
		fmt.Printf("Error loading bookmarks: %v\n", err)
		return
	}

	fmt.Println("List of bookmarks:")
	for _, bookmark := range bookmarkList {
		fmt.Printf("- %s\n", bookmark.Name)
	}
}

func init() {
	rootCmd.AddCommand(NewListCmd())
}
