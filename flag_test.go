package flaggy

import (
	"testing"
)

// debugOff makes defers easier
func debugOff() {
	DebugMode = false
}

func TestParseArgWithValue(t *testing.T) {
	testCases := make(map[string][]string)
	testCases["-f=test"] = []string{"f", "test"}
	testCases["--f=test"] = []string{"f", "test"}
	testCases["--flag=test"] = []string{"flag", "test"}
	testCases["-flag=test"] = []string{"flag", "test"}
	testCases["----flag=--test"] = []string{"--flag", "--test"}
	testCases["-b"] = []string{"b", ""}
	testCases["--bool"] = []string{"bool", ""}

	for arg, correctValues := range testCases {
		key, value := parseArgWithValue(arg)
		if key != correctValues[0] {
			t.Fatalf("Flag %s parsed key as %s but expected key %s", arg, key, correctValues[0])
		}
		if value != correctValues[1] {
			t.Fatalf("Flag %s parsed value as %s but expected value %s", arg, value, correctValues[1])
		}
		t.Logf("Flag %s parsed key as %s and value as %s correctly", arg, key, value)
	}
}

func TestDetermineArgType(t *testing.T) {
	testCases := make(map[string]string)
	testCases["-f"] = argIsFlagWithSpace
	testCases["--f"] = argIsFlagWithSpace
	testCases["-flag"] = argIsFlagWithSpace
	testCases["--flag"] = argIsFlagWithSpace
	testCases["positionalArg"] = argIsPositional
	testCases["subcommand"] = argIsPositional
	testCases["sub--+/\\324command"] = argIsPositional
	testCases["--flag=CONTENT"] = argIsFlagWithValue
	testCases["-flag=CONTENT"] = argIsFlagWithValue
	testCases["-anotherfl-ag=CONTENT"] = argIsFlagWithValue
	testCases["--anotherfl-ag=CONTENT"] = argIsFlagWithValue
	testCases["1--anotherfl-ag=CONTENT"] = argIsPositional

	for arg, correctArgType := range testCases {
		argType := determineArgType(arg)
		if argType != correctArgType {
			t.Fatalf("Flag %s determined to be type %s but expected type %s", arg, argType, correctArgType)
		} else {
			t.Logf("Flag %s correctly determined to be type %s", arg, argType)
		}
	}
}
