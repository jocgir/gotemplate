package template

import (
	"fmt"
	"strings"

	"github.com/coveooss/multilogger/reutils"
)

var alreadyIssued = make(map[string]int)

func assignExpression(repl replacement, match string) string {
	return assignExpressionInternal(repl, match, false)
}

func assignExpressionAcceptError(repl replacement, match string) string {
	return assignExpressionInternal(repl, match, true)
}

func assignExpressionInternal(repl replacement, match string, acceptError bool) string {
	matches, _ := reutils.MultiMatch(match, repl.re)
	tp := matches["type"]
	id := matches["id"]
	expr := matches["expr"]
	assign := matches["assign"]
	if tp == "" || id == "" || expr == "" || assign == "" {
		InternalLog.Errorf("Invalid assign regex %s: %s for '%s', must contains type, id and expr", repl.name, repl.expr, match)
		return match
	}

	global := false
	internal := true
	switch tp {
	case "@$.":
		tp = "@"
		fallthrough
	case "@", "@.":
		global = true
	case "@{", "@$":
		internal = strings.Contains(id, ".")
	}

	switch assign {
	case ":=", "=":
		break
	default:
		// This is an assignment operator (i.e. +=, /=, <<=, etc.)
		operator := assign[:len(assign)-1]
		assign = "="
		if tp == "@{" {
			expr = fmt.Sprintf("$%[1]s %[2]s (%[3]s)", id, operator, expr)
		} else {
			value := map[string]string{"@": "$.", "@.": ".", "@$": "$"}[tp] + id
			expr = fmt.Sprintf("%[1]s %[2]s (%[3]s)", value, operator, expr)
		}
	}

	var err error
	if expr, err = expressionParserInternal(exprRepl, expr, true, internal); err != nil && !acceptError {
		return match
	}

	if !global && !internal {
		return fmt.Sprintf("%s- $%s %s %s %s", repl.delimiters[0], id, assign, expr, repl.delimiters[1])
	}

	parts := strings.Split(id, ".")
	object := strings.Join(parts[:len(parts)-1], ".")
	id = parts[len(parts)-1]

	if strings.HasSuffix(tp, ".") {
		// The object is assigned with a syntax such as:
		//   @.a := ...
		//   @.a.b := ...
		object = "." + object
	} else if global {
		// The object is assigned with a syntax such as:
		//   @a := ...
		//   @a.b := ...
		object = iif(object == "", "$", "$."+object).(string)
	} else {
		// The object is assigned with a syntax such as:
		//   @$a := ...
		//   @{a} := ...
		//   @{a := ...}
		//   @$a.b := ...
		//   @{a.b} := ...
		//   @{a.b := ...}
		object = "$" + object
	}

	if assign == ":=" || StrictAssignationMode == AssignationValidationDisabled {
		// We do not check if the variable already exist (or not)
		return fmt.Sprintf(`%[1]s- set %[3]s "%[4]s" %[5]s %[2]s`, repl.delimiters[0], repl.delimiters[1], object, id, expr)
	}
	objectID := fmt.Sprintf("%s%s", iif(strings.HasSuffix(object, "."), object, object+"."), id)
	validateCode := iif(StrictAssignationMode == AssignationValidationWarning, "assertWarning", "assert").(string)
	validateCode += fmt.Sprintf(` (not (isNil %[1]s)) "%[1]s does not exist, use := to declare new variable"`, objectID)
	return fmt.Sprintf(`%[1]s- %[3]s %[2]s%[1]s- set %[4]s "%[5]s" %[6]s %[2]s`, repl.delimiters[0], repl.delimiters[1], validateCode, object, id, expr)
}

// AssignationValidationType is the enum type to define valid global variables validation mode.
type AssignationValidationType uint8

// Valid values for AssignationValidationType
const (
	AssignationValidationDisabled AssignationValidationType = iota
	AssignationValidationWarning
	AssignationValidationStrict
)

// StrictAssignationMode defines the global assignation validation mode.
var StrictAssignationMode = AssignationValidationWarning
