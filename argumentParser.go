package flaggy

// ArgumentParser represents the set of vars and subcommands we are expecting
// from our input args.
type ArgumentParser struct {
	Subcommand
	TrailingArguments []string // everything after a -- is placed here
}

// NewArgumentParser creates a new ArgumentParser ready to parse inputs
func NewArgumentParser() *ArgumentParser {
	return &ArgumentParser{}
}

// Parse calculates all flags and subcommands
func (ap *ArgumentParser) Parse() error {
	for _, command := range ap.Subcommands {
		err := command.parse(ap, 1) // initial depth of parsing is one command deep
		if err != nil {
			return err
		}
	}
	return nil
}
