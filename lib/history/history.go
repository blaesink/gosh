package history

import (
	"fmt"
	"gopkg.in/yaml.v3"
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

	goshHistoryLocation = homeDir + "/.config/gosh/goshHistory.yaml"
}

// A struct that holds information about any command run.
//
// Attributes:
// 	Command string => The command issued by the user.
// 	Result int => The result of the command (0 if success, other if failure).
// 	Invocations uint => How many times the user has used this command.
type GoshCommand struct {
	Command     string `yaml:"command"`
	Invocations uint   `yaml:"invocations"`
	result      int
}

func (gc *GoshCommand) command() string {
	return gc.Command
}

func (gc *GoshCommand) res() int {
	return gc.result
}

func NewCommand(text string, result int) *GoshCommand {
	return &GoshCommand{text, 1, result}
}

type GoshHistory struct {
	Commands map[uint32]*GoshCommand `yaml:"commands"`
}

func NewHistory() *GoshHistory {
	commands := make(map[uint32]*GoshCommand, 0)
	return &GoshHistory{commands}
}

func (g *GoshHistory) AddToHistory(c *GoshCommand) (uint32, error) {
	commandHash := hash(c.command())

	if cmd := g.retrieveCommand(commandHash); cmd != nil {
		cmd.Invocations++
	} else {
		g.Commands[commandHash] = c
	}

	return commandHash, nil
}

func (g *GoshHistory) retrieveCommand(hash uint32) *GoshCommand {
	cmd, ok := g.Commands[hash]

	if !ok {
		return nil
	}

	return cmd
}

// Cleans all commands with a non-zero result.
// This keeps the user from entering bad commands.
func (g *GoshHistory) Clean() {
	for h, cmd := range g.Commands {
		if cmd.res() != 0 {
			delete(g.Commands, h)
		}
	}
}

func (g *GoshHistory) size() uint {
	return uint(len(g.Commands))
}

// Writes the GoshHistory struct to JSON ([]byte) for writing.
func (g *GoshHistory) toYAML() ([]byte, error) {
	return yaml.Marshal(g)
}

// Writes the json to file.
func (g *GoshHistory) SaveToFile() {
	content, err := g.toYAML()

	// Probs shouldn't be a panic.
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(goshHistoryLocation, content, 0777)
}

// Loads config file from location.
// TODO: Need some sort of init script for the first package run to make the
// TODO: config file location if it doesn't exist. Which it wouldn't.
func FromConfigFile() (*GoshHistory, error) {
	// fmt.Printf("Looking for history file in %s\n", goshHistoryLocation)
	h := NewHistory()
	content, err := ioutil.ReadFile(goshHistoryLocation)

	if err != nil {
		return h, fmt.Errorf("No config file found!")
	}

	err = yaml.Unmarshal(content, &h)
	return h, err
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
