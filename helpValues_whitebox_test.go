package flaggy

import (
	"testing"
)

func TestMakeSpacer(t *testing.T) {
	if spacer := makeSpacer("short", 20); len(spacer) != 15 {
		t.Errorf("spacer length expected to be 15, got %d.", len(spacer))
	}

	if spacer := makeSpacer("very long", 20); len(spacer) != 11 {
		t.Errorf("spacer length expected to be 11, got %d.", len(spacer))
	}

	if spacer := makeSpacer("very long", 0); len(spacer) != 0 {
		t.Errorf("spacer length expected to be 0, got %d.", len(spacer))
	}
}

func TestGetLongestNameLength(t *testing.T) {
	input := []string{"short", "longer", "very-long"}
	var subcommands []*Subcommand
	var flags []*Flag
	var positionalValues []*PositionalValue

	for _, name := range input {
		subcommands = append(subcommands, NewSubcommand(name))
		flags = append(flags, &Flag{LongName: name})
		positionalValues = append(positionalValues, &PositionalValue{Name: name})
	}

	if l := getLongestNameLength(subcommands, 0); l != 9 {
		t.Errorf("should have returned 9, got %d.", l)
	}

	if l := getLongestNameLength(subcommands, 15); l != 15 {
		t.Errorf("should have returned 15, got %d.", l)
	}

	if l := getLongestNameLength(flags, 0); l != 9 {
		t.Errorf("should have returned 9, got %d.", l)
	}

	if l := getLongestNameLength(flags, 15); l != 15 {
		t.Errorf("should have returned 15, got %d.", l)
	}

	if l := getLongestNameLength(positionalValues, 0); l != 9 {
		t.Errorf("should have returned 15, got %d.", l)
	}

	if l := getLongestNameLength(positionalValues, 15); l != 15 {
		t.Errorf("should have returned 9, got %d.", l)
	}
}
