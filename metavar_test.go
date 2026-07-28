package flaggy_test

import (
	"testing"

	"github.com/integrii/flaggy"
)

// TestGlobalMetavarToggle verifies package-level helpers control the default
// parser without requiring callers to access its fields directly.
func TestGlobalMetavarToggle(t *testing.T) {
	flaggy.ResetParser()
	defer flaggy.ResetParser()

	if flaggy.DefaultParser.ShowMetavars {
		t.Fatal("metavars should be disabled after ResetParser")
	}

	flaggy.EnableMetavars()
	if !flaggy.DefaultParser.ShowMetavars {
		t.Fatal("EnableMetavars should enable metavars")
	}

	flaggy.DisableMetavars()
	if flaggy.DefaultParser.ShowMetavars {
		t.Fatal("DisableMetavars should disable metavars")
	}
}
