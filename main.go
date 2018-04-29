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
	"net"
	"os"
	"strings"
	"time"
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

// AddStringFlag adds a new string flag
func AddStringFlag(assignmentVar *string, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddStringSliceFlag adds a new slice of strings flag
// Specify the flag multiple times to fill the slice
func AddStringSliceFlag(assignmentVar *[]string, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddBoolFlag adds a new bool flag
func AddBoolFlag(assignmentVar *bool, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddBoolSliceFlag adds a new slice of bools flag
// Specify the flag multiple times to fill the slice
func AddBoolSliceFlag(assignmentVar *[]bool, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddByteSliceFlag adds a new slice of bytes flag
// Specify the flag multiple times to fill the slice.  Takes hex as input.
func AddByteSliceFlag(assignmentVar *[]byte, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddDurationFlag adds a new time.Duration flag.
// Input format is described in time.ParseDuration().
// Example values: 1h, 1h50m, 32s
func AddDurationFlag(assignmentVar *time.Duration, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddDurationSliceFlag adds a new time.Duration flag.
// Input format is described in time.ParseDuration().
// Example values: 1h, 1h50m, 32s
// Specify the flag multiple times to fill the slice.
func AddDurationSliceFlag(assignmentVar *[]time.Duration, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddFloat32Flag adds a new float32 flag.
func AddFloat32Flag(assignmentVar *float32, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddFloat32SliceFlag adds a new float32 flag.
// Specify the flag multiple times to fill the slice.
func AddFloat32SliceFlag(assignmentVar *[]float32, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddFloat64Flag adds a new float64 flag.
func AddFloat64Flag(assignmentVar *float64, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddFloat64SliceFlag adds a new float64 flag.
// Specify the flag multiple times to fill the slice.
func AddFloat64SliceFlag(assignmentVar *[]float64, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddIntFlag adds a new int flag
func AddIntFlag(assignmentVar *int, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddIntSliceFlag adds a new int slice flag.
// Specify the flag multiple times to fill the slice.
func AddIntSliceFlag(assignmentVar *[]int, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUIntFlag adds a new uint flag
func AddUIntFlag(assignmentVar *uint, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUIntSliceFlag adds a new uint slice flag.
// Specify the flag multiple times to fill the slice.
func AddUIntSliceFlag(assignmentVar *[]uint, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUInt64Flag adds a new uint64 flag
func AddUInt64Flag(assignmentVar *uint64, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUInt64SliceFlag adds a new uint64 slice flag.
// Specify the flag multiple times to fill the slice.
func AddUInt64SliceFlag(assignmentVar *[]uint64, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUInt32Flag adds a new uint32 flag
func AddUInt32Flag(assignmentVar *uint32, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUInt32SliceFlag adds a new uint32 slice flag.
// Specify the flag multiple times to fill the slice.
func AddUInt32SliceFlag(assignmentVar *[]uint32, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUInt16Flag adds a new uint16 flag
func AddUInt16Flag(assignmentVar *uint16, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUInt16SliceFlag adds a new uint16 slice flag.
// Specify the flag multiple times to fill the slice.
func AddUInt16SliceFlag(assignmentVar *[]uint16, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUInt8Flag adds a new uint8 flag
func AddUInt8Flag(assignmentVar *uint8, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddUInt8SliceFlag adds a new uint8 slice flag.
// Specify the flag multiple times to fill the slice.
func AddUInt8SliceFlag(assignmentVar *[]uint8, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddInt64Flag adds a new int64 flag
func AddInt64Flag(assignmentVar *int64, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddInt64SliceFlag adds a new int64 slice flag.
// Specify the flag multiple times to fill the slice.
func AddInt64SliceFlag(assignmentVar *[]int64, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddInt32Flag adds a new int32 flag
func AddInt32Flag(assignmentVar *int32, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddInt32SliceFlag adds a new int32 slice flag.
// Specify the flag multiple times to fill the slice.
func AddInt32SliceFlag(assignmentVar *[]int32, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddInt16Flag adds a new int16 flag
func AddInt16Flag(assignmentVar *int16, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddInt16SliceFlag adds a new int16 slice flag.
// Specify the flag multiple times to fill the slice.
func AddInt16SliceFlag(assignmentVar *[]int16, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddInt8Flag adds a new int8 flag
func AddInt8Flag(assignmentVar *int8, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddInt8SliceFlag adds a new int8 slice flag.
// Specify the flag multiple times to fill the slice.
func AddInt8SliceFlag(assignmentVar *[]int8, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddIPFlag adds a new net.IP flag.
func AddIPFlag(assignmentVar *net.IP, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddIPSliceFlag adds a new int8 slice flag.
// Specify the flag multiple times to fill the slice.
func AddIPSliceFlag(assignmentVar *[]net.IP, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddHardwareAddrFlag adds a new net.HardwareAddr flag.
func AddHardwareAddrFlag(assignmentVar *net.HardwareAddr, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddHardwareAddrSliceFlag adds a new net.HardwareAddr slice flag.
// Specify the flag multiple times to fill the slice.
func AddHardwareAddrSliceFlag(assignmentVar *[]net.HardwareAddr, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddIPMaskFlag adds a new net.IPMask flag. IPv4 Only.
func AddIPMaskFlag(assignmentVar *net.IPMask, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddIPMaskSliceFlag adds a new net.HardwareAddr slice flag. IPv4 only.
// Specify the flag multiple times to fill the slice.
func AddIPMaskSliceFlag(assignmentVar *[]net.IPMask, shortName string, longName string, description string) error {
	return mainParser.addFlag(assignmentVar, shortName, longName, description)
}

// AddSubcommand adds a subcommand for parsing
func AddSubcommand(newSC *Subcommand, relativePosition int) error {
	return mainParser.AddSubcommand(newSC, relativePosition)
}

// ShowHelp shows parser help
func ShowHelp(message string) {
	mainParser.ShowHelpWithMessage(message)
}

// SetDescription sets the description of the default package command parser
func SetDescription(description string) {
	mainParser.Description = description
}

// SetVersion sets the version of the default package command parser
func SetVersion(version string) {
	mainParser.Version = version
}

// SetName sets the name of the default package command parser
func SetName(name string) {
	mainParser.Name = name
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
