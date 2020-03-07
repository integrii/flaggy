package flaggy

import (
	"errors"
	"reflect"
	"strconv"
)

// PositionalValue represents a value which is determined by its position
// relative to where a subcommand was detected.
type PositionalValue struct {
	Name          string // used in documentation only
	Description   string
	AssignmentVar interface{}
	Position      int    // the position, not including switches, of this variable
	Required      bool   // this subcommand must always be specified
	Found         bool   // was this positional found during parsing?
	Hidden        bool   // indicates this positional value should be hidden from help
	defaultValue  string // used for help output
	parsed        bool   // indicates that this positional has already been parsed
}

// identifyAndAssignValue identifies the type of the incoming value
// and assigns it to the AssignmentVar pointer's target value.  If
// the value is a type that needs parsing, that is performed as well.
func (p *PositionalValue) identifyAndAssignValue(value string) error {
	var err error

	// Only parse this positional default value once. This keeps us from
	// overwriting the default value in help output
	if !p.parsed {
		p.parsed = true
		// parse the default value as a string and remember it for help output
		p.defaultValue, err = p.returnAssignmentVarValueAsString()
		if err != nil {
			return err
		}
	}

	debugPrint("attempting to assign value", value, "to positional", p.Name)

	// depending on the type of the assignment variable, we convert the
	// incoming string and assign it.  We only use pointers to variables
	// in flaggy. No returning vars by value.
	switch p.AssignmentVar.(type) {
	case *string:
		v, _ := (p.AssignmentVar).(*string)
		*v = value
	case *bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		a, _ := (p.AssignmentVar).(*bool)
		*a = v
	case *int:
		v, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		e := p.AssignmentVar.(*int)
		*e = v
	default:
		return errors.New("Unknown positional assignmentVar supplied in " + p.Name)
	}

	return err
}

// returnAssignmentVarValueAsString returns the value of the positional's
// assignment variable as a string. This is used to display the default
// value of positionals before they are assigned (like when help is output).
func (p *PositionalValue) returnAssignmentVarValueAsString() (string, error) {
	debugPrint("returning current value of assignment var of positional", p.Name)

	var err error

	// depending on the type of the assignment variable, we convert the
	// incoming string and assign it. We only use pointers to variables
	// in flaggy. No returning vars by value.
	switch p.AssignmentVar.(type) {
	case *string:
		v, _ := (p.AssignmentVar).(*string)
		return *v, err
	case *bool:
		a, _ := (p.AssignmentVar).(*bool)
		return strconv.FormatBool(*a), err
	case *int:
		a := p.AssignmentVar.(*int)
		return strconv.Itoa(*a), err
	default:
		return "", errors.New("Unknown positional assignmentVar found in " + p.Name + ". " +
			"Type not supported: " + reflect.TypeOf(p.AssignmentVar).String())
	}
}
