package flaggy

import "testing"

func TestDetermineArgType(t *testing.T) {
	testCases := make(map[string]string)
	testCases["-f"] = ArgIsFlagWithSpace
	testCases["--f"] = ArgIsFlagWithSpace
	testCases["-flag"] = ArgIsFlagWithSpace
	testCases["--flag"] = ArgIsFlagWithSpace
	testCases["positionalArg"] = ArgIsPositional
	testCases["subcommand"] = ArgIsPositional
	testCases["sub--+/\\324command"] = ArgIsPositional
	testCases["--flag=CONTENT"] = ArgIsFlagWithValue
	testCases["-flag=CONTENT"] = ArgIsFlagWithValue
	testCases["-anotherfl-ag=CONTENT"] = ArgIsFlagWithValue
	testCases["--anotherfl-ag=CONTENT"] = ArgIsFlagWithValue
	testCases["1--anotherfl-ag=CONTENT"] = ArgIsPositional

	for arg, correctArgType := range testCases {
		argType := determineArgType(arg)
		if argType != correctArgType {
			t.Fatalf("Flag %s determined to be type %s but expected type %s", arg, argType, correctArgType)
		} else {
			t.Logf("Flag %s correctly determined to be type %s", arg, argType)
		}
	}
}
