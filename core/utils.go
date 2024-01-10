package core

import (
	"fmt"
	"strings"
)

// AskYesNo asks a yes/no question to the user and returns the answer as
// a boolean value, by reading from stdin.
func AskYesNo(prompt string) bool {
	var response string
	fmt.Print(prompt)
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}
