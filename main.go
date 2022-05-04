package main

import (
	"bufio"
	"fmt"
	"gosh/lib/history"
	"gosh/lib/parser"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	goshHistory, loadErr := history.FromConfigFile()

	if loadErr != nil {
		panic(loadErr)
	}
	AwaitCloseSignal(goshHistory)

	fmt.Println("Welcome to gosh!")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ") // Could be some user defined one too.
		text, err := reader.ReadString('\n')

		if err != nil {
			panic("Oh no")
		}

		cmd, err := parser.GoshExecCommand(text)
		if err != nil {
			fmt.Printf("%s not found!\n", cmd.Cmd.Command)
		}
		goshHistory.AddToHistory(cmd)
	}
}

func AwaitCloseSignal(h *history.GoshHistory) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		h.Clean()
		h.SaveToFile()
		os.Exit(0)
	}()
}
