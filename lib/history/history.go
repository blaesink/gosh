package history

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type GoshCommand struct {
	text   string
	result int
}

func NewCommand(text string, result int) GoshCommand {
	return GoshCommand{text, result}
}

type GoshHistory struct {
	commands []GoshCommand
}

func (g GoshHistory) marshall() string {
	u, err := json.Marshal(g)

	if err != nil {
		return "Failure!"
	}

	return string(u)
}

func (g GoshHistory) WriteToFile() {
	file, _ := json.MarshalIndent(g, "", " ")
	ioutil.WriteFile("/usr/share/gosh/goshHistory.json", file, 0777)
}

func newEmptyHistory() GoshHistory {
	commands := make([]GoshCommand, 0)
	return GoshHistory{commands}
}

func FromFile() GoshHistory {
	user, _ := user.Current()
	configDir := "/home/" + user.Username + "/.config/gosh/"
	configPath := configDir + "goshHistory.json"

	_, statErr := os.Stat(configDir)

	if statErr != nil {
		os.MkdirAll(configDir, 0777)
	}

	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		_, rrr := os.Create(configPath)
		if rrr != nil {
			fmt.Println(rrr)
			panic("Unable to create history file" + configPath)
		}
		return newEmptyHistory()
	}

	var history GoshHistory

	json.Unmarshal(file, &history)

	return history
}

func (g *GoshHistory) AddToHistory(c GoshCommand) {
	g.commands = append(g.commands, c)
}

// Clears all commands with a non-zero exit code.
// Useful since we don't want to rerun a bad command!
func (g *GoshHistory) CleanHistory() {
	for i, cmd := range g.commands {
		if cmd.result != 0 {
			// FIXME: this may be inefficient!
			g.commands = append(g.commands[:i], g.commands[i+1:]...)
		}
	}
}

func (g GoshHistory) ShowHistory() {
	fmt.Println(g.commands)
}
