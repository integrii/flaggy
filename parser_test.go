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
