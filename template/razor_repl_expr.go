package template

import (
	"fmt"
	"go/parser"
	"go/scanner"
	"regexp"
	"strings"

	"github.com/coveo/gotemplate/v3/collections"
	"github.com/coveo/gotemplate/v3/utils"
	"github.com/fatih/color"
	"github.com/op/go-logging"
)

const (
	protectString = "_=LONG_STRING="
	literalAt     = "_=!AT!=_"
	literalStart  = `{{ "{{" }}`
	stringRep     = "__StRiNg__"
	rangeExpr     = "__RaNgE__"
	defaultExpr   = "__DeFaUlT__"
	funcExpr      = "__FuNc__"
	funcCall      = "__FuNcAlL__"
	typeExpr      = "__TyPe__"
	mapExpr       = "__MaP__"
	dotRep        = "__DoT__"
	ellipsisRep   = "__ElLiPsIs__"
	globalRep     = "__GlObAl__"
)

var dotPrefix = regexp.MustCompile(`(?P<prefix>^|[^\w\)\]])\.(?P<value>\w[\w\.]*)?`)
var idRegex = regexp.MustCompile(`^[\p{L}\d_]+$`)

func expressionParser(repl replacement, match string) string {
	expr, _ := expressionParserInternal(repl, match, false, false)
	return expr
}

func expressionParserSkipError(repl replacement, match string) string {
	expr, _ := expressionParserInternal(repl, match, true, false)
	return expr
}

func expressionParserInternal(repl replacement, match string, skipError, internal bool) (result string, err error) {
	matches, _ := utils.MultiMatch(match, repl.re)
	var expr, expression string
	var replacements []collections.Replacement
	if expression = matches["expr"]; expression != "" {
		if getLogLevelInternal() >= logging.DEBUG {
			defer func() {
				if !debugMode && result != match {
					log.Debug("Resulting expression =", result)
				}
			}()
		}

		// We first protect strings declared in the expression
		protected, includedStrings := String(expression).Protect()
		protected, replacements = protected.BatchReplaceReversible(expressionTransformer)

		// We transform the expression into a valid go statement
		// for k, v := range map[string]string{"$": stringRep, "range": rangeExpr, "default": defaultExpr, "func": funcExpr, "...": ellipsisRep, "type": typeExpr, "map": mapExpr} {
		// 	protected = protected.Replace(k, v)
		// }
		//protected = String(dotPrefix.ReplaceAllString(protected.Str(), fmt.Sprintf("${prefix}%s${value}", dotRep)))
		//protected = protected.Replace(ellipsisRep, "...")

		// for k, v := range map[string]string{"<>": "!=", "÷": "/", "≠": "!=", "≦": "<=", "≧": ">=", "«": "<<", "»": ">>"} {
		// 	protected = protected.Replace(k, v)
		// }

		// for key, val := range ops {
		// 	protected = protected.Replace(" "+val+" ", key)
		// }
		// We add support to partial slice
		//protected = String(indexExpression(protected.Str()))

		// We restore the strings into the expression
		expr = protected.RestoreProtected(includedStrings).Str()
	}

	if indexExpr := matches["index"]; indexExpr != "" {
		indexExpr = indexExpression(indexExpr)
		indexExpr = indexExpr[1 : len(indexExpr)-1]

		sep, slicer, limit2 := ",", "extract", false
		if strings.Contains(indexExpr, ":") {
			sep, slicer, limit2 = ":", "slice", true
		}
		values := strings.Split(indexExpr, sep)
		if !debugMode && limit2 && len(values) > 2 {
			log.Errorf("Only one : character is allowed in slice expression: %s", match)
		}
		for i := range values {
			if values[i], err = expressionParserInternal(exprRepl, values[i], true, true); err != nil {
				return match, err
			}
		}
		indexExpr = strings.Replace(strings.Join(values, " "), `$`, `$$`, -1)
		repl.replace = strings.Replace(repl.replace, "${index}", indexExpr, -1)
		repl.replace = strings.Replace(repl.replace, "${slicer}", slicer, -1)
	}

	if selectExpr := matches["selection"]; selectExpr != "" {
		if selectExpr, err = expressionParserInternal(exprRepl, selectExpr, true, true); err != nil {
			return match, err
		}
		repl.replace = strings.Replace(repl.replace, "${selection}", selectExpr, -1)
	}

	if expr != "" {
		node := nodeValue
		if internal {
			node = nodeValueInternal
		}
		tr, err := parser.ParseExpr(expr)
		if err != nil {
			pos := err.(scanner.ErrorList)[0].Pos.Offset
			expr = expr[:pos]
			if tr2, err2 := parser.ParseExpr(expr); err2 == nil {
				newRegex := exprFix.ReplaceAllString(repl.re.String(), strings.Replace(regexp.QuoteMeta(utils.BatchReplaceRevert(expr, replacements)), "$", "$$", -1))
				//remaining = expr[pos:]
				if must(regexp.MatchString(newRegex, match)).(bool) {
					err = nil
					tr = tr2
				}
			}
		}
		if err == nil {
			result, err := node(tr)
			if err == nil {
				result = String(result).BatchReplace(revertExpressions).Replace("$", "$$").Str()
				repl.replace = strings.Replace(repl.replace, "${expr}", result, -1)
				result = repl.re.ReplaceAllString(match, repl.replace)
				return result, nil
			}
		}
		if !debugMode && err != nil && getLogLevelInternal() >= 6 {
			log.Debug(color.CyanString(fmt.Sprintf("Invalid expression '%s' : %v", expression, err)))
		}
		if skipError {
			return match, err
		}
		repl.replace = strings.Replace(repl.replace, "${expr}", strings.Replace(expression, "$", "$$", -1), -1)
	}

	return repl.re.ReplaceAllString(match, repl.replace), nil
}

