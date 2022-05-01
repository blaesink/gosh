package history

import (
	"bytes"
	"testing"
)

func mockHistory() *GoshHistory {
	h := NewHistory()
	return &h
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
		inp      []GoshCommand
		expected uint
	}

	// Test matrix.
	tests := []testHistoryStruct{
		{[]GoshCommand{{"ls", 0}}, 1},
		{[]GoshCommand{{"ls", 1}}, 0},
		{[]GoshCommand{{"ls", 0}, {"rm -rf /", 1}}, 1},
		{[]GoshCommand{{"ls", 1}, {"rm -rf /", 1}}, 0}}

	for _, tt := range tests {
		// Make a new History for each loop and test the input from each test.
		h := GoshHistory{tt.inp}
		h.Clean()

		if !isEqual(h.Size(), tt.expected) {
			t.Errorf("h.Clean() with commands %v expected size %d, have size %d", tt.inp, tt.expected, h.Size())
		}
	}
}

func TestToJSON(t *testing.T) {
	h := NewHistory()
	h.AddToHistory(mockCommand())

	expected := []byte("{\"commands\":[{\"command\":\"rm -rf /\",\"result\":0}]}")
	js, err := h.ToJSON()

	if !isNil(err) {
		t.Fatalf("Unable to convert json for %v", h)
	}

	// Compare exact bytes.
	if bytes.Compare(js, expected) != 0 {
		t.Fatalf("Have %v, want %v", js, expected)
	}
}

func isEqual(a, b uint) bool {
	return a == b
}

func isNil(e error) bool {
	return e == nil
}
