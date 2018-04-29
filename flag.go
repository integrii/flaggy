package flaggy

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// Flag holds the base methods for all flag types
type Flag struct {
	ShortName     string
	LongName      string
	Description   string
	Hidden        bool // indicates this flag should be hidden from help and suggestions
	AssignmentVar interface{}
}

// HasName indicates that this flag's short or long name matches the
// supplied name string
func (f *Flag) HasName(name string) bool {
	name = strings.TrimSpace(name)
	if f.ShortName == name || f.LongName == name {
		return true
	}
	return false
}

// identifyAndAssignValue identifies the type of the incoming value
// and assigns it to the AssignmentVar pointer's target value.  If
// the value is a type that needs parsing, that is performed as well.
func (f *Flag) identifyAndAssignValue(value string) error {

	fmt.Println("attempting to assign value", value, "to flag", f.LongName)

	var err error

	// depending on the type of the assignment variable, we convert the
	// incoming string and assign it.  We only use pointers to variables
	// in flagy.  No returning vars by value.
	switch f.AssignmentVar.(type) {
	case *string:
		f.AssignmentVar = &value
		fmt.Println("after", f.AssignmentVar)
		fmt.Println("twice")
	case *[]string:
		v := f.AssignmentVar.(*[]string)
		a := append(*v, value)
		f.AssignmentVar = &a
	case *bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		f.AssignmentVar = &v
	case *[]bool:
		// parse the incoming bool
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		// cast the assignment var
		existing := f.AssignmentVar.(*[]bool)
		// deref the assignment var and append to it
		v := append(*existing, b)
		// pointer the new value and assign it
		f.AssignmentVar = &v
	case *time.Duration:
		v, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		f.AssignmentVar = &v
	case *[]time.Duration:
		t, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		existing := f.AssignmentVar.(*[]time.Duration)
		// deref the assignment var and append to it
		v := append(*existing, t)
		// pointer the new value and assign it
		f.AssignmentVar = &v
	case *float32:
		v, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		float := float32(v)
		f.AssignmentVar = &float
	case *[]float32:
		v, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		existing := f.AssignmentVar.(*[]float32)
		float := float32(v)
		new := append(*existing, float)
		f.AssignmentVar = &new
	case *float64:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		f.AssignmentVar = &v
	case *[]float64:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		existing := f.AssignmentVar.(*[]float64)
		new := append(*existing, v)
		f.AssignmentVar = &new
	case *int:
		v, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		f.AssignmentVar = &v
	case *[]int:
		v, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		existing := f.AssignmentVar.(*[]int)
		new := append(*existing, v)
		f.AssignmentVar = &new
	case *uint:
		v, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		val := uint(v)
		f.AssignmentVar = &val
	case *[]uint:
		v, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		existing := f.AssignmentVar.(*[]uint)
		new := append(*existing, uint(v))
		f.AssignmentVar = &new
	case *uint64:
		v, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		f.AssignmentVar = &v
	case *[]uint64:
		v, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		existing := f.AssignmentVar.(*[]uint64)
		new := append(*existing, v)
		f.AssignmentVar = &new
	case *uint32:
		v, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		val := uint32(v)
		f.AssignmentVar = &val
	case *[]uint32:
		v, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		existing := f.AssignmentVar.(*[]uint32)
		new := append(*existing, uint32(v))
		f.AssignmentVar = &new
	case *uint16:
		v, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
		val := uint16(v)
		f.AssignmentVar = &val
	case *[]uint16:
		v, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
		existing := f.AssignmentVar.(*[]uint16)
		new := append(*existing, uint16(v))
		f.AssignmentVar = &new
	case *uint8:
		v, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return err
		}
		val := uint8(v)
		f.AssignmentVar = &val
	case *[]uint8:
		// parse a hex string to a slice of bytes
		src := []byte(value)
		dst := make([]byte, hex.DecodedLen(len(src)))
		_, err := hex.Decode(dst, src)
		if err != nil {
			return err
		}
		f.AssignmentVar = &dst
	case *int64:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		f.AssignmentVar = &v
	case *[]int64:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		existingSlice := f.AssignmentVar.(*[]int64)
		newSlice := append(*existingSlice, v)
		f.AssignmentVar = &newSlice
	case *int32:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		converted := int32(v)
		f.AssignmentVar = &converted
	case *[]int32:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		existingSlice := f.AssignmentVar.(*[]int32)
		newSlice := append(*existingSlice, int32(v))
		f.AssignmentVar = &newSlice
	case *int16:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		converted := int16(v)
		f.AssignmentVar = &converted
	case *[]int16:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		existingSlice := f.AssignmentVar.(*[]int16)
		newSlice := append(*existingSlice, int16(v))
		f.AssignmentVar = &newSlice
	case *int8:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		converted := int8(v)
		f.AssignmentVar = &converted
	case *[]int8:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		existingSlice := f.AssignmentVar.(*[]int8)
		newSlice := append(*existingSlice, int8(v))
		f.AssignmentVar = &newSlice
	case *net.IP:
		v := net.ParseIP(value)
		f.AssignmentVar = &v
	case *[]net.IP:
		v := net.ParseIP(value)
		existing := f.AssignmentVar.(*[]net.IP)
		new := append(*existing, v)
		f.AssignmentVar = &new
	case *net.HardwareAddr:
		v, err := net.ParseMAC(value)
		if err != nil {
			return err
		}
		f.AssignmentVar = &v
	case *[]net.HardwareAddr:
		v, err := net.ParseMAC(value)
		if err != nil {
			return err
		}
		existing := f.AssignmentVar.(*[]net.HardwareAddr)
		new := append(*existing, v)
		f.AssignmentVar = &new
	case *net.IPMask:
		v := net.IPMask(net.ParseIP(value).To4())
		f.AssignmentVar = &v
	case *[]net.IPMask:
		v := net.IPMask(net.ParseIP(value).To4())
		existing := f.AssignmentVar.(*[]net.IPMask)
		new := append(*existing, v)
		f.AssignmentVar = &new
	default:
		return errors.New("Unknown flag assignmentVar supplied in flag " + f.LongName + " " + f.ShortName)
	}

	return err
}

