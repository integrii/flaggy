package flaggy_test

import (
	"os"
	"testing"

	"github.com/integrii/flaggy"
)

// TestSubcommandParse tests paring of a single subcommand
func TestSubcommandParse(t *testing.T) {
	flaggy.DebugMode = true
	defer debugOff()

	// create the argument parser
	ap := flaggy.NewArgumentParser()
	newSC := flaggy.NewSubcommand(1)
	newSC.ShortName = "t"
	newSC.LongName = "test"
	newSC.Description = "Test subcommand"
	ap.AddSubcommand(newSC)

	os.Args = []string{"test"}
	err := ap.Parse()
	if err != nil {
		t.Fatal("Error parsing args: " + err.Error())
	}

	if !newSC.SubcommandUsed {
		t.Fatal("Subcommand was not used, but it was expected to be")
	}
	t.Log("Subcommand was used.")

	//  test what happens if you add a second arg (it should crash with unknown field)
	os.Args = []string{"test", "test2"}
	err = ap.Parse()
	if err != nil {
		t.Fatal("Error parsing args: " + err.Error())
	}

}

// debugOff makes defers easier
func debugOff() {
	flaggy.DebugMode = false
}