var exprFix = regexp.MustCompile(`\(\?P<expr>.+?\)`)

var exprRepl = replacement{
	name:    "Expression",
	re:      regexp.MustCompile(`^(?P<expr>.*)$`),
	replace: `${expr}`,
}

func indexExpression(expr string) string {
	expr = negativeSlice.ReplaceAllString(expr, "[${index}:0]")
	expr = strings.Replace(expr, "[]", "[0:-1]", -1)
	expr = strings.Replace(expr, "[:", "[0:", -1)
	expr = strings.Replace(expr, ":]", ":-1]", -1)
	return expr
}

var negativeSlice = regexp.MustCompile(`\[(?P<index>-\d+):]`)

// Build the batch regular expression to replace elements in expression
// and ensure that the resulting string is valid
var expressionTransformer = func() *collections.BatchRegex {
	rp := collections.NewReplacementPair
	replacementPairs := []collections.ReplacementPair{
		rp(`\$`, stringRep), // $ is not a valid character
		// Go reserved word
		rp("range", rangeExpr),
		rp("default", defaultExpr),
		rp("func", funcExpr),
		rp("type", typeExpr),
		rp("map", mapExpr),
		// A variable cannot starts with . in go
		rp(`(\B\.\b|^\.$)`, dotRep),
		// Symbol aliases
		rp("<>", "!="),
		rp("÷", "/"),
		rp("≠", "!="),
		rp("≦", "<="),
		rp("≧", ">="),
		rp("«", "<<"),
		rp("»", ">>"),
		// Partial index support
		rp(`\[(?P<index>-\d+):]`, "[${index}:0]"),
		rp(`\[\]`, "[0:-1]"),
		rp(`\[:`, "[0:"),
		rp(`:]`, ":-1]"),
	}
	for key, val := range ops {
		replacementPairs = append(replacementPairs, rp(" "+val+" ", key))
	}

	return must(utils.BuildBatchRegex(replacementPairs...)).(*collections.BatchRegex)
}()

var revertExpressions = func() *collections.BatchRegex {
	rp := collections.NewReplacementPair
	backPairs := []collections.ReplacementPair{
		rp(stringRep, "$"),
		rp(rangeExpr, "range"),
		rp(defaultExpr, "default"),
		rp(funcExpr, "func"),
		rp(typeExpr, "type"),
		rp(mapExpr, "map"),
		rp(dotRep, "."),
		rp(globalRep, "$."),
		rp(funcCall, ""),
		rp(`\$\.slice `, "slice $."),
	}
	return must(utils.BuildBatchRegex(backPairs...)).(*collections.BatchRegex)
}()
