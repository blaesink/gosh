package main

import (
	"bufio"
	"fmt"
	"gosh/lib/history"
	"gosh/lib/parser"
	"os"
)

var _ = history.GoshCommand{}

func main() {
	goshHistory, loadErr := history.FromConfigFile()

	if loadErr != nil {
		panic(loadErr)
	}

	fmt.Println("Welcome to gosh!")
	// This should probably be some History struct
	// commandHistory := history.NewHistory()
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("> ")
		text, err := reader.ReadString('\n')

		if err != nil {
			panic("Oh no")
		}

		cmd, err := parser.GoshExecCommand(text)
		if err != nil {
			fmt.Printf("%s not found!\n", cmd.Command)
		}

		goshHistory.AddToHistory(cmd)
	}
}
