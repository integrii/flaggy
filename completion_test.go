package flaggy_test

import (
	"strings"
	"testing"

	"github.com/integrii/flaggy"
)

// newCompletionParser builds a parser populated with the fictitious fleet commands used
// to validate shell completion generators across every supported shell variant.
func newCompletionParser() (*flaggy.Parser, []string) {
	p := flaggy.NewParser("starfleet")
	var route string
	p.String(&route, "w", "warp", "Enable warp calibration during deployment")
	p.AddPositionalValue(&route, "sector", 2, false, "Target sector coordinate")

	commands := []string{"deploy", "destroy", "diagnose", "dock"}
	for _, name := range commands {
		sub := flaggy.NewSubcommand(name)
		sub.Description = "Handle " + name + " operations"
		p.AttachSubcommand(sub, 1)
	}

	return p, commands
}

// verifyCompletionCoverage ensures that each fictitious command surfaces within the
// produced completion output so developers can trust the generator to describe real CLIs.
func verifyCompletionCoverage(t *testing.T, output string, commands []string, shell string) {
	t.Helper()
	for _, name := range commands {
		if !strings.Contains(output, name) {
			t.Fatalf("expected %s completion to list command %s: %s", shell, name, output)
		}
	}
}

// TestGenerateBashCompletion exercises bash completions with the shared fleet commands.
func TestGenerateBashCompletion(t *testing.T) {
	p, commands := newCompletionParser()
	out := flaggy.GenerateBashCompletion(p)
	verifyCompletionCoverage(t, out, commands, "bash")
	if !strings.Contains(out, "--warp") {
		t.Fatalf("expected bash completion to include long flag name: %s", out)
	}
	if !strings.Contains(out, "-w") {
		t.Fatalf("expected bash completion to include short flag name: %s", out)
	}
	if !strings.Contains(out, "sector") {
		t.Fatalf("expected bash completion to include positional name: %s", out)
	}
}

// TestGenerateZshCompletion exercises zsh completions with the shared fleet commands.
func TestGenerateZshCompletion(t *testing.T) {
	p, commands := newCompletionParser()
	out := flaggy.GenerateZshCompletion(p)
	verifyCompletionCoverage(t, out, commands, "zsh")
	if !strings.Contains(out, "--warp") {
		t.Fatalf("expected zsh completion to include long flag name: %s", out)
	}
	if !strings.Contains(out, "-w") {
		t.Fatalf("expected zsh completion to include short flag name: %s", out)
	}
	if !strings.Contains(out, "sector") {
		t.Fatalf("expected zsh completion to include positional name: %s", out)
	}
}

// TestGenerateFishCompletion exercises fish completions with the shared fleet commands.
func TestGenerateFishCompletion(t *testing.T) {
	p, commands := newCompletionParser()
	out := flaggy.GenerateFishCompletion(p)
	verifyCompletionCoverage(t, out, commands, "fish")
	if !strings.Contains(out, "complete -c starfleet") {
		t.Fatalf("expected fish completion to target starfleet: %s", out)
	}
	if !strings.Contains(out, "-l warp") {
		t.Fatalf("expected fish completion to include long flag name: %s", out)
	}
	if !strings.Contains(out, "-s w") {
		t.Fatalf("expected fish completion to include short flag name: %s", out)
	}
	if !strings.Contains(out, "sector") {
		t.Fatalf("expected fish completion to include positional name: %s", out)
	}
}

// TestGeneratePowerShellCompletion exercises PowerShell completions with fleet commands.
func TestGeneratePowerShellCompletion(t *testing.T) {
	p, commands := newCompletionParser()
	out := flaggy.GeneratePowerShellCompletion(p)
	verifyCompletionCoverage(t, out, commands, "powershell")
	if !strings.Contains(out, "Register-ArgumentCompleter") {
		t.Fatalf("expected powershell completion to register completer: %s", out)
	}
	if !strings.Contains(out, "--warp") {
		t.Fatalf("expected powershell completion to include long flag name: %s", out)
	}
	if !strings.Contains(out, "-w") {
		t.Fatalf("expected powershell completion to include short flag name: %s", out)
	}
	if !strings.Contains(out, "sector") {
		t.Fatalf("expected powershell completion to include positional name: %s", out)
	}
}

// TestGenerateNushellCompletion exercises Nushell completions with the fleet commands.
func TestGenerateNushellCompletion(t *testing.T) {
	p, commands := newCompletionParser()
	out := flaggy.GenerateNushellCompletion(p)
	verifyCompletionCoverage(t, out, commands, "nushell")
	if !strings.Contains(out, "extern \"starfleet\"") {
		t.Fatalf("expected nushell completion to expose extern signature: %s", out)
	}
	if !strings.Contains(out, "\"--warp\"") {
		t.Fatalf("expected nushell completion to include long flag name: %s", out)
	}
	if !strings.Contains(out, "\"-w\"") {
		t.Fatalf("expected nushell completion to include short flag name: %s", out)
	}
	if !strings.Contains(out, "\"sector\"") {
		t.Fatalf("expected nushell completion to include positional name: %s", out)
	}
}
