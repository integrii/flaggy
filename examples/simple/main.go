package main

import (
	"fmt"
	"github.com/integrii/flaggy"
)

func main() {
	var stringValue = "abcdef"
	var boolValue = false
	var intValue = 123456

	flaggy.PositionalString(&stringValue, "str", 1, true, "String value")
	flaggy.PositionalBool(&boolValue, "bool", 2, true, "Boolean value")
	flaggy.PositionalInt(&intValue, "int", 3, true, "Integer value")
	flaggy.Parse()

	fmt.Printf("%T: %v\n", stringValue, stringValue)
	fmt.Printf("%T: %v\n", boolValue, boolValue)
	fmt.Printf("%T: %v\n", intValue, intValue)
}
