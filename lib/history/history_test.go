package history

import (
	"bytes"
	"testing"
)

var _ = bytes.Compare

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

func TestAddToHistory(t *testing.T) {
	h := mockHistory()
	h.AddToHistory(mockCommand())

	if !isEqual(h.Size(), 1) {
		t.Fatalf("Have command length of %d, want 1", len(h.Commands))
	}
}

func TestCleanHistory(t *testing.T) {
	// A struct for the TestCleanHistory test matrices.
	type testHistoryStruct struct {
		inp      map[uint32]GoshCommand
		expected uint
	}

	tests := []testHistoryStruct{
		{mockMap([]GoshCommand{{"ls", 0, 1}}), 1}}

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
