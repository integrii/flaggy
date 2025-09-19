package flaggy_test

import (
	"os"
	"testing"
	"time"

	"github.com/integrii/flaggy"
)

// TestPrintHelpToStdout builds a representative parser and prints the current
// help output to stdout for manual inspection. Run with:
//
//	go test -run TestPrintHelpToStdout -v
func TestPrintHelpToStdout(t *testing.T) {
	p := flaggy.NewParser("help-dump")
	p.Description = "Sample command showing current help formatting."

	// Set up a couple of subcommands
	scA := flaggy.NewSubcommand("subA")
	scA.ShortName = "a"
	scA.Description = "Subcommand A description."

	scB := flaggy.NewSubcommand("subB")
	scB.ShortName = "b"
	scB.Description = "Subcommand B description."

	p.AttachSubcommand(scA, 1)
	scA.AttachSubcommand(scB, 1)

	// Add a positional to demonstrate the section
	var posA = "defaultPosA"
	scA.AddPositionalValue(&posA, "posA", 2, false, "Example positional for A.")

	// Add a few flags of different types
	var s string = "defaultStringHere"
	var i int
	var b bool
	var d time.Duration
	p.String(&s, "s", "stringFlag", "Example string flag.")
	p.Int(&i, "i", "intFlag", "Example int flag.")
	p.Bool(&b, "b", "boolFlag", "Example bool flag.")
	p.Duration(&d, "d", "durationFlag", "Example duration flag.")

	// Optional extra help lines to show placement in template
	p.AdditionalHelpPrepend = "This is a prepend for help"
	p.AdditionalHelpAppend = "This is an append for help"

	// Parse to set subcommand context to scB
	if err := p.ParseArgs([]string{"subA", "subB"}); err != nil {
		t.Fatalf("parse: unexpected error: %v", err)
	}

	// Redirect help output from stderr to stdout for visibility under `go test -v`.
	savedStderr := os.Stderr
	os.Stderr = os.Stdout
	defer func() { os.Stderr = savedStderr }()

	// Print current help to stdout
	p.ShowHelpWithMessage("This is a help message on exit")
}
