package flaggy

import (
	"fmt"
	"os"
)

// Parser represents the set of vars and subcommands we are expecting
// from our input args, and the parser than handles them all.
type Parser struct {
	Subcommand
	Version              string   // the optional version of the paser.
	ShowHelpWithHFlag    bool     // display help when -h or --help passed
	ShowVersionWithVFlag bool     // display the version when -v or --version passed
	ShowHelpOnUnexpected bool     // display help when an unexpected flag is passed
	TrailingArguments    []string // everything after a -- is placed here
}

// NewParser creates a new ArgumentParser ready to parse inputs
func NewParser(name string) *Parser {
	// this can not be done inline because of struct embedding
	p := &Parser{}
	p.Name = name
	p.Version = defaultVersion
	p.ShowHelpOnUnexpected = true
	p.ShowHelpWithHFlag = true
	p.ShowVersionWithVFlag = true
	p.SetHelpTemplate(DefaultHelpTemplate)
	return p
}

// ParseArgs parses as if the passed args were the os.Args, but without the
// binary at the 0 position in the array.  An error is returned if there
// is a low level issue converting flags to their proper type.  No error
// is returned for invalid arguments or missing require subcommands.
func (p *Parser) ParseArgs(args []string) error {
	debugPrint("Kicking off parsing with depth of 0 and args:", args)
	return p.parse(p, args, 0)
}

// ShowVersionAndExit shows the version of this parser
func (p *Parser) ShowVersionAndExit() {
	fmt.Println("Version:", p.Version)
	os.Exit(0)
}

// Parse calculates all flags and subcommands
func (p *Parser) Parse() error {

	err := p.ParseArgs(os.Args[1:])
	if err != nil {
		return err
	}
	TrailingArguments = p.TrailingArguments
	return nil

}
