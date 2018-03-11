package flaggy

// setValueForParsers sets the value for a specified key in the
// specified parsers (which include Parser and Subcommand) until a
// parser accepts takes the value, then it returns
func setValueForParsers(key string, value string, parsers ...ArgumentParser) error {
	for _, p := range parsers {
		valueSet, err := p.SetValueForKey(key, value)
		if err != nil {
			return err
		}
		// dont continue setting past the first parser
		if valueSet {
			return nil
		}
	}
	return nil
}

// ArgumentParser represements a parser or subcommand
type ArgumentParser interface {
	SetValueForKey(key string, value string) (bool, error)
}
