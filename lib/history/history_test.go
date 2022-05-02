package history

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

var _ = bytes.Compare
var currPath string

func init() {
	currPath, _ = os.Getwd()
	goshHistoryLocation = currPath + "/dummyHistory.json"

	fmt.Println(goshHistoryLocation)
}

func mockHistory() *GoshHistory {
	h := NewHistory()
	return h
}

func mockCommand() GoshCommand {
	return NewCommand("rm -rf /", 0)
}

func TestNewHistory(t *testing.T) {
	h := mockHistory()

	if !isEqual(h.Size(), 0) {
		t.Fatalf("Have command length of %d, want 0", len(h.Commands))
	}
}

func TestRetrieveCommand(t *testing.T) {
	h := mockHistory()
	id, err := h.AddToHistory(mockCommand())

	if err != nil {
		t.Fatalf("%v", err)
	}

	cmd, ok := h.RetrieveCommand(id)

	if !ok {
		t.Fatalf("Command %s does not exist", cmd.Command)
	}
}

func TestAddToHistory(t *testing.T) {
	type historyTest struct {
		commands   []GoshCommand
		wantedSize uint
	}

	tests := []historyTest{
		{[]GoshCommand{{"ls", 0, 1}, {"ps", 0, 1}}, 2}}

	for _, tt := range tests {
		h := mockHistory()
		for _, cmd := range tt.commands {
			h.AddToHistory(cmd)
		}

		if !isEqual(h.Size(), tt.wantedSize) {
			t.Fatalf("Have size %d, want size %d", h.Size(), tt.wantedSize)
		}
	}
}

func TestAddDuplicateCommand(t *testing.T) {
	h := NewHistory()
	h.AddToHistory(mockCommand())
	id, err := h.AddToHistory(mockCommand())

	expectedInvocations := 2

	if err != nil {
		t.Fatalf("%v", err)
	}

	if inv := h.Commands[id].Invocations; inv != uint(expectedInvocations) {
		t.Fatalf("Have %d invocation(s) for command %s, want %d invocations",
			inv, mockCommand().Command, expectedInvocations)
	}
}

func TestCleanHistory(t *testing.T) {
	// A struct for the TestCleanHistory test matrices.
	type testHistoryStruct struct {
		inp      map[uint32]GoshCommand
		expected uint
	}

	tests := []testHistoryStruct{
		{mockMap([]GoshCommand{{"ls", 0, 1}}), 1},
		{mockMap([]GoshCommand{{"ls", 0, 1}, {"ks", -1, 1}}), 1}}

	for _, tt := range tests {
		h := mockHistory()
		h.Commands = tt.inp
		h.Clean()

		if s := h.Size(); s != tt.expected {
			t.Fatalf("TestCleanHistory: have %d want %d", s, tt.expected)
		}
	}
}

func TestToJSON(t *testing.T) {
	h := NewHistory()
	h.AddToHistory(mockCommand())

	// expected := []byte("{\"commands\":[{\"command\":\"rm -rf /\",\"result\":0}]}")
	_, err := h.ToJSON()

	if !isNil(err) {
		t.Fatalf("Unable to convert json for %v", h)
	}

	// // Compare exact bytes.
	// if bytes.Compare(js, expected) != 0 {
	// 	t.Fatalf("Have %v, want %v", js, expected)
	// }
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
	if cmd, ok := h.Commands[1446109160]; !ok {
		t.Fatalf("No command found!")
	} else {
		if cmd.Command != "ls" {
			t.Fatalf("Have command %s, want \"ls\"", cmd.Command)
		}
	}
}

func isEqual(a, b uint) bool {
	return a == b
}

func isNil(e error) bool {
	return e == nil
}

func mockMap(cmds []GoshCommand) (m map[uint32]GoshCommand) {
	m = make(map[uint32]GoshCommand)

	for _, cmd := range cmds {
		hash := hash(cmd.Command)
		m[hash] = cmd
	}
	return
}
