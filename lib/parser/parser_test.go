package parser

import (
	"gosh/lib/history"
	"testing"
)

func TestGoshExecCommand(t *testing.T) {
	// Just the most inoffensive thing that can be done.
	cmd := "ls"
	expected := history.NewCommand("ls", 0)

	res, err := GoshExecCommand(cmd)

	if err != nil {
		t.Fatalf("Unexpected error while running GoshExecCommand")
	}

	// Dereference since the addresses aren't the same.
	if *res != *expected {
		t.Errorf("Have %v, want %v", res, expected)
	}
}
