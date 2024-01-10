package cmd

import (
	"fmt"

	"github.com/mirkobrombin/lsh/core"
	"github.com/spf13/cobra"
)

func NewExportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "export <exportPath>",
		Short: "Export bookmarks",
		Args:  cobra.ExactArgs(1),
		Run:   exportCmdRun,
	}
}

func exportCmdRun(cmd *cobra.Command, args []string) {
	exportPath := args[0]
	finalExportPath, err := core.ExportBookmarks(exportPath)
	if err != nil {
		fmt.Printf("Error exporting bookmarks: %v\n", err)
		return
	}
	fmt.Printf("Bookmarks exported to %s\n", finalExportPath)
}

func init() {
	rootCmd.AddCommand(NewExportCmd())
}
