package flaggy

// setValueForParsers sets the value for a specified key in the
// specified parsers (which normally include a Parser and Subcommand) until a
// parser accept takes the value, then it returns immediately.  The return
// values represent the key being set, and any errors returned when setting the
// key, such as failures to conver the string into the appropriate flag value.
func setValueForParsers(key string, value string, parsers ...ArgumentParser) (bool, error) {

	var valueWasSet bool
	var err error

	for _, p := range parsers {
		valueWasSet, err = p.SetValueForKey(key, value)
		if err != nil {
			return false, err
		}
	}

	return valueWasSet, nil
}

// ArgumentParser represements a parser or subcommand
type ArgumentParser interface {
	SetValueForKey(key string, value string) (bool, error)
}
