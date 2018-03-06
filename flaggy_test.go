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

	var boolT bool

	parser := flaggy.NewArgumentParser()
	parser.AddBoolFlag(&boolT, "t", "", "-t test flag for bool arg")

}
