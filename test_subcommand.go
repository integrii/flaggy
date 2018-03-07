package flaggy

import (
	"os"
	"testing"
)

func TestSubcommandParse(t *testing.T) {
	DebugMode = true
	defer debugOff()

	// create the argument parser
	ap := NewArgumentParser()
	newSC := NewSubCommand(1)
	newSC.ShortName = "t"
	newSC.LongName = "test"
	newSC.Description = "Test subcommand"
	ap.AddSubcommand(newSC)

	os.Args = []string{"test"}
	err := ap.Parse()
	if err != nil {
		t.Fatal("Error parsing args: " + err.Error())
	}

}

// debugOff makes defers easier
func debugOff() {
	DebugMode = false
}
