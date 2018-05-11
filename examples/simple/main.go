package main

import "github.com/integrii/flaggy"

func main() {
	// Declare variables and their defaults
	var stringFlag = "defaultValue"

	// Add a flag
	flaggy.String(&stringFlag, "f", "flag", "A test string flag")

	// Parse the flag
	flaggy.Parse()

	// Use the flag
	print(stringFlag)
}
