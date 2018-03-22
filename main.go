// Package flaggy is a input flag parsing tool that supports subcommands
// positional values, and any-position flags without unnecessary complexeties.
/*

Supported Flag Types

Strings and Ints
 -key=var
 --key=var
 --key var
 -key var

Booleans (sets to true if flag is specified without value)
 --key
 --key true
 --key=false
 -k false
 -k=true
 -k
 
 All arguments after a double dash (--) are added as strings to the 
 TrailingArguments slice.


*/
package flaggy

import (
	"fmt"
	"os"
	"strings"
)

// defaultVersion is applied to parsers when they are created
const defaultVersion = "0.0.0"

// DebugMode indicates that debug output should be enabled
var DebugMode bool

// DefaultHelpTemplate is the help template that will be used
// for newly created subcommands and commands
var DefaultHelpTemplate string

var mainParser *Parser

// TrailingArguments holds trailing arguments in the main parser after parsing
// has been run.
var TrailingArguments []string

func init() {

	// Users may set DefaultHelpTemplate to save repeated template
	// assignment on every subcommand
	DefaultHelpTemplate = defaultHelpTemplate

	// set the default help template
	// allow usage like flaggy.StringVar by enabling a default Parser
	ResetParser()
}

// ResetParser resets the main default parser to a fresh instance.
// Normally used in tests.
func ResetParser() {
	if len(os.Args) > 0 {
		chunks := strings.Split(os.Args[0], "/")
		mainParser = NewParser(chunks[len(chunks)-1])
	} else {
		mainParser = NewParser("default")
	}
}

// Parse parses flags as requested in the default package parser
func Parse() error {
	err := mainParser.Parse()
	TrailingArguments = mainParser.TrailingArguments
	return err
}

// ParseArgs parses the passed args as if they were the arguments to the
// running binary.  Targets the default main parser for the package.
func ParseArgs(args []string) error {
	err := mainParser.ParseArgs(args)
	TrailingArguments = mainParser.TrailingArguments
	return err
}

// AddBoolFlag adds a bool flag for parsing, at the global level of the
// default parser
func AddBoolFlag(assignmentVar *bool, shortName string, longName string, description string) error {
	return mainParser.AddBoolFlag(assignmentVar, shortName, longName, description)
}

// AddIntFlag adds an int flag for parsing, at the global level of the
// default parser
func AddIntFlag(assignmentVar *int, shortName string, longName string, description string) error {
	return mainParser.AddIntFlag(assignmentVar, shortName, longName, description)
}

// AddStringFlag adds a string flag for parsing, at the global level of the
// default parser
func AddStringFlag(assignmentVar *string, shortName string, longName string, description string) error {
	return mainParser.AddStringFlag(assignmentVar, shortName, longName, description)
}

// AddSubcommand adds a subcommand for parsing
func AddSubcommand(newSC *Subcommand, relativePosition int) error {
	return mainParser.AddSubcommand(newSC, relativePosition)
}

// ShowHelp shows parser help
func ShowHelp(message string) {
	mainParser.ShowHelpWithMessage(message)
}

// ShowHelpAndExit shows parser help and exits with status code 2
func ShowHelpAndExit(message string) {
	ShowHelp(message)
	os.Exit(2)
}

// AddPositionalValue adds a positional value to the main parser at the global
// context
func AddPositionalValue(assignmentVar *string, name string, relativePosition int, required bool, description string) error {
	return mainParser.AddPositionalValue(assignmentVar, name, relativePosition, required, description)
}

// debugPrint prints if debugging is enabled
func debugPrint(i ...interface{}) {
	if DebugMode {
		fmt.Println(i...)
	}
}
