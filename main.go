// Package flaggy is a input flag parsing tool that supports
// subcommands and any-position flags without complexeties.
//
// Parsing Order:
//   - Parse and asign any flags found in the format -key=var,
//     --key=var, or '-key var'.  Remove these variables from
//     further consideration.
//   - Detect any positional values
//   - Detect any subcommands and parse them
//   - Repeat parsing order on subcommands until out of subcommands
//
package flaggy

import (
	"fmt"
	"os"
)

// DebugMode indicates that debug output should be enabled
var DebugMode bool

var mainParser *Parser

func init() {
	// allow usage like flaggy.StringVar by enabling a default Parser
	if len(os.Args) > 0 {
		mainParser = NewParser(os.Args[0])
	} else {
		mainParser = NewParser("default")
	}
}

// Parse parses flags as requested in the default package parser
func Parse() {
	mainParser.Parse()
}

// AddBoolFlag adds a bool flag for parsing, at the global level of the
// default parser
func AddBoolFlag(assignmentVar *bool, shortName string, longName string, description string) {
	mainParser.AddBoolFlag(assignmentVar, shortName, longName, description)
}

// AddIntFlag adds an int flag for parsing, at the global level of the
// default parser
func AddIntFlag(assignmentVar *int, shortName string, longName string, description string) {
	mainParser.AddIntFlag(assignmentVar, shortName, longName, description)
}

// AddStringFlag adds a string flag for parsing, at the global level of the
// default parser
func AddStringFlag(assignmentVar *string, shortName string, longName string, description string) {
	mainParser.AddStringFlag(assignmentVar, shortName, longName, description)
}

// AddSubcommand adds a subcommand for parsing
func AddSubcommand(newSC *Subcommand) error {
	return mainParser.AddSubcommand(newSC)
}

// AddPositionalValue adds a positional value to the main parser at the global
// context
func AddPositionalValue(relativePosition int, assignmentVar *string, name string, description string) error {
	return mainParser.AddPositionalValue(relativePosition, assignmentVar, name, description)
}

// debugPrint prints if debugging is enabled
func debugPrint(i ...interface{}) {
	if DebugMode {
		fmt.Println(i...)
	}
}
