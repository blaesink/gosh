package builtins

import (
	"fmt"
	"gosh/lib/history"
	"os"
	"strings"
)

func CheckForBuiltin(text string, history history.GoshHistory) bool {
	switch text {
	case "help":
		help()
		return true
	case "exit":
		history.WriteToFile()
		exit()
		return true
	case "clear":
		clear()
		return true
	case "history":
		history.ShowHistory()
	}
	return false
}

// Prints out some general help information.
func help() {
	fmt.Printf("Well hello there!\n")
}

// Exits the terminal.
func exit() {
	fmt.Println("Bye!")
	os.Exit(0)
}

// Hide your mistakes, or declutter your eyes.
// TODO: this always puts the prompt on the bottom.
func clear() {
	fmt.Println(strings.Repeat("\n", 100))
}
