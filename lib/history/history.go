package history

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var goshHistoryLocation string

func init() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic("No user specified!")
	}

	goshHistoryLocation = homeDir + "/.config/gosh/goshHistory.json"
}

type GoshCommand struct {
	Command string `json:"command"`
	Result  int    `json:"result"`
}

func NewCommand(text string, result int) GoshCommand {
	return GoshCommand{text, result}
}

type GoshHistory struct {
	Commands []GoshCommand `json:"commands"`
}

func NewHistory() GoshHistory {
	commands := make([]GoshCommand, 0)
	return GoshHistory{commands}
}

func (g *GoshHistory) AddToHistory(c GoshCommand) {
	g.Commands = append(g.Commands, c)
}

// Cleans all commands with a non-zero hsult.
// This keeps the user from entering bad commands.
// NOTE: this is technically reslicing, and can be very slow as size increases.
func (g *GoshHistory) Clean() {
	for i, cmd := range g.Commands {
		if cmd.Result != 0 {
			if i < len(g.Commands) {
				g.Commands = append(g.Commands[:i], g.Commands[i+1:]...)
			} else {
				g.Commands = g.Commands[:i-1]
			}
		}
	}
}

func (g *GoshHistory) Size() uint {
	return uint(len(g.Commands))
}

func (g *GoshHistory) ToJSON() ([]byte, error) {
	return json.Marshal(g)
}

// Loads config file from location.
// TODO: Need some sort of init script for the first package run to make the
// TODO: config file location if it doesn't exist. Which it wouldn't.
func FromConfigFile() (*GoshHistory, error) {
	fmt.Printf("Looking for history file in %s\n", goshHistoryLocation)

	h := GoshHistory{}
	content, err := ioutil.ReadFile(goshHistoryLocation)

	if err != nil {
		return &GoshHistory{}, fmt.Errorf("No config file found!")
	}

	err = json.Unmarshal(content, &h)
	return &h, err
}
