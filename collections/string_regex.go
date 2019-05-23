package collections

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/coveo/gotemplate/v3/errors"
)

// MultiMatch returns a map of matching elements from a list of regular expressions (returning the first matching element)
func (s String) MultiMatch(expressions ...*regexp.Regexp) (map[string]string, int) {
	for i, re := range expressions {
		if matches := re.FindStringSubmatch(s.Str()); len(matches) != 0 {
			results := make(map[string]string)
			for i, key := range re.SubexpNames() {
				if key != "" {
					results[key] = matches[i]
				}
			}
			return results, i
		}
	}
	return nil, -1
}

// GetRegexGroup caches compiled regex to avoid multiple interpretation of the same regex
func GetRegexGroup(key string, definitions []string) (result []*regexp.Regexp, err error) {
	if result, ok := cachedRegex[key]; ok {
		return result, nil
	}

	result = make([]*regexp.Regexp, len(definitions))
	for i := range definitions {
		regex, err := regexp.Compile(definitions[i])
		if err != nil {
			return nil, err
		}
		result[i] = regex
	}

	cachedRegex[key] = result
	return
}

var cachedRegex = map[string][]*regexp.Regexp{}

// BuildBatchRegex builds a regular expression to replace all specified expression by their replacement
func BuildBatchRegex(pairs ...ReplacementPair) (re *BatchRegex, err error) {
	if len(pairs) == 0 {
		return nil, fmt.Errorf("There must be at least 1 pair")
	}
	re = &BatchRegex{literals: make(map[string]string)}
	expressions := make([]string, len(pairs))
	var errs errors.Array
	for i := range pairs {
		expressions[i] = pairs[i].expression
		if exp := regexp.QuoteMeta(pairs[i].expression); exp == pairs[i].expression {
			re.literals[exp] = pairs[i].replacement
		} else {
			if exp, err := regexp.Compile(pairs[i].expression); err != nil {
				errs = append(errs, err)
			} else {
				re.expressions = append(re.expressions, exp)
				re.replacements = append(re.replacements, pairs[i].replacement)
			}
		}
	}
	if errs.AsError() != nil {
		return nil, errs
	}

	re.Regexp, err = regexp.Compile(fmt.Sprintf("(%s)", strings.Join(expressions, "|")))
	if err != nil {
		re = nil
	}
	return re, err
}

// BatchReplace returns the modified string
func (s String) BatchReplace(re *BatchRegex) String {
	s, _ = s.BatchReplaceReversible(re)
	return s
}

// BatchReplaceReversible returns the modified string and all the replacements that occurred
func (s String) BatchReplaceReversible(re *BatchRegex) (String, []Replacement) {
	result, repl := re.Replace(s.Str())
	return String(result), repl
}

// BatchReplaceRevert revert the previously applied changes to restore the string as its original value
func (s String) BatchReplaceRevert(replacements []Replacement) String {
	for _, r := range replacements {
		if r.Position > len(s) {
			break
		}
		s = s.ReplaceN(r.New, r.Original, 1)
	}
	return s
}

// NewReplacementPair returns an object used to feed BuildBatchRegex
func NewReplacementPair(expression, replacement string) ReplacementPair {
	return ReplacementPair{expression, replacement}
}

// Replacement represents a replacement text that occurred during Replace on BatchRegex
type Replacement struct {
	Original string
	New      string
	Position int
}

func (r Replacement) String() string {
	return fmt.Sprintf("('%s' at %d => '%s')", r.Original, r.Position, r.New)
}

// ReplacementPair represent a replacement combination
type ReplacementPair struct {
	expression  string
	replacement string
}

// BatchRegex is a regular expression that is able to replace many expression
// and returns the replaced text with their position
type BatchRegex struct {
	*regexp.Regexp
	literals     map[string]string
	expressions  []*regexp.Regexp
	replacements []string
}

// Replace returns the modified string and all the replacements that occurred
func (re *BatchRegex) Replace(s string) (result string, replacements []Replacement) {
	matches := re.FindAllStringIndex(s, -1)
	if matches == nil {
		return s, nil
	}
	i := 0
	result = re.ReplaceAllStringFunc(s, func(s string) (replaced string) {
		replaced, literal := re.literals[s]
		if !literal {
			_, index := String(s).MultiMatch(re.expressions...)
			if index == -1 {
				errors.Raise("Invalid match on '%s'", s)
			}
			replaced = re.expressions[index].ReplaceAllString(s, re.replacements[index])
		}
		replacements = append(replacements, Replacement{Original: s, New: replaced, Position: matches[i][0]})
		i++
		return
	})
	return
}
