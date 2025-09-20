package flaggy_test

import (
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/integrii/flaggy"
)

func TestMinimalHelpOutput(t *testing.T) {
	p := flaggy.NewParser("TestMinimalHelpOutput")

	rd, wr, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: error: %s", err)
	}
	savedStderr := os.Stderr
	os.Stderr = wr

	defer func() {
		os.Stderr = savedStderr
	}()

	p.ShowHelp()

	buf := make([]byte, 1024)
	n, err := rd.Read(buf)
	if err != nil {
		t.Fatalf("read: error: %s", err)
	}
	got := strings.Split(string(buf[:n]), "\n")
	// Updated to match current help template output (single leading/trailing blank line)
	want := []string{
		"",
		"TestMinimalHelpOutput",
		"",
		"  Subcommands:",
		"    completion   Generate shell completion script for bash or zsh.",
		"",
		"  Flags:",
		"        --version   Displays the program version string.",
		"    -h  --help      Displays help with available flag, subcommand, and positional value parameters.",
		"",
	}

	if len(got) != len(want) {
		t.Fatalf("help length mismatch: got %d lines, want %d lines\nGot:\n%q\nWant:\n%q", len(got), len(want), got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("help line %d mismatch:\nGot:  %q\nWant: %q", i, got[i], want[i])
		}
	}
}

func TestShowHelpBeforeParseIncludesSubcommands(t *testing.T) {
	p := flaggy.NewParser("root-help")
	alpha := flaggy.NewSubcommand("alpha")
	alpha.ShortName = "a"
	beta := flaggy.NewSubcommand("beta")
	beta.ShortName = "b"
	p.AttachSubcommand(alpha, 1)
	p.AttachSubcommand(beta, 2)

	rd, wr, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: error: %s", err)
	}
	savedStderr := os.Stderr
	os.Stderr = wr

	p.ShowHelp()

	if err := wr.Close(); err != nil {
		t.Fatalf("close: error: %s", err)
	}
	os.Stderr = savedStderr

	data, err := io.ReadAll(rd)
	if err != nil {
		t.Fatalf("read all: error: %s", err)
	}
	output := string(data)

	if !strings.Contains(output, "alpha") {
		t.Fatalf("expected alpha subcommand in help, got:\n%s", output)
	}
	if !strings.Contains(output, "beta") {
		t.Fatalf("expected beta subcommand in help, got:\n%s", output)
	}
}

func TestRootHelpDoesNotShowGlobalFlagsSection(t *testing.T) {
	p := flaggy.NewParser("rootCmd")
	p.Description = "Root description"

	var rootString string
	var rootBool bool
	p.String(&rootString, "s", "string", "Root string flag")
	p.Bool(&rootBool, "b", "bool", "Root bool flag")

	sub := flaggy.NewSubcommand("child")
	var subFlag string
	sub.String(&subFlag, "", "sub-flag", "Subcommand flag")
	p.AttachSubcommand(sub, 1)

	rd, wr, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: error: %s", err)
	}
	savedStderr := os.Stderr
	os.Stderr = wr

	p.ShowHelp()

	if err := wr.Close(); err != nil {
		t.Fatalf("close: error: %s", err)
	}
	os.Stderr = savedStderr

	data, err := io.ReadAll(rd)
	if err != nil {
		t.Fatalf("read: error: %s", err)
	}
	output := string(data)

	if strings.Contains(output, "Global Flags:") {
		t.Fatalf("root help should not contain 'Global Flags:' section:\n%s", output)
	}
	if !strings.Contains(output, "  Flags:") {
		t.Fatalf("expected root Flags section in help output:\n%s", output)
	}
	if !strings.Contains(output, "--string") || !strings.Contains(output, "--bool") {
		t.Fatalf("expected root flags in output:\n%s", output)
	}
}

func TestHelpWithMissingSCName(t *testing.T) {
	defer func() {
		r := recover()
		gotMsg := r.(string)
		wantMsg := "Panic instead of exit with code: 2"
		if gotMsg != wantMsg {
			t.Fatalf("error: got: %s; want: %s", gotMsg, wantMsg)
		}
	}()
	flaggy.ResetParser()
	flaggy.PanicInsteadOfExit = true
	sc := flaggy.NewSubcommand("")
	sc.ShortName = "sn"
	flaggy.AttachSubcommand(sc, 1)
	flaggy.ParseArgs([]string{"x"})
}

