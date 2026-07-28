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

// TestMetavarsAreDisabledByDefault verifies existing help output remains
// unchanged until a parser explicitly enables metavars.
func TestMetavarsAreDisabledByDefault(t *testing.T) {
	p := NewParser("metavar-default")
	p.ShowHelpWithHFlag = false
	p.ShowVersionWithVersionFlag = false

	var value string
	p.String(&value, "c", "config-file", "configuration path")

	var help Help
	help.ExtractValues(p, "")

	if len(help.Flags) != 1 {
		t.Fatalf("expected one help flag, got %d", len(help.Flags))
	}
	if help.Flags[0].Metavar != "" {
		t.Fatalf("expected metavars to be disabled, got %q", help.Flags[0].Metavar)
	}
}

// TestMetavarsUseFlagNames verifies long names become uppercase underscore
// metavars and short names provide the fallback when no long name exists.
func TestMetavarsUseFlagNames(t *testing.T) {
	p := NewParser("metavar-names")
	p.ShowHelpWithHFlag = false
	p.ShowVersionWithVersionFlag = false
	p.EnableMetavars()

	var config string
	var shortOnly string
	var toggle bool
	p.String(&config, "c", "config-file", "configuration path")
	p.String(&shortOnly, "p", "", "short-only value")
	p.Bool(&toggle, "v", "verbose", "verbose output")

	var help Help
	help.ExtractValues(p, "")

	if got := help.Flags[0].Metavar; got != "CONFIG_FILE" {
		t.Fatalf("expected CONFIG_FILE metavar, got %q", got)
	}
	if got := help.Flags[1].Metavar; got != "P" {
		t.Fatalf("expected P metavar for short-only flag, got %q", got)
	}
	if got := help.Flags[2].Metavar; got != "" {
		t.Fatalf("expected toggle flag to have no metavar, got %q", got)
	}
	if got := help.Flags[0].ShortDisplay; got != "-c CONFIG_FILE  " {
		t.Fatalf("expected short display with metavar, got %q", got)
	}
	if got := help.Flags[0].LongDisplay; got != "--config-file CONFIG_FILE   " {
		t.Fatalf("expected long display with metavar, got %q", got)
	}
}
