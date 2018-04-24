package flaggy

// setValueForParsers sets the value for a specified key in the
// specified parsers (which normally include a Parser and Subcommand).
// The return values represent the key being set, and any errors
// returned when setting the key, such as failures to conver the string
// into the appropriate flag value.
func setValueForParsers(key string, value string, parsers ...ArgumentParser) (bool, error) {

	var valueWasSet bool

	for _, p := range parsers {
		valueWasSetKey, err := p.SetValueForKey(key, value)
		if err != nil {
			return valueWasSet, err
		}
		if valueWasSetKey {
			valueWasSet = true
		}
	}

	return valueWasSet, nil
}

// ArgumentParser represements a parser or subcommand
type ArgumentParser interface {
	SetValueForKey(key string, value string) (bool, error)
}
