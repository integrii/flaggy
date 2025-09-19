package flaggy

import (
	"strings"
	"testing"
)

func TestNestedSubcommandsParseLocalFlags(t *testing.T) {
	// Command: ./app subcommandA subcommandB -flagA 5
	// Expect: subcommandA and subcommandB used; flagA parsed as int 5 on subcommandB.
	t.Parallel()

	// Setup root parser and nested subcommands to mirror the CLI hierarchy.
	p := NewParser("app")
	subA := NewSubcommand("subcommandA")
	subB := NewSubcommand("subcommandB")
	p.AttachSubcommand(subA, 1)
	subA.AttachSubcommand(subB, 1)

	// Define flagA on subcommandB with storage for the parsed integer.
	var flagA int
	subB.Int(&flagA, "", "flagA", "int flag scoped to subcommandB")

	// Parse the CLI input exactly as described in the command comment.
	if err := p.ParseArgs([]string{"subcommandA", "subcommandB", "-flagA", "5"}); err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	// Assert both subcommands are marked as used and the flag value is correct.
	if !subA.Used || !subB.Used {
		t.Fatalf("expected subcommands to be marked used: subA=%v subB=%v", subA.Used, subB.Used)
	}
	if flagA != 5 {
		t.Fatalf("expected flagA to be 5, got %d", flagA)
	}
}

func TestRootAndChildFlagsAreIsolated(t *testing.T) {
	// Command: ./app -flagA=5 --flagB 5 subcommandA --flagC hello
	// Expect: root flagA=5, root flagB=5, subcommandA used with flagC="hello".
	t.Parallel()

	// Setup root parser with a single subcommand to host child-scoped flags.
	p := NewParser("app")
	subA := NewSubcommand("subcommandA")
	p.AttachSubcommand(subA, 1)

	// Define storage locations for the root and child flag values.
	var flagA int
	var flagB int
	var flagC string

	// Register two root flags and one child flag matching the command contract.
	p.Int(&flagA, "", "flagA", "root int flag")
	p.Int(&flagB, "", "flagB", "second root int flag")
	subA.String(&flagC, "", "flagC", "child string flag")

	// Parse the CLI input exactly as described in the command comment.
	args := []string{"-flagA=5", "--flagB", "5", "subcommandA", "--flagC", "hello"}
	if err := p.ParseArgs(args); err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	// Verify root flags resolve correctly and the child flag receives "hello".
	if flagA != 5 {
		t.Fatalf("expected flagA to be 5, got %d", flagA)
	}
	if flagB != 5 {
		t.Fatalf("expected flagB to be 5, got %d", flagB)
	}
	if flagC != "hello" {
		t.Fatalf("expected flagC to be hello, got %q", flagC)
	}
	// Confirm the subcommand was invoked during parsing.
	if !subA.Used {
		t.Fatalf("expected subcommandA to be used")
	}
}

func TestFlagNameCollisionWithSubcommand(t *testing.T) {
	// Command: ./app -flagA=test flagA
	// Expect: root flagA string set to "test" while flagA subcommand is used.
	t.Parallel()

	// Setup root parser with a subcommand whose name collides with a root flag.
	p := NewParser("app")
	subFlagA := NewSubcommand("flagA")
	p.AttachSubcommand(subFlagA, 1)

	// Register the root-level string flag sharing the subcommand's name.
	var rootFlag string
	p.String(&rootFlag, "", "flagA", "root string flag")

	// Parse the CLI input exactly as described in the command comment.
	if err := p.ParseArgs([]string{"-flagA=test", "flagA"}); err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	// Validate both the root flag value and the subcommand usage status.
	if rootFlag != "test" {
		t.Fatalf("expected root flag to be \"test\", got %q", rootFlag)
	}
	if !subFlagA.Used {
		t.Fatalf("expected subcommand flagA to be used")
	}
}

