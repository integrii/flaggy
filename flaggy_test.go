package flaggy_test

import (
	"os"
	"testing"

	"github.com/integrii/flaggy"
)

func TestParsePositionalsA(t *testing.T) {
	flaggy.DebugMode = true
	inputLine := []string{"thisBinary", "subcommand", "-t", "-n", "dashN", "positionalA", "positionalB"}
	os.Args = inputLine

	parser := flaggy.NewArgumentParser()
	parser.AddBoolFlag(flaggy.BoolFlag{
		ShortName:   "t",
		LongName:    "",
		Description: "Example boolean flag.",
	})

}