// TestHelpOutput tests the display of help with -h
func TestHelpOutput(t *testing.T) {
	flaggy.ResetParser()
	// flaggy.DebugMode = true
	// defer debugOff()

	p := flaggy.NewParser("testCommand")
	p.Description = "Description goes here.  Get more information at https://github.com/integrii/flaggy."
	scA := flaggy.NewSubcommand("subcommandA")
	scA.ShortName = "a"
	scA.Description = "Subcommand A is a command that does stuff"
	scB := flaggy.NewSubcommand("subcommandB")
	scB.ShortName = "b"
	scB.Description = "Subcommand B is a command that does other stuff"
	scX := flaggy.NewSubcommand("subcommandX")
	scX.Description = "This should be hidden."
	scX.Hidden = true

	var posA = "defaultPosA"
	var posB string
	p.AttachSubcommand(scA, 1)
	scA.AttachSubcommand(scB, 1)
	scA.AddPositionalValue(&posA, "testPositionalA", 2, false, "Test positional A does some things with a positional value.")
	scB.AddPositionalValue(&posB, "hiddenPositional", 1, false, "Hidden test positional B does some less than serious things with a positional value.")
	scB.PositionalFlags[0].Hidden = true
	var stringFlag = "defaultStringHere"
	var intFlag int
	var boolFlag bool
	var durationFlag time.Duration
	p.String(&stringFlag, "s", "stringFlag", "This is a test string flag that does some stringy string stuff.")
	p.Int(&intFlag, "i", "intFlg", "This is a test int flag that does some interesting int stuff.")
	p.Bool(&boolFlag, "b", "boolFlag", "This is a test bool flag that does some booly bool stuff.")
	p.Duration(&durationFlag, "d", "durationFlag", "This is a test duration flag that does some untimely stuff.")
	var subFlag string
	scB.String(&subFlag, "", "subFlag", "This is a subcommand-specific flag.")
	p.AdditionalHelpPrepend = "This is a prepend for help"
	p.AdditionalHelpAppend = "This is an append for help"

	rd, wr, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: error: %s", err)
	}
	savedStderr := os.Stderr
	os.Stderr = wr

	defer func() {
		os.Stderr = savedStderr
	}()

	if err := p.ParseArgs([]string{"subcommandA", "subcommandB", "hiddenPositional1"}); err != nil {
		t.Fatalf("got: %s; want: no error", err)
	}
	p.ShowHelpWithMessage("This is a help message on exit")

	buf := make([]byte, 1024)
	n, err := rd.Read(buf)
	if err != nil {
		t.Fatalf("read: error: %s", err)
	}
	got := strings.Split(string(buf[:n]), "\n")
	// Updated to match current help template output without completion subcommand on nested help.
	want := []string{
		"",
		"subcommandB - Subcommand B is a command that does other stuff",
		"",
		"  Flags:",
		"      --subFlag   This is a subcommand-specific flag.",
		"",
		"  Global Flags:",
		"        --version        Displays the program version string.",
		"    -h  --help           Displays help with available flag, subcommand, and positional value parameters.",
		"    -s  --stringFlag     This is a test string flag that does some stringy string stuff. (default: defaultStringHere)",
		"    -i  --intFlg         This is a test int flag that does some interesting int stuff. (default: 0)",
		"    -b  --boolFlag       This is a test bool flag that does some booly bool stuff.",
		"    -d  --durationFlag   This is a test duration flag that does some untimely stuff. (default: 0s)",
		"",
		"This is a help message on exit",
		"",
	}

	if len(got) != len(want) {
		t.Fatalf("help length mismatch: got %d lines, want %d lines\nGot:\n%q\nWant:\n%q", len(got), len(want), got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("help line %d mismatch:\nGot:  %q\nWant: %q", i, got[i], want[i])
		}
	}
}
