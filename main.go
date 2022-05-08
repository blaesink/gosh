package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"gosh/lib/history"
	"gosh/lib/parser"
	"os"
)

func main() {
	goshHistory, loadErr := history.FromConfigFile()

	if loadErr != nil {
		panic(loadErr)
	}

	fmt.Println("Welcome to gosh!")
	rl, err := readline.New("> ")

	if err != nil {
		panic("Couldn't instantiate readline!")
	}

	loadRecentsToHistory(goshHistory)

	defer rl.Close()
	defer cleanup(goshHistory)

	for {

		line, err := rl.Readline()

		if err != nil {
			break
		}

		if len(line) == 0 {
			continue
		}

		cmd, err := parser.GoshExecCommand(line)
		if err != nil {
			fmt.Printf("%s not found!\n", cmd.Command)
		}
		goshHistory.AddToHistory(cmd)
	}
}

func cleanup(h *history.GoshHistory) {
	h.Clean()
	h.SaveToFile()
	os.Exit(0)
}

func loadRecentsToHistory(h *history.GoshHistory) error {
	for _, cmd := range h.Recents {
		fmt.Printf("Adding %s to history\n", cmd)
		err := readline.AddHistory(cmd)

		if err != nil {
			panic("Unable to add loaded history to readline!")
		}
	}
	return nil
}
