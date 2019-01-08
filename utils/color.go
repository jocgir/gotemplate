package utils

import (
	"fmt"
	"strings"

	"github.com/coveo/gotemplate/collections"
	"github.com/coveo/gotemplate/errors"
	"github.com/fatih/color"
)

// Attribute is imported from color attribute
type Attribute color.Attribute

// The following constant are copied from the color package in order to get
// the actual names
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground attributes
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground attributes high intensity
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background attributes
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background attributes high intensity
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

//go:generate stringer -type=Attribute -output generated_colors.go

// Color returns a color attribute build from supplied attribute names
func Color(attributes ...string) (*color.Color, error) {
	if nameValues == nil {
		nameValues = make(map[string]color.Attribute, BgHiWhite)
		for i := Reset; i < BgHiWhite; i++ {
			name := strings.ToLower(Attribute(i).String())
			if strings.HasPrefix(name, "attribute(") {
				continue
			}
			nameValues[name] = color.Attribute(i)
			if strings.HasPrefix(name, "fg") {
				nameValues[name[2:]] = color.Attribute(i)
			}
		}
	}

	result := color.New()
	var containsColor bool
	var err errors.Array
	for _, attr := range attributes {
		for _, attr := range String(attr).FieldsID().Strings() {
			if a, match := nameValues[attr.ToLower().Str()]; match {
				result.Add(a)
				containsColor = true

			} else {
				err = append(err, fmt.Errorf("Attribute not found %s", attr))
			}
		}
	}

	if !containsColor {
		return result, fmt.Errorf("No color specified")
	}
	if len(err) > 0 {
		return result, err
	}
	return result, nil
}

// SprintColor returns a string formated with attributes that are supplied before
func SprintColor(args ...interface{}) (string, error) {
	var i int
	colorArgs := make([]string, len(args))
	for i = 0; i < len(args); i++ {
		if _, err := Color(fmt.Sprint(args[i])); err != nil {
			break
		}
		colorArgs[i] = fmt.Sprint(args[i])
	}

	c, _ := Color(colorArgs...)
	return c.Sprint(collections.FormatMessage(args[i:]...)), nil
}

var nameValues map[string]color.Attribute

// ColorPrintln call standard fmt.Println function but using the color out stream.
func ColorPrintln(args ...interface{}) (int, error) {
	return fmt.Fprintln(color.Output, args...)
}

// ColorPrintf call standard fmt.Printf function but using the color out stream.
func ColorPrintf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(color.Output, format, args...)
}

// ColorPrint call standard fmt.Printf function but using the color out stream.
func ColorPrint(args ...interface{}) (int, error) {
	return fmt.Fprint(color.Output, args...)
}

// ColorErrorPrintln call standard fmt.Println function but using the color out stream.
func ColorErrorPrintln(args ...interface{}) (int, error) {
	return fmt.Fprintln(color.Error, args...)
}

// ColorErrorPrintf call standard fmt.Printf function but using the color out stream.
func ColorErrorPrintf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(color.Error, format, args...)
}

// ColorErrorPrint call standard fmt.Printf function but using the color out stream.
func ColorErrorPrint(args ...interface{}) (int, error) {
	return fmt.Fprint(color.Error, args...)
}
