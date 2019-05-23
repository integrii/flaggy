package main

import "github.com/integrii/flaggy"

func main() {
	// Declare variables and their defaults
	var stringSliceFlag []string
	var boolSliceFlag []bool

	// Add a slice flag
	flaggy.DefaultParser.AdditionalHelpAppend = "Example: ./sliceFlag -b -b -s one -s two -b=false"
	flaggy.StringSlice(&stringSliceFlag, "s", "string", "A test string slice flag")
	flaggy.BoolSlice(&boolSliceFlag, "b", "bool", "A test bool slice flag")

	// Parse the flag
	flaggy.Parse()

	// output the flag contents
	for i := range stringSliceFlag {
		println(stringSliceFlag[i])
	}

	for i := range boolSliceFlag {
		println(boolSliceFlag[i])
	}

	// ./sliceFlag -b -b -s one -s two -b=false
	// output:
	// one
	// two
	// true
	// true
	// false
}
