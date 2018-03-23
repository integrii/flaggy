![flaggy logo](/logo.png)

[![flaggy go report](https://goreportcard.com/badge/github.com/integrii/flaggy)](https://goreportcard.com/report/github.com/integrii/flaggy) [![Build Status](https://travis-ci.org/integrii/flaggy.svg?branch=master)](https://travis-ci.org/integrii/flaggy) [![godoc](https://camo.githubusercontent.com/d48cccd1ce67ddf8ba7fc356ec1087f3f7aa6d12/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f6c696c65696f2f6c696c653f7374617475732e737667)](http://godoc.org/github.com/integrii/flaggy) [![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](http://unlicense.org/)



Sensible flag parsing with support for subcommands, positional values, and flags that can be at any position.  No required project or package layout like [Cobra](https://github.com/spf13/Cobra), and **no third party package dependencies**.  

Check out the [godoc](http://godoc.org/github.com/integrii/flaggy), [examples directory](https://github.com/integrii/flaggy/tree/master/examples), and [examples in this readme](https://github.com/integrii/flaggy#super-simple-example) to get started quickly.

Open an issue if you hate something, or better yet, fix it and make a pull request!

# Key Features

- Very easy to use ([see examples below](https://github.com/integrii/flaggy#super-simple-example))
- Any flag can be at at any position
- Pretty and readable help output by default
- Positional subcommands
- Positional parameters
- Suggested subcommands when a subcommand is typo'd
- Nested subcommands
- Both global and subcommand specific flags
- Both global and subcommand specific positional parameters
- Customizable help templates for both the global command and subcommands
- Customizable appended/prepended message for both the global command and subcommands
- Simple function that displays help followed by a custom message string
- Flags and subcommands may have both a short and long name
- Unlimited trailing arguments after a `--`
- Flags can use a single dash or double dash (`--flag`, `-flag`, `-f`, `--f`)
- Flags can have `=` assignment operators, or use a space (`--flag=value`, `--flag value`)
- Flags support single quote globs with spaces (`--flag 'this is all one value'`)
- Optional but default version output with `-v` or `--version`
- Optional but default help output with `-h` or `--help`
- Optional but default show help when any invalid or unknown parameter is passed


# Example Help Output

```
testCommand - Description goes here.  Get more information at http://my.website
This is an optional prepend for help output

  Positional Variables:
    testPositionalA (Position 2) (Required) Test positional A does some things with a positional value.

  Subommands:
    subcommandA (a) (Position 1) Subcommand A is a command that does stuff
    subcommandB (b) (Position 1) Subcommand B is a command that does other stuff
    subcommandC (c) (Position 1) Subcommand C is a command that does SERIOUS stuff

 Flags:
    --stringFlag (-s) This is a test string flag that does some stringy string stuff.
    --intFlg (-i) This is a test int flag that does some interesting int stuff.
    --boolFlag (-b) This is a test bool flag that does some booly bool stuff.

This is an optional append for help
This is an optional help add-on message
```


# Super Simple Example

`./yourApp -f test`

```go
// Declare variables and their defaults
var stringFlag = "defaultValue"

// Add a flag
flaggy.AddStringFlag(&stringFlag, "f", "flag", "A test string flag")

// Parse the flag
flaggy.Parse()

// Use the flag
print(stringFlag)
```


# Example with Subcommand

`./yourApp subcommandExample -f test`

```go
// Declare variables and their defaults
var stringFlag = "defaultValue"

// Create the subcommand
subcommand := flaggy.NewSubcommand("subcommandExample")

// Add a flag to the subcommand
subcommand.AddStringFlag(&stringFlag, "f", "flag", "A test string flag")

// Add the subcommand to the parser at position 1
flaggy.AddSubcommand(subcommand, 1)

// Parse the subcommand and all flags
flaggy.Parse()

// Use the flag
print(stringFlag)
```

# Example with Nested Subcommands, Various Flags and Trailing Arguments

`./yourApp subcommandExample --flag=5 nestedSubcommand -t test -y -- trailingArg`

```go
// Declare variables and their defaults
var stringFlagF = "defaultValueF"
var intFlagT = 3
var boolFlagB bool

// Create the subcommand
subcommandExample := flaggy.NewSubcommand("subcommandExample")
nestedSubcommand := flaggy.NewSubcommand("nestedSubcommand")

// Add a flag to the subcommand
subcommandExample.AddStringFlag(&stringFlagF, "t", "testFlag", "A test string flag")

nestedSubcommand.AddIntFlag(&intFlagT, "f", "flag", "A test int flag")

// add a global bool flag for fun
flaggy.AddBoolFlag(&boolFlagB, "y", "yes", "A sample boolean flag")

// Add the nested subcommand to the parent subcommand at position 1
subcommandExample.AddSubcommand(nestedSubcommand, 1)
// Add the base subcommand to the parser at position 1
flaggy.AddSubcommand(subcommandExample, 1)

// Parse the subcommand and all flags
flaggy.Parse()

// Use the flags and trailing arguments
print(stringFlagF)
print(intFlagT)
print(boolFlagB)
print(flaggy.TrailingArguments[0])
```
