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

func TestHelpWithMissingSCName(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Expected panic with subcommand avilability at position, but did not get one")
		}
	}()
	flaggy.ResetParser()
	sc := flaggy.NewSubcommand("")
	sc.ShortName = "sn"
	flaggy.AttachSubcommand(sc, 1)
	flaggy.ParseArgs([]string{"x"})
}

// TestHelpOutput tests the dislay of help with -h
func TestHelpOutput(t *testing.T) {
	flaggy.ResetParser()
	// flaggy.DebugMode = true
	// defer debugOff()

	p := flaggy.NewParser("testCommand")
	p.Description = "Description goes here.  Get more information at https://github.com/integrii/flaggy."
	scA := flaggy.NewSubcommand("subcommandA")
	scA.ShortName = "a"
	scA.Description = "Subcommand A is a command that does stuff"
	scB := flaggy.NewSubcommand("subcommandB")
	scB.ShortName = "b"
	scB.Description = "Subcommand B is a command that does other stuff"
	scX := flaggy.NewSubcommand("subcommandX")
	scX.Description = "This should be hidden."
	scX.Hidden = true

	var posA = "defaultPosA"
	var posB string
	p.AttachSubcommand(scA, 1)
	scA.AttachSubcommand(scB, 1)
	scA.AddPositionalValue(&posA, "testPositionalA", 2, false, "Test positional A does some things with a positional value.")
	scB.AddPositionalValue(&posB, "hiddenPositional", 1, false, "Hidden test positional B does some less than serious things with a positional value.")
	scB.PositionalFlags[0].Hidden = true
	var stringFlag = "defaultStringHere"
	var intFlag int
	var boolFlag bool
	var durationFlag time.Duration
	p.String(&stringFlag, "s", "stringFlag", "This is a test string flag that does some stringy string stuff.")
	p.Int(&intFlag, "i", "intFlg", "This is a test int flag that does some interesting int stuff.")
	p.Bool(&boolFlag, "b", "boolFlag", "This is a test bool flag that does some booly bool stuff.")
	p.Duration(&durationFlag, "d", "durationFlag", "This is a test duration flag that does some untimely stuff.")
	p.AdditionalHelpPrepend = "This is a prepend for help"
	p.AdditionalHelpAppend = "This is an append for help"
	p.ParseArgs([]string{"subcommandA", "subcommandB", "hiddenPositional1"})
	p.ShowHelpWithMessage("This is a help message on exit")
}
