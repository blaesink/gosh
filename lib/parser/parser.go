package parser

import (
	"os"
	"os/exec"
	"strings"
)

// Args:
// 	text string => The user input text to parse.
// Returns:
// 	history.GoshCommand => A formed GoshCommand that contains the original text and its result.
func GoshParseLine(text string) []string {
	// TODO: function composition ("piping") as `a _ _ . b` => b(a(_,_))
	commands := strings.Split(text, " ")

	return commands
}

// Args:
// 	commands [][]string => the command(s) to exe    arr := make([]string, 0)cute.
// Returns:
// 	history.GoshCommand => A struct containing the result code and original text.
func GoshExecCommand(text string) ([]string, error) {
	// Remove the newline character.
	commandText := strings.TrimSuffix(text, "\n")
	args := strings.Split(commandText, " ")

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return args, cmd.Run()
}
