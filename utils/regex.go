package utils

import (
	"regexp"

	"github.com/coveo/gotemplate/v3/collections"
)

// MultiMatch returns a map of matching elements from a list of regular expressions (returning the first matching element)
func MultiMatch(s string, expressions ...*regexp.Regexp) (map[string]string, int) {
	return String(s).MultiMatch(expressions...)
}

// Imports from collections
var (
	GetRegexGroup   = collections.GetRegexGroup
	BuildBatchRegex = collections.BuildBatchRegex
)

// BatchReplace returns the modified string
func BatchReplace(s string, re *collections.BatchRegex) string {
	return String(s).BatchReplace(re).Str()
}

// BatchReplaceReversible returns the modified string and all the replacements that occurred
func BatchReplaceReversible(s string, re *collections.BatchRegex) (string, []collections.Replacement) {
	result, list := String(s).BatchReplaceReversible(re)
	return result.Str(), list
}

// BatchReplaceRevert revert the previously applied changes to restore the string as its original value
func BatchReplaceRevert(s string, replacements []collections.Replacement) string {
	return String(s).BatchReplaceRevert(replacements).Str()
}
