# flaggy

[![flaggy go report](https://goreportcard.com/badge/github.com/integrii/flaggy)](https://goreportcard.com/report/github.com/integrii/flaggy)

Sensible flag parsing with support for subcommands, positional values, and flags that can be at any position.  No required project or package layout like [Cobra](https://github.com/spf13/Cobra), and **no third party package dependencies**.  

Check out the [godoc](http://godoc.org/github.com/integrii/flaggy), [godoc examples](https://godoc.org/github.com/integrii/flaggy#pkg-examples), and [examples in this readme](https://github.com/integrii/flaggy#super-simple-example) to get started quickly.

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
print(flaggy.TrailingArguments)
```
