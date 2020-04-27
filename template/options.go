package template

import (
	"fmt"
	"reflect"
)

// ExtendedOption represents the additional options enabled by this template package.
type ExtendedOption uint32

// Options values
const (
	Razor ExtendedOption = 1 << iota
	Extension
	Math
	Sprig
	Data
	Logging
	Runtime
	Utils
	Net
	OS
	Git
	Overwrite
	OutputStdout
	RenderingDisabled
	AcceptNoValue
	StrictErrorCheck

	DefaultOptions = Git<<1 - 1
	AllOptions     = StrictErrorCheck<<1 - 1
)

//go:generate stringer -type=options -output generated_options.go

// IsSet checks if all options specified are set.
// If you want to test with logical OR, submitted options can be combined with binary or.
// If you want to test with logical AND, submitted options must be provided separately.
func (o ExtendedOption) IsSet(opts ...ExtendedOption) bool {
	for _, opt := range opts {
		if o&opt == 0 {
			return false
		}
	}
	return o != 0
}

// Set the bits on the object t
func (o *ExtendedOption) Set(opt ExtendedOption, on bool) *ExtendedOption {
	if on {
		*o |= opt
	} else {
		*o ^= opt
	}
	return o
}

// Strings generates a string with all enabled options separated by sep.
func (o ExtendedOption) Strings(sep ...interface{}) (result string) {
	var separator = "+"
	if len(sep) > 0 {
		separator = fmt.Sprint(sep...)
	}
	if (o != 0) && ((o & (o - 1)) == 0) {
		// There is only one byte set, so we don't have to
		// loop and we can delegate to the generated constants.
		return o.String()
	}
	for current := ExtendedOption(1); current <= AllOptions; current <<= 1 {
		if o&current != 0 {
			if result != "" {
				result += separator
			}
			result += current.String()
		}
	}
	return
}

// List returns an array with all options available.
func (o ExtendedOption) List() []ExtendedOption {
	result := make([]ExtendedOption, 0, reflect.TypeOf(o).Bits())
	for current := ExtendedOption(1); current <= ExtendedOption(AllOptions); current <<= 1 {
		if o&current != 0 {
			result = append(result, current)
		}
	}
	return result
}
