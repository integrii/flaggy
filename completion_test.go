package flaggy_test

import (
	"strings"
	"testing"

	"github.com/integrii/flaggy"
)

func TestGenerateBashCompletion(t *testing.T) {
	p := flaggy.NewParser("testapp")
	var str string
	p.String(&str, "", "testflag", "flag for testing")
	p.AddPositionalValue(&str, "pos", 2, false, "positional")
	sub := flaggy.NewSubcommand("sub")
	p.AttachSubcommand(sub, 1)

	out := flaggy.GenerateBashCompletion(p)
	if !strings.Contains(out, "--testflag") {
		t.Fatalf("expected flag in bash completion output: %s", out)
	}
	if !strings.Contains(out, "sub") {
		t.Fatalf("expected subcommand in bash completion output: %s", out)
	}
	if !strings.Contains(out, "pos") {
		t.Fatalf("expected positional in bash completion output: %s", out)
	}
}

func TestGenerateZshCompletion(t *testing.T) {
	p := flaggy.NewParser("testapp")
	var str string
	p.String(&str, "", "testflag", "flag for testing")
	p.AddPositionalValue(&str, "pos", 2, false, "positional")
	sub := flaggy.NewSubcommand("sub")
	p.AttachSubcommand(sub, 1)

	out := flaggy.GenerateZshCompletion(p)
	if !strings.Contains(out, "--testflag") {
		t.Fatalf("expected flag in zsh completion output: %s", out)
	}
	if !strings.Contains(out, "sub") {
		t.Fatalf("expected subcommand in zsh completion output: %s", out)
	}
	if !strings.Contains(out, "pos") {
		t.Fatalf("expected positional in zsh completion output: %s", out)
	}
}
