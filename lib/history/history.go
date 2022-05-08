package history

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"hash/fnv"
	"io/ioutil"
	"os"
	"time"
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
// 	LastRunAt string => RFC3339 formatted timestamp of when the command was last issued.
type GoshCommand struct {
	Command     string `yaml:"command"`
	Invocations uint   `yaml:"invocations"`
	result      int
	LastRunAt   string `yaml:"lastRunAt"`
}

func (gc *GoshCommand) command() string {
	return gc.Command
}

func (gc *GoshCommand) res() int {
	return gc.result
}

func NewCommand(text string, result int) *GoshCommand {
	callTime := time.Now().Format(time.RFC3339)
	return &GoshCommand{text, 1, result, callTime}
}

// The GoshHistory struct holds all the information needed to interface with
// the user's history of interacting with the shell.
//
// Commands map[uint32]*GoshCommand => Map of all commands issued.
// 		Any time a user runs a command, it is reflected and updated here.
// RecentL []string => Text of run commands in an array.
// 		In the goshHistory.yaml file, the most recent run command is at the bottom (LIFO).
type GoshHistory struct {
	Commands map[uint32]*GoshCommand `yaml:"commands"`
	Recents  []string                `yaml:"recents"` // FIXME: test me.
}

func NewHistory() *GoshHistory {
	commands := make(map[uint32]*GoshCommand, 0)
	return &GoshHistory{commands, []string{}}
}

func (g *GoshHistory) AddToHistory(c *GoshCommand) (uint32, error) {
	commandHash := hash(c.command())

	if cmd := g.retrieveCommand(commandHash); cmd != nil {
		cmd.Invocations++           // Increment how many time's we've called.
		cmd.LastRunAt = c.LastRunAt // Update the last time we called.
	} else {
		g.Commands[commandHash] = c
	}

	g.Recents = append(g.Recents, c.Command)

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
// Also cleans out duplicates from the RecentL list.
// TODO: this should be done in parallel via goroutines.
func (g *GoshHistory) Clean() {
	l := []string{}
	for hash, cmd := range g.Commands {
		if cmd.res() != 0 {
			delete(g.Commands, hash)
		}
	}

	for _, cmd := range g.Recents {
		if _, ok := g.Commands[hash(cmd)]; ok {
			l = append(l, cmd)
		}
	}
	g.Recents = l
}

func (g *GoshHistory) size() uint {
	return uint(len(g.Commands))
}

// Writes the GoshHistory struct to JSON ([]byte) for writing.
func (g *GoshHistory) toYAML() ([]byte, error) {
	return yaml.Marshal(g)
}

// Writes the json to file.
func (g *GoshHistory) SaveToFile() error {
	content, err := g.toYAML()

	// Probs shouldn't be a panic.
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(goshHistoryLocation, content, 0777)

	if err != nil {
		return fmt.Errorf("Unable to write history to %s", goshHistoryLocation)
	}

	return nil
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