const argIsPositional = "positional"       // subcommand or positional value
const argIsFlagWithSpace = "flagWithSpace" // -f path or --file path
const argIsFlagWithValue = "flagWithValue" // -f=path or --file=path
const argIsFinal = "final"                 // the final argument only '--'

// determineArgType determines if the specified arg is a flag with space
// separated value, a flag with a connected value, or neither (positional)
func determineArgType(arg string) string {

	// if the arg is --, then its the final arg
	if arg == "--" {
		return argIsFinal
	}

	// if it has the prefix --, then its a long flag
	if strings.HasPrefix(arg, "--") {
		// if it contains an equals, it is a joined value
		if strings.Contains(arg, "=") {
			return argIsFlagWithValue
		}
		return argIsFlagWithSpace
	}

	// if it has the prefix -, then its a short flag
	if strings.HasPrefix(arg, "-") {
		// if it contains an equals, it is a joined value
		if strings.Contains(arg, "=") {
			return argIsFlagWithValue
		}
		return argIsFlagWithSpace
	}

	return argIsPositional
}

// parseArgWithValue parses a key=value concatentated argument
func parseArgWithValue(arg string) (key string, value string) {

	// remove up to two minuses from start of flag
	arg = strings.TrimPrefix(arg, "-")
	arg = strings.TrimPrefix(arg, "-")

	// debugPrint("parseArgWithValue parsing", arg)

	// break at the equals
	args := strings.SplitN(arg, "=", 2)

	// if its a bool arg, with no explicit value, we return a blank
	if len(args) == 1 {
		return args[0], ""
	}

	// if its a key and value pair, we return those
	if len(args) == 2 {
		// debugPrint("parseArgWithValue parsed", args[0], args[1])
		return args[0], args[1]
	}

	fmt.Println("Warning: attempted to parseArgWithValue but did not have correct parameter count.", arg, "->", args)
	return "", ""
}

// parseFlagToName parses a flag with space value down to a key name:
//     --path -> path
//     -p -> p
func parseFlagToName(arg string) string {
	// remove minus from start
	arg = strings.TrimLeft(arg, "-")
	arg = strings.TrimLeft(arg, "-")
	return arg
}

// flagIsBool determines if the flag is a bool within the specified parser
// and subcommand's context
func flagIsBool(sc *Subcommand, p *Parser, key string) bool {
	for _, f := range sc.Flags {
		if f.HasName(key) {
			_, isBool := f.AssignmentVar.(*bool)
			_, isBoolSlice := f.AssignmentVar.(*[]bool)
			if isBool || isBoolSlice {
				return true
			}
		}
	}

	// by default, the answer is false
	return false
}
