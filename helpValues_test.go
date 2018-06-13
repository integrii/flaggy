package flaggy_test

import (
	"testing"
	"time"

	"github.com/integrii/flaggy"
)

func TestMinimalHelpOutput(t *testing.T) {
	p := flaggy.NewParser("TestMinimalHelpOutput")
	p.ShowHelp()
}

// TestHelpOutput tests the dislay of help with -h
func TestHelpOutput(t *testing.T) {
	// flaggy.DebugMode = true
	// defer debugOff()
	p := flaggy.NewParser("testCommand")
	p.Description = "Description goes here.  Get more information at http://flaggy.flag."
	scA := flaggy.NewSubcommand("subcommandA")
	scA.ShortName = "a"
	scA.Description = "Subcommand A is a command that does stuff"
	scB := flaggy.NewSubcommand("subcommandB")
	scB.ShortName = "b"
	scB.Description = "Subcommand B is a command that does other stuff"
	scC := flaggy.NewSubcommand("subcommandC")
	scC.ShortName = "c"
	scC.Description = "Subcommand C is a command that does SERIOUS stuff"
	var posA = "defaultPosA"
	var posB string
	p.AttachSubcommand(scA, 1)
	p.AttachSubcommand(scB, 1)
	p.AttachSubcommand(scC, 1)
	p.AddPositionalValue(&posA, "testPositionalA", 2, true, "Test positional A does some things with a positional value.")
	p.AddPositionalValue(&posB, "testPositionalB", 3, false, "Test positional B does some less than serious things with a positional value.")
	var stringFlag string
	var intFlag int
	var boolFlag bool
	var durationFlag time.Duration
	p.String(&stringFlag, "s", "stringFlag", "This is a test string flag that does some stringy string stuff.")
	p.Int(&intFlag, "i", "intFlg", "This is a test int flag that does some interesting int stuff.")
	p.Bool(&boolFlag, "b", "boolFlag", "This is a test bool flag that does some booly bool stuff.")
	p.Duration(&durationFlag, "d", "durationFlag", "This is a test duration flag that does some untimely stuff.")
	p.AdditionalHelpPrepend = "This is a prepend for help"
	p.AdditionalHelpAppend = "This is an append for help"
	p.ShowHelpWithMessage("This is a help addon message")
}
