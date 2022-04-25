package main

import (
	"bufio"
	"fmt"
	"gosh/lib/history"
	"gosh/lib/parser"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to gosh!")
	reader := bufio.NewReader(os.Stdin)
	commandHistory := history.FromFile()

	for {
		fmt.Printf("> ")
		text, err := reader.ReadString('\n')

		if err != nil {
			panic("Oh no")
		}

		command := parser.GoshExecCommand(text, commandHistory)

		fmt.Println(command)
	}
}

// A utility function used for checking if a command is going to slow down a user's "flow".
// Could be enabled by some sort of flag.
func callTimesOut(in <-chan interface{}, out chan<- interface{}) {
	time.Sleep(2 * time.Second)
	out <- 1
}
