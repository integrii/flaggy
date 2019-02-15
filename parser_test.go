package flaggy

import "testing"

func TestDoubleParse(t *testing.T) {
	ResetParser()

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

	if DefaultParser.ShowVersionWithVersionFlag != true {
		t.Fatal("Why not true?")
	}

	DefaultParser.DisableShowVersionWithVersion()

	if DefaultParser.ShowVersionWithVersionFlag != false {
		t.Fatal("ShowVersionWithVersionFlag should have been false.")
	}
}