func TestBlankAndWhitespaceValues(t *testing.T) {
	// Command: ./app -flagA "" subcommandA -a "" -b " " -c XYZ
	// Expect: root flagA blank, -a blank, -b single space, -c "XYZ", subcommandA used.
	t.Parallel()

	// Setup root parser with subcommandA to hold scoped string flags.
	p := NewParser("app")
	subA := NewSubcommand("subcommandA")
	p.AttachSubcommand(subA, 1)

	// Define storage for root and subcommand flag values.
	var rootFlag string
	var flagA string
	var flagB string
	var flagC string

	// Register the root flag and subcommand flags matching the CLI usage.
	p.String(&rootFlag, "", "flagA", "root string flag")
	subA.String(&flagA, "a", "", "blank string flag")
	subA.String(&flagB, "b", "", "single space flag")
	subA.String(&flagC, "c", "", "non blank flag")

	// Parse the CLI input exactly as described in the command comment.
	args := []string{"-flagA", "", "subcommandA", "-a", "", "-b", " ", "-c", "XYZ"}
	if err := p.ParseArgs(args); err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	// Assert blank and whitespace values were preserved for each flag.
	if rootFlag != "" {
		t.Fatalf("expected root flagA to be blank, got %q", rootFlag)
	}
	if flagA != "" {
		t.Fatalf("expected -a flag to be blank, got %q", flagA)
	}
	if flagB != " " {
		t.Fatalf("expected -b flag to be a single space, got %q", flagB)
	}
	if flagC != "XYZ" {
		t.Fatalf("expected -c flag to be XYZ, got %q", flagC)
	}
	// Confirm the subcommand was invoked during parsing.
	if !subA.Used {
		t.Fatalf("expected subcommandA to be used")
	}
}

func TestRootBoolAfterSubcommand(t *testing.T) {
	// Command: ./app subcommandA --output
	// Expect: subcommandA used; root --output bool flag set to true.
	t.Parallel()

	// Setup root parser with subcommandA to mirror the CLI input.
	p := NewParser("app")
	subA := NewSubcommand("subcommandA")
	p.AttachSubcommand(subA, 1)

	// Register the root-level bool flag that follows the subcommand.
	var output bool
	p.Bool(&output, "", "output", "root bool flag")

	// Parse the CLI input exactly as described in the command comment.
	if err := p.ParseArgs([]string{"subcommandA", "--output"}); err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	// Assert the bool flag is true and the subcommand is marked as used.
	if !output {
		t.Fatalf("expected --output to set output to true")
	}
	if !subA.Used {
		t.Fatalf("expected subcommandA to be used")
	}
}

func TestNestedSubcommandTrailingArguments(t *testing.T) {
	// Command: ./app one two --test -- abc 123 xyz
	// Expect: subcommands one & two used; --test bool true; trailing args joined to "abc 123 xyz".
	t.Parallel()

	// Setup root parser with nested subcommands to reflect the CLI command.
	p := NewParser("app")
	subOne := NewSubcommand("one")
	subTwo := NewSubcommand("two")
	p.AttachSubcommand(subOne, 1)
	subOne.AttachSubcommand(subTwo, 1)

	// Register the bool flag on the deepest subcommand per the command contract.
	var test bool
	subTwo.Bool(&test, "", "test", "bool flag on nested subcommand")

	// Parse the CLI input exactly as described in the command comment.
	args := []string{"one", "two", "--test", "--", "abc", "123", "xyz"}
	if err := p.ParseArgs(args); err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	// Validate subcommands were used, flag set to true, and trailing args preserved.
	if !subOne.Used || !subTwo.Used {
		t.Fatalf("expected nested subcommands to be used: one=%v two=%v", subOne.Used, subTwo.Used)
	}
	if !test {
		t.Fatalf("expected --test to set the nested bool flag to true")
	}
	if got := strings.Join(p.TrailingArguments, " "); got != "abc 123 xyz" {
		t.Fatalf("expected trailing arguments to join to %q, got %q", "abc 123 xyz", got)
	}
}
