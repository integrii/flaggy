package flaggy

import "os"

// Parser represents the set of vars and subcommands we are expecting
// from our input args, and the parser than handles them all.
type Parser struct {
	Subcommand
	TrailingArguments []string // everything after a -- is placed here
}

// NewParser creates a new ArgumentParser ready to parse inputs
func NewParser(name string) *Parser {
	p := &Parser{}
	p.Name = name
	return p
}

// ParseArgs parses as if the passed args were the os.Args, but without the
// binary at the 0 position in the array.
func (p *Parser) ParseArgs(args []string) error {
	debugPrint("Kicking off parsing with depth of 0 and args:", args)
	return p.parse(p, args, 0)
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
