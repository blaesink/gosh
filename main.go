package main

import (
	"bufio"
	"fmt"
	// "gosh/lib/history"
	"gosh/lib/parser"
	"os"
)

func main() {
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

		if args, err := parser.GoshExecCommand(text); err != nil {
			fmt.Printf("%s not found!\n", args[0])
		}

	}
}
