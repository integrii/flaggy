package flaggy

import "testing"

func TestDoubleParse(t *testing.T) {
	ResetParser()
	DefaultParser.ShowHelpOnUnexpected = false

	err := DefaultParser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	err = DefaultParser.Parse()
	if err == nil {
		t.Fatal(err)
	}
}

func TestDisableShowVersionFlag(t *testing.T) {
	ResetParser()

	// if this fails the function tested might be useless.
	// Review if it's still useful and adjust.
	if DefaultParser.ShowVersionWithVersionFlag != true {
		t.Fatal("The tested function might not make sense any more.")
	}

	DefaultParser.DisableShowVersionWithVersion()

	if DefaultParser.ShowVersionWithVersionFlag != false {
		t.Fatal("ShowVersionWithVersionFlag should have been false.")
	}
}

func TestFindArgsNotInParsedValues(t *testing.T) {
	t.Parallel()

	// ensure all 'test.' values are skipped
	args := []string{"test.timeout=10s", "test.v=true"}
	parsedValues := []parsedValue{}
	unusedArgs := findArgsNotInParsedValues(args, parsedValues)
	if len(unusedArgs) > 0 {
		t.Fatal("Found 'test.' args as unused when they should be ignored")
	}

	// ensure regular values are not skipped
	parsedValues = []parsedValue{
		parsedValue{
			Key:   "flaggy",
			Value: "testing",
		},
	}
	args = []string{"flaggy", "testing", "unusedFlag"}
	unusedArgs = findArgsNotInParsedValues(args, parsedValues)
	t.Log(unusedArgs)
	if len(unusedArgs) == 0 {
		t.Fatal("Found no args as unused when --flaggy=testing should have been detected")
	}
	if len(unusedArgs) != 1 {
		t.Fatal("Invalid number of unused args found.  Expected 1 but found", len(unusedArgs))
	}
}
