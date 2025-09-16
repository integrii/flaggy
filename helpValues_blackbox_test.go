package flaggy_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
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
    want := []string{
        "",
        "",
        "  Subcommands: ",
        "    completion   Generate shell completion script for bash or zsh.",
        "",
        "  Flags: ",
        "       --version   Displays the program version string.",
        "    -h --help      Displays help with available flag, subcommand, and positional value parameters.",
    }

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("help mismatch (-want +got):\n%s", diff)
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
    want := []string{
        "subcommandB - Subcommand B is a command that does other stuff",
        "",
        "  Subcommands: ",
        "    completion   Generate shell completion script for bash or zsh.",
        "",
        "  Flags: ",
        "       --version        Displays the program version string.",
        "    -h --help           Displays help with available flag, subcommand, and positional value parameters.",
        "    -s --stringFlag     This is a test string flag that does some stringy string stuff. (default: defaultStringHere)",
        "    -i --intFlg         This is a test int flag that does some interesting int stuff. (default: 0)",
        "    -b --boolFlag       This is a test bool flag that does some booly bool stuff.",
        "    -d --durationFlag   This is a test duration flag that does some untimely stuff. (default: 0s)",
        "",
        "This is a help message on exit",
    }

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("help mismatch (-want +got):\n%s", diff)
	}
}
