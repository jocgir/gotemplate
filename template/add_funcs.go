package template

import (
	"sort"

	"github.com/coveo/gotemplate/utils"
)

const (
	goTemplateBse = "Base go template functions"
)

var baseGoTemplateFuncs = funcTableMap{
	"and":      {nil, goTemplateBse, nil, []string{}, `Returns the boolean AND of its arguments by returning the first empty argument or the last argument, that is, "and x y" behaves as "if x then y else x". All the arguments are evaluated.`},
	"call":     {nil, goTemplateBse, nil, []string{}, `Returns the result of calling the first argument, which must be a function, with the remaining arguments as parameters. Thus "call .X.Y 1 2" is, in Go notation, dot.X.Y(1, 2) where Y is a func-valued field, map entry, or the like. The first argument must be the result of an evaluation that yields a value of function type (as distinct from a predefined function such as print). The function must return either one or two result values, the second of which is of type error. If the arguments don't match the function or the returned error value is non-nil, execution stops.`},
	"html":     {nil, goTemplateBse, nil, []string{}, `Returns the escaped HTML equivalent of the textual representation of its arguments. This function is unavailable in html/template, with a few exceptions.`},
	"index":    {nil, goTemplateBse, nil, []string{}, `Returns the result of indexing its first argument by the following arguments. Thus "index x 1 2 3" is, in Go syntax, x[1][2][3]. Each indexed item must be a map, slice, or array.`},
	"js":       {nil, goTemplateBse, nil, []string{}, `Returns the escaped JavaScript equivalent of the textual representation of its arguments.`},
	"len":      {nil, goTemplateBse, nil, []string{}, `Returns the integer length of its argument.`},
	"not":      {nil, goTemplateBse, nil, []string{}, `Returns the boolean negation of its single argument.`},
	"or":       {nil, goTemplateBse, nil, []string{}, `Returns the boolean OR of its arguments by returning the first non-empty argument or the last argument, that is, "or x y" behaves as "if x then x else y". All the arguments are evaluated.`},
	"print":    {nil, goTemplateBse, nil, []string{}, `An alias for fmt.Sprint`},
	"printf":   {nil, goTemplateBse, nil, []string{}, `An alias for fmt.Sprintf`},
	"println":  {nil, goTemplateBse, nil, []string{}, `An alias for fmt.Sprintln`},
	"urlquery": {nil, goTemplateBse, nil, []string{}, `Returns the escaped value of the textual representation of its arguments in a form suitable for embedding in a URL query. This function is unavailable in html/template, with a few exceptions.`},

	"eq": {nil, goTemplateBse, nil, []string{}, `Returns the boolean truth of arg1 == arg2`},
	"ge": {nil, goTemplateBse, nil, []string{}, `Returns the boolean truth of arg1 >= arg2`},
	"gt": {nil, goTemplateBse, nil, []string{}, `Returns the boolean truth of arg1 > arg2`},
	"le": {nil, goTemplateBse, nil, []string{}, `Returns the boolean truth of arg1 <= arg2`},
	"lt": {nil, goTemplateBse, nil, []string{}, `Returns the boolean truth of arg1 < arg2`},
	"ne": {nil, goTemplateBse, nil, []string{}, `Returns the boolean truth of arg1 != arg2`},
}

// Add additional functions to the go template context
func (t *Template) addFuncs() {
	t.AddFunctions(baseGoTemplateFuncs)

	if t.options[Sprig] {
		t.addSprigFuncs()
	}

	if t.options[Math] {
		t.addMathFuncs()
	}

	if t.options[Data] {
		t.addDataFuncs()
	}

	if t.options[Logging] {
		t.addLoggingFuncs()
	}

	if t.options[Runtime] {
		t.addRuntimeFuncs()
	}

	if t.options[Utils] {
		t.addUtilsFuncs()
	}
}

// Apply all regular expressions replacements to the supplied string
func (t Template) substitute(content string) string {
	return utils.Substitute(content, t.substitutes...)
}

// List the available template names
func (t Template) getTemplateNames() []string {
	templates := t.Templates()
	result := make([]string, len(templates))
	for i := range templates {
		result[i] = templates[i].Name()
	}
	sort.Strings(result)
	return result
}