// Package flaggy is a input flag parsing tool that supports both subcommands
// and any-position flags without unnecessary complexeties.
/*

Supported Flag Types

Strings and Ints
 -key=var
 --key=var
 --key var
 -key var

Booleans (sets to true if no var specified)
 --key
 --key var
 -k var
 -k


*/
package flaggy

import (
	"fmt"
	"os"
)

// DebugMode indicates that debug output should be enabled
var DebugMode bool

var mainParser *Parser

// TrailingArguments holds trailing arguments in the main parser after parsing
// has been run.
var TrailingArguments []string

func init() {
	// allow usage like flaggy.StringVar by enabling a default Parser
	if len(os.Args) > 0 {
		mainParser = NewParser(os.Args[0])
	} else {
		mainParser = NewParser("default")
	}
}

// Parse parses flags as requested in the default package parser
func Parse() error {
	return mainParser.Parse()
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
func AddSubcommand(newSC *Subcommand, relativePosition int) error {
	return mainParser.AddSubcommand(newSC, relativePosition)
}

// AddPositionalValue adds a positional value to the main parser at the global
// context
func AddPositionalValue(assignmentVar *string, name string, relativePosition int, description string) error {
	return mainParser.AddPositionalValue(assignmentVar, name, relativePosition, description)
}

// debugPrint prints if debugging is enabled
func debugPrint(i ...interface{}) {
	if DebugMode {
		fmt.Println(i...)
	}
}
