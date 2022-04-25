package parser

import (
	"gosh/lib/builtins"
	"gosh/lib/history"
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
func GoshExecCommand(text string, h history.GoshHistory) history.GoshCommand {
	// Remove the newline character.
	commandText := strings.TrimSuffix(text, "\n")
	args := strings.Split(commandText, " ")

	if builtins.CheckForBuiltin(commandText, h) == true {
		return history.NewCommand(text, 0)
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	err := cmd.Run()

	if err != nil {
		return history.NewCommand(text, -1)
	}

	return history.NewCommand(text, 0)
}

// TODO: Runs a command through a channel in order to provide an "async" feel.
// func GoshExecAsync(in <-chan string, out chan<- history.GoshCommand) {
// 	// Get command string from channel.
// 	commandString := <-in
// 	result := GoshExecCommand(commandString)
// 	out <- result
// }
