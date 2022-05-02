package history

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
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

// A struct that holds information about any command run.
//
// Attributes:
// 	Command string => The command issued by the user.
// 	Result int => The result of the command (0 if success, other if failure).
// 	Invocations uint => How many times the user has used this command.
type GoshCommand struct {
	Command     string `json:"command"`
	Result      int    `json:"result"`
	Invocations uint   `json:"invocations"`
}

func NewCommand(text string, result int) GoshCommand {
	return GoshCommand{text, result, 1}
}

type GoshHistory struct {
	Commands map[uint32]GoshCommand `json:"commands"`
}

func NewHistory() *GoshHistory {
	commands := make(map[uint32]GoshCommand, 0)
	return &GoshHistory{commands}
}

func (g *GoshHistory) AddToHistory(c GoshCommand) (uint32, error) {
	commandHash := hash(c.Command)

	if cmd, ok := g.Commands[commandHash]; !ok {
		g.Commands[commandHash] = c
	} else {
		// Increment the amount of times we've run this command.
		cmd.Invocations = cmd.Invocations + 1
	}
	return commandHash, nil
}

// Cleans all commands with a non-zero hsult.
// This keeps the user from entering bad commands.
// NOTE: this is technically reslicing, and can be very slow as size increases.
func (g *GoshHistory) Clean() {
	for h, cmd := range g.Commands {
		if cmd.Result != 0 {
			delete(g.Commands, h)
		}
	}
}

func (g *GoshHistory) Size() uint {
	return uint(len(g.Commands))
}

// Writes the GoshHistory struct to JSON ([]byte) for writing.
func (g *GoshHistory) ToJSON() ([]byte, error) {
	return json.Marshal(g)
}

// Writes the json to file.
func (g *GoshHistory) SaveToFile() error {
	content, err := g.ToJSON()

	// Probs shouldn't be a panic.
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(goshHistoryLocation, content, 0777)
	return nil
}

// Loads config file from location.
// TODO: Need some sort of init script for the first package run to make the
// TODO: config file location if it doesn't exist. Which it wouldn't.
func FromConfigFile() (*GoshHistory, error) {
	fmt.Printf("Looking for history file in %s\n", goshHistoryLocation)

	h := NewHistory()
	content, err := ioutil.ReadFile(goshHistoryLocation)

	if err != nil {
		return h, fmt.Errorf("No config file found!")
	}

	err = json.Unmarshal(content, &h)
	return h, err
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
