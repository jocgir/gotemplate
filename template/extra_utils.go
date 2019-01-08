package template

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/coveo/gotemplate/collections"
	"github.com/coveo/gotemplate/utils"
)

const (
	utilsBase = "Other utilities"
)

var utilsFuncs = dictionary{
	"center":     center,
	"color":      utils.SprintColor,
	"concat":     collections.Concat,
	"formatList": utils.FormatList,
	"id":         id,
	"IIf":        IIf,
	"joinLines":  collections.JoinLines,
	"lorem":      lorem,
	"mergeList":  utils.MergeLists,
	"repeat":     repeat,
	"indent":     indent,
	"nIndent":    nIndent,
	"sIndent":    sIndent,
	"splitLines": collections.SplitLines,
	"wrap":       wrap,
}

var utilsFuncsArgs = arguments{
	"center":     {"width"},
	"formatList": {"format", "list"},
	"id":         {"identifier", "replaceChar"},
	"IIf":        {"testValue", "valueTrue", "valueFalse"},
	"joinLines":  {"format"},
	"lorem":      {"loremType", "params"},
	"mergeList":  {"lists"},
	"repeat":     {"n", "element"},
	"indent":     {"nbSpace"},
	"nIndent":    {"nbSpace"},
	"sIndent":    {"spacer"},
	"splitLines": {"content"},
	"wrap":       {"width"},
}

var utilsFuncsAliases = aliases{
	"center":     {"centered"},
	"color":      {"colored", "enhanced"},
	"id":         {"identifier"},
	"IIf":        {"ternary"},
	"lorem":      {"loremIpsum"},
	"nIndent":    {"nindent"},
	"formatList": {"autoWrap", "aWrap", "awrap"},
	"sIndent":    {"sindent", "spaceIndent", "autoIndent", "aindent", "aIndent"},
	"wrap":       {"wrapped"},
}

var utilsFuncsHelp = descriptions{
	"center": "Returns the concatenation of supplied arguments centered within width.",
	"color": String(`
		Colors the rendered string.

		The first arguments are interpretated as color attributes until the first non color attribute. Attributes are case insensitive.

		Valid attributes are:
		    Reset, Bold, Faint, Italic, Underline, BlinkSlow, BlinkRapid, ReverseVideo, Concealed, CrossedOut

		Valid color are:
		    Black, Red, Green, Yellow, Blue, Magenta, Cyan, White

		Color can be prefixed by:
		    Fg:   Meaning foreground (Fg is assumed if not specified)
		    FgHi: Meaning high intensity forground
		    Bg:   Meaning background"
		    BgHi: Meaning high intensity background
	`).UnIndent().TrimSpace(),
	"concat": "Returns the concatenation (without separator) of the string representation of objects.",
	"formatList": String(`
		Return a list of strings by applying the format to each element of the supplied list.

		You can also use autoWrap as Razor expression if you don't want to specify the format.
		The format is then automatically induced by the context around the declaration).
		Valid aliases for autoWrap are: aWrap, awrap.

		Ex:
		    Hello @<autoWrap(to(10)) World!
	`).UnIndent().TrimSpace(),
	"id":        "Returns a valid go identifier from the supplied string (replacing any non compliant character by replacement, default _ ).",
	"IIf":       "If testValue is empty, returns falseValue, otherwise returns trueValue.\n    WARNING: All arguments are evaluated and must by valid.",
	"indent":    "Indents every line in a given string to the specified indent width. This is useful when aligning multi-line strings.",
	"joinLines": "Merge the supplied objects into a newline separated string.",
	"lorem":     "Returns a random string. Valid types are be word, words, sentence, para, paragraph, host, email, url.",
	"mergeList": "Return a single list containing all elements from the lists supplied.",
	"nindent":   "Work as indent but add a newline before.",
	"repeat":    "Returns an array with the item repeated n times.",
	"sIndent": String(`
		Indents the elements using the provided spacer.
		
		You can also use autoIndent as Razor expression if you don't want to specify the spacer.
		Spacer will then be auto determined by the spaces that precede the expression.
		Valid aliases for autoIndent are: aIndent, aindent.
	`).UnIndent().TrimSpace(),
	"splitLines": "Returns a list of strings from the supplied object with newline as the separator.",
	"wrap":       "Wraps the rendered arguments within width.",
}

func (t *Template) addUtilsFuncs() {
	t.AddFunctions(utilsFuncs, utilsBase, FuncOptions{
		FuncHelp:    utilsFuncsHelp,
		FuncArgs:    utilsFuncsArgs,
		FuncAliases: utilsFuncsAliases,
	})
}

func lorem(funcName interface{}, params ...int) (result String, err error) {
	kind, err := utils.GetLoremKind(asStdString(funcName))
	if err == nil {
		result, err = utils.Lorem(kind, params...)
	}
	return
}

func center(width interface{}, args ...interface{}) (String, error) {
	w, err := strconv.Atoi(fmt.Sprintf("%v", width))
	if err != nil {
		return "", fmt.Errorf("width must be integer")
	}
	return FormatMessage(args...).Center(w), nil
}

func wrap(width interface{}, args ...interface{}) (String, error) {
	w, err := strconv.Atoi(fmt.Sprintf("%v", width))
	if err != nil {
		return "", fmt.Errorf("width must be integer")
	}
	return FormatMessage(args...).Wrap(w), nil
}

func indent(space int, args ...interface{}) String {
	args = convertArgs(nil, args...).AsArray()
	return String("\n").Join(args...).Indent(strings.Repeat(" ", space))
}

func nIndent(space int, args ...interface{}) String {
	return "\n" + indent(space, args...)
}

func sIndent(spacer string, args ...interface{}) String {
	args = convertArgs(nil, args...).AsArray()
	return String("\n").Join(args...).Indent(spacer)
}

func id(id string, replace ...interface{}) String {
	// By default, replacement char for invalid chars would be _
	// but it is possible to specify an alternative string to act as the replacement
	replacement := fmt.Sprint(replace...)
	if replacement == "" {
		replacement = "_"
	}

	dup := duplicateUnderscore
	if replacement != "_" {
		// If the replacement string is not the default one, we generate a special substituter to remove duplicates
		// taking into account regex special chars such as +, ?, etc.
		dup = regexp.MustCompile(fmt.Sprintf(`(?:%s)+`, regexSpecial.ReplaceAllString(replacement, `\$0`)))
	}

	return String(dup.ReplaceAllString(validChars.ReplaceAllString(id, replacement), replacement))
}

var validChars = regexp.MustCompile(`[^\p{L}\d_]`)
var duplicateUnderscore = regexp.MustCompile(`_+`)
var regexSpecial = regexp.MustCompile(`[\+\.\?\(\)\[\]\{\}\\]`)

func repeat(n int, a interface{}) (result IGenericList, err error) {
	if n < 0 {
		err = fmt.Errorf("n must be greater or equal than 0")
		return
	}
	result = collections.CreateList(n)
	for i := 0; i < n; i++ {
		result.Set(i, a)
	}
	return
}
