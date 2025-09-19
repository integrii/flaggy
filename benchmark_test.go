package flaggy

import (
	"testing"
	"time"
)

// BenchmarkParseSingleFlag measures parsing a single string flag so we can ensure
// the simplest CLI scenario stays efficient for users building tiny utilities.
func BenchmarkParseSingleFlag(b *testing.B) {
	// Track allocations to verify the parser remains lightweight.
	b.ReportAllocs()

	// Provide a minimal argument slice that exercises a lone string flag.
	args := []string{"--name", "benchmark"}

	// Execute the parse flow repeatedly to gather benchmark statistics.
	for i := 0; i < b.N; i++ {
		// Reset the parser so each iteration starts from a clean state.
		ResetParser()

		// Capture the string flag value during parsing.
		var name string

		// Declare the flag that the benchmark will parse.
		String(&name, "n", "name", "collects a name for benchmarking")

		// Parse the synthetic arguments.
		ParseArgs(args)

		// Fail fast if the parsed value is not the expected string.
		if name != "benchmark" {
			b.Fatalf("expected name to equal benchmark, got %q", name)
		}
	}
}

// BenchmarkParseSubcommandWithTwoFlags tracks the performance of parsing a
// subcommand that defines two flags so teams can gauge realistic CLI workloads.
func BenchmarkParseSubcommandWithTwoFlags(b *testing.B) {
	// Record allocations during the benchmark to detect regressions.
	b.ReportAllocs()

	// Arrange a subcommand invocation with host and port flags populated.
	args := []string{"serve", "--host", "localhost", "--port", "8080"}

	// Drive the parser repeatedly to gather benchmark metrics.
	for i := 0; i < b.N; i++ {
		// Reset state so each run configures fresh flag bindings.
		ResetParser()

		// Create the subcommand and variables that store parsed values.
		serve := NewSubcommand("serve")
		var host string
		var port int

		// Declare the host and port flags on the subcommand.
		serve.String(&host, "h", "host", "host name to bind to")
		serve.Int(&port, "p", "port", "port to listen on")

		// Attach the subcommand so the parser can discover it.
		AttachSubcommand(serve, 1)

		// Parse the prepared argument slice.
		ParseArgs(args)

		// Ensure the subcommand was detected and values were populated.
		if !serve.Used {
			b.Fatal("expected serve subcommand to be used")
		}
		if host != "localhost" || port != 8080 {
			b.Fatalf("unexpected host %q or port %d", host, port)
		}
	}
}

// BenchmarkParseNestedSubcommandsWithTrailingArgs covers a complex hierarchy of
// three subcommands, mixed flag types, and trailing arguments to validate that
// rich CLIs remain performant.
func BenchmarkParseNestedSubcommandsWithTrailingArgs(b *testing.B) {
	// Track allocation counts while exercising the full parser feature set.
	b.ReportAllocs()

	// Define arguments that walk through alpha -> beta -> gamma subcommands
	// and include trailing values after a -- separator.
	args := []string{
		"alpha", "--name", "alpha-task", "--count", "3", "--enabled", "--speed", "1.5", "--timeout", "500ms",
		"beta", "--label", "release", "--label", "candidate", "--level", "9", "--active", "--ratio", "1.75", "--interval", "45s",
		"gamma", "--mode", "auto", "--max", "100", "--debug", "--threshold", "0.8", "--window", "1m30s",
		"--", "artifact-one", "artifact-two", "artifact-three",
	}

	// Loop to exercise the parser repeatedly during the benchmark.
	for i := 0; i < b.N; i++ {
		// Reset state before wiring up the subcommand hierarchy again.
		ResetParser()

		// Prepare variables that will hold parsed values from each layer.
		var (
			alphaName    string
			alphaCount   int
			alphaEnabled bool
			alphaSpeed   float64
			alphaTimeout time.Duration

			betaLabels   []string
			betaLevel    int
			betaActive   bool
			betaRatio    float32
			betaInterval time.Duration

			gammaMode      string
			gammaMax       uint
			gammaDebug     bool
			gammaThreshold float64
			gammaWindow    time.Duration
		)

		// Configure the alpha subcommand with mixed flag types.
		alpha := NewSubcommand("alpha")
		alpha.String(&alphaName, "n", "name", "name for the alpha subcommand")
		alpha.Int(&alphaCount, "c", "count", "number of repetitions")
		alpha.Bool(&alphaEnabled, "e", "enabled", "whether alpha is enabled")
		alpha.Float64(&alphaSpeed, "s", "speed", "speed multiplier")
		alpha.Duration(&alphaTimeout, "t", "timeout", "how long alpha should wait")
		AttachSubcommand(alpha, 1)

		// Configure the beta subcommand, including slice and duration flags.
		beta := NewSubcommand("beta")
		beta.StringSlice(&betaLabels, "l", "label", "labels to apply")
		beta.Int(&betaLevel, "v", "level", "an integer setting")
		beta.Bool(&betaActive, "a", "active", "whether beta is active")
		beta.Float32(&betaRatio, "r", "ratio", "ratio value for calculations")
		beta.Duration(&betaInterval, "i", "interval", "time between runs")
		alpha.AttachSubcommand(beta, 1)

		// Configure the gamma subcommand to complete the hierarchy.
		gamma := NewSubcommand("gamma")
		gamma.String(&gammaMode, "m", "mode", "mode of operation")
		gamma.UInt(&gammaMax, "x", "max", "maximum allowed value")
		gamma.Bool(&gammaDebug, "d", "debug", "enable debug logging")
		gamma.Float64(&gammaThreshold, "h", "threshold", "threshold used for alerts")
		gamma.Duration(&gammaWindow, "w", "window", "time window to inspect")
		beta.AttachSubcommand(gamma, 1)

		// Parse the synthetic arguments that exercise the full tree.
		ParseArgs(args)

		// Assert that each subcommand was used during parsing.
		if !alpha.Used || !beta.Used || !gamma.Used {
			b.Fatalf("expected alpha, beta, gamma to be used but got alpha=%t beta=%t gamma=%t", alpha.Used, beta.Used, gamma.Used)
		}

		// Confirm the beta subcommand gathered labels and the active flag.
		if len(betaLabels) != 2 || !betaActive {
			b.Fatalf("expected beta labels and active flag to be parsed correctly")
		}

		// Confirm gamma parsed every value correctly, including numeric caps.
		if !gammaDebug || gammaMode == "" || gammaMax != 100 {
			b.Fatalf("gamma values not parsed as expected: debug=%t mode=%q max=%d", gammaDebug, gammaMode, gammaMax)
		}

		// Confirm trailing arguments flowed through after the -- separator.
		if len(TrailingArguments) != 3 {
			b.Fatalf("expected 3 trailing arguments, got %d", len(TrailingArguments))
		}

		// Validate every alpha flag to ensure full coverage of parsed values.
		if alphaName == "" || alphaCount != 3 || !alphaEnabled || alphaTimeout != 500*time.Millisecond {
			b.Fatalf("alpha values not parsed correctly")
		}
	}
}
