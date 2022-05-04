package history

import (
	"bytes"
	"os"
	"testing"
)

var _ = bytes.Compare
var currPath string

func init() {
	currPath, _ = os.Getwd()
	goshHistoryLocation = currPath + "/dummyHistory.yaml"
}

func mockHistory() *GoshHistory {
	h := NewHistory()
	return h
}

func mockCommand() *GoshCommand {
	cmd := NewCommand("rm -rf /", 0)
	return cmd
}

func TestNewHistory(t *testing.T) {
	h := mockHistory()

	if !isEqual(h.size(), 0) {
		t.Fatalf("Have command length of %d, want 0", len(h.Commands))
	}
}

func TestRetrieveCommand(t *testing.T) {
	h := mockHistory()
	h.Commands[1234] = mockCommand()

	cmd := h.retrieveCommand(1234)

	if cmd == nil {
		t.Fatalf("Nil command retrieved.")
	}
}

func TestRetrieveNilCommand(t *testing.T) {
	h := mockHistory()
	h.Commands[1234] = mockCommand()

	// Get command that doesn't exist.
	cmd := h.retrieveCommand(1237)

	if cmd != nil {
		t.Fatalf("Command %v is not nil!", cmd)
	}
}

func TestAddToHistory(t *testing.T) {

	tests := []historyTest{
		{[]*GoshCommand{NewCommand("ls", 0), NewCommand("ps", 0)}, 2},
		{[]*GoshCommand{NewCommand("ls", 0), NewCommand("ps -ak", 1)}, 2}}

	for _, tt := range tests {
		h := mockHistory()
		for _, cmd := range tt.commands {
			h.AddToHistory(cmd)
		}

		if !isEqual(h.size(), tt.wanted) {
			t.Fatalf("Have size %d, want size %d", h.size(), tt.wanted)
		}
	}
}

func TestAddDuplicateCommand(t *testing.T) {

	tests := []historyTest{
		{[]*GoshCommand{
			NewCommand("test1", 0)},
			1},
		{[]*GoshCommand{
			NewCommand("test2", 0),
			NewCommand("test2", 0)},
			2},
		{[]*GoshCommand{
			NewCommand("test3", 0),
			NewCommand("test3", 0),
			NewCommand("test3", 0)},
			3},
	}

	for _, tt := range tests {
		h := NewHistory()

		for _, cmd := range tt.commands {
			h.AddToHistory(cmd)
		}

		id := hash(tt.commands[0].command())

		if cmd := h.retrieveCommand(id); cmd.Invocations != tt.wanted {
			t.Fatalf("Have %d invocations for %s, want %d", cmd.Invocations, cmd.command(), tt.wanted)
		}
	}

}

func TestCleanHistory(t *testing.T) {
	// A struct for the TestCleanHistory test matrices.
	type testHistoryStruct struct {
		inp      map[uint32]*GoshCommand
		expected uint
	}

	tests := []testHistoryStruct{
		{mockMap([]*GoshCommand{NewCommand("ls", 0)}), 1},
		{mockMap([]*GoshCommand{NewCommand("ls", 0), NewCommand("ks", -1)}), 1}}

	for _, tt := range tests {
		h := mockHistory()
		h.Commands = tt.inp
		h.Clean()

		if s := h.size(); s != tt.expected {
			t.Fatalf("TestCleanHistory: have %d want %d", s, tt.expected)
		}
	}
}

func TestToYAML(t *testing.T) {
	h := NewHistory()
	h.AddToHistory(mockCommand())

	_, err := h.toYAML()

	if !isNil(err) {
		t.Fatalf("Unable to convert yaml for %v", h)
	}
}

func TestFromConfigFile(t *testing.T) {
	h, err := FromConfigFile()

	if !isNil(err) {
		t.Fatalf("Could not read config file from %s", goshHistoryLocation)
	}

	if h.Commands == nil {
		t.Fatalf("Have nil map for GoshHistory from file %s", goshHistoryLocation)
	}

	// Look for specific command.
	// In this case, `ls`
	cmd := h.retrieveCommand(1446109160)
	if cmd.command() != "ls" {
		t.Fatalf("Have command %s, want \"ls\"", cmd.command())
	}
}

func TestGetters(t *testing.T) {
	cmd := mockCommand()

	if cmd.command() != "rm -rf /" {
		t.Fatalf("Have %s, want \"rm -rf /\"", cmd.command())
	}

	if cmd.res() != 0 {
		t.Fatalf("Have %d, want 0", cmd.res())
	}
}

func isEqual(a, b uint) bool {
	return a == b
}

func isNil(e error) bool {
	return e == nil
}

func mockMap(cmds []*GoshCommand) (m map[uint32]*GoshCommand) {
	m = make(map[uint32]*GoshCommand)

	for _, cmd := range cmds {
		hash := hash(cmd.command())
		m[hash] = cmd
	}
	return
}

type historyTest struct {
	commands []*GoshCommand
	wanted   uint
}
