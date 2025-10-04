package flaggy_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/integrii/flaggy"
)

// runParserWithArgs executes a new parser using the provided os.Args tail while capturing
// stdout, stderr, and the panic payload triggered by exitOrPanic so tests can assert behavior.
func runParserWithArgs(t *testing.T, args []string) (string, string, any) {
	t.Helper()

	originalArgs := os.Args
	os.Args = append([]string{"starfleet"}, args...)
	defer func() {
		os.Args = originalArgs
	}()

	parser := flaggy.NewParser("starfleet")

	originalStdout := os.Stdout
	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}

	originalStderr := os.Stderr
	stderrReader, stderrWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stderr pipe: %v", err)
	}

	os.Stdout = stdoutWriter
	os.Stderr = stderrWriter
	defer func() {
		os.Stdout = originalStdout
		os.Stderr = originalStderr
	}()

	originalPanicSetting := flaggy.PanicInsteadOfExit
	flaggy.PanicInsteadOfExit = true
	defer func() {
		flaggy.PanicInsteadOfExit = originalPanicSetting
	}()

	var recovered any
	func() {
		defer func() {
			recovered = recover()
		}()
		_ = parser.Parse()
	}()

	stdoutWriter.Close()
	stderrWriter.Close()

	stdoutBytes, err := io.ReadAll(stdoutReader)
	if err != nil {
		t.Fatalf("failed to read stdout pipe: %v", err)
	}
	stderrBytes, err := io.ReadAll(stderrReader)
	if err != nil {
		t.Fatalf("failed to read stderr pipe: %v", err)
	}

	stdoutReader.Close()
	stderrReader.Close()

	return string(stdoutBytes), string(stderrBytes), recovered
}

// TestParseCompletionRequiresShell ensures that invoking the completion command without a shell
// prints guidance that lists every supported shell so users know their options.
func TestParseCompletionRequiresShell(t *testing.T) {
	_, stderr, recovered := runParserWithArgs(t, []string{"completion"})
	if recovered == nil {
		t.Fatalf("expected parser to exit when no shell provided")
	}
	expectedExit := "Panic instead of exit with code: 2"
	if recovered != expectedExit {
		t.Fatalf("expected exit code 2 when shell missing: %v", recovered)
	}
	want := "Supported shells: bash zsh fish powershell nushell"
	if !strings.Contains(stderr, want) {
		t.Fatalf("expected stderr to describe all supported shells: %s", stderr)
	}
}

// TestParseCompletionSupportsAllShells confirms that every advertised completion shell runs
// through the parser and emits shell-specific output when provided on the command line.
func TestParseCompletionSupportsAllShells(t *testing.T) {
	cases := []struct {
		shell    string
		expected string
	}{
		{shell: "bash", expected: "# bash completion"},
		{shell: "zsh", expected: "#compdef"},
		{shell: "fish", expected: "# fish completion"},
		{shell: "powershell", expected: "# PowerShell completion"},
		{shell: "nushell", expected: "# nushell completion"},
	}

	for _, tc := range cases {
		stdout, stderr, recovered := runParserWithArgs(t, []string{"completion", tc.shell})
		if recovered == nil {
			t.Fatalf("expected parser to exit after writing %s completion", tc.shell)
		}
		expectedExit := "Panic instead of exit with code: 0"
		if recovered != expectedExit {
			t.Fatalf("expected exit code 0 for %s completion: %v", tc.shell, recovered)
		}
		if len(stderr) > 0 {
			t.Fatalf("expected stderr to be empty for %s completion: %s", tc.shell, stderr)
		}
		if !strings.Contains(stdout, tc.expected) {
			t.Fatalf("expected stdout for %s completion to include %q: %s", tc.shell, tc.expected, stdout)
		}
	}
}
