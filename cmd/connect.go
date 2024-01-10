package cmd

import (
	"fmt"
	"strings"

	"github.com/mirkobrombin/lsh/core"
	"github.com/spf13/cobra"
)

func NewConnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "connect <bookmarkName | connection string>",
		Short: "Connect to a saved bookmark",
		Args:  cobra.ExactArgs(1),
		Run:   connectCmdRun,
	}
}

func connectCmdRun(cmd *cobra.Command, args []string) {
	bookmarkName := args[0]
	bookmark, err := core.LoadBookmark(bookmarkName)

	if err != nil {
		if strings.Contains(bookmarkName, "@") {
			saveBookmark := core.AskYesNo("Detected a connection string. Do you want to save it as a bookmark? (y/n): ")
			if saveBookmark {
				fmt.Print("Enter a name for the bookmark: ")
				fmt.Scanln(&bookmarkName)

				newBookmark := core.Bookmark{
					Name:       bookmarkName,
					Connection: args[0],
				}
				err := core.SaveBookmark(newBookmark)
				if err != nil {
					fmt.Printf("Error saving bookmark %s: %v\n", bookmarkName, err)
					return
				}
				fmt.Printf("Bookmark %s saved.\n", bookmarkName)
				bookmark, err = core.LoadBookmark(bookmarkName)
				if err != nil {
					fmt.Printf("Error loading bookmark %s: %v\n", bookmarkName, err)
					return
				}
			} else {
				return
			}
		} else {
			fmt.Printf("No bookmark found with name %s\n", bookmarkName)
			return
		}
	}

	fmt.Printf("Connecting to bookmark: %s\n", bookmarkName)
	if err := bookmark.Connect(); err != nil {
		fmt.Printf("Error connecting to bookmark %s: %v\n", bookmarkName, err)
		return
	}
}

func init() {
	rootCmd.AddCommand(NewConnectCmd())
}
