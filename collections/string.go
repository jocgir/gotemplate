package collections

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// AsString transform any object into a String
func AsString(i interface{}) String {
	switch s := i.(type) {
	case []byte:
		return String(s)
	case string:
		return String(s)
	case String:
		return s
	default:
		return String(fmt.Sprint(i))
	}
}

// AsStdString transform any object into a string
func AsStdString(i interface{}) string { return string(AsString(i)) }

// String is an enhanced class implementation of the standard go string library.
// This is convenient when manipulating go template string to have it considered as an object.
type String string

// IString is just a type that is used to supply argument that will be converted to String
type IString interface{}

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func (s String) Compare(b IString) int { return strings.Compare(string(s), AsStdString(b)) }

// Contains reports whether substr is within s.
func (s String) Contains(substr IString) bool { return strings.Contains(string(s), AsStdString(substr)) }

// ContainsAny reports whether any Unicode code points in chars are within s.
func (s String) ContainsAny(chars IString) bool {
	return strings.ContainsAny(string(s), AsStdString(chars))
}

// ContainsRune reports whether the Unicode code point r is within s.
func (s String) ContainsRune(r rune) bool { return strings.ContainsRune(string(s), r) }

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func (s String) Count(substr IString) int { return strings.Count(string(s), AsStdString(substr)) }

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding.
func (s String) EqualFold(t IString) bool { return strings.EqualFold(string(s), AsStdString(t)) }

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning an array of substrings of s or an
// empty list if s contains only white space.
func (s String) Fields() StringArray { return AsList(strings.Fields(string(s))).Strings() }

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
// FieldsFunc makes no guarantees about the order in which it calls f(c).
// If f does not return consistent results for a given c, FieldsFunc may crash.
func (s String) FieldsFunc(f func(rune) bool) StringArray {
	return AsList(strings.FieldsFunc(string(s), f)).Strings()
}

// HasPrefix tests whether the string s begins with prefix.
func (s String) HasPrefix(prefix IString) bool {
	return strings.HasPrefix(string(s), AsStdString(prefix))
}

// HasSuffix tests whether the string s ends with suffix.
func (s String) HasSuffix(suffix IString) bool {
	return strings.HasSuffix(string(s), AsStdString(suffix))
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func (s String) Index(substr IString) int { return strings.Index(string(s), AsStdString(substr)) }

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func (s String) IndexAny(chars IString) int { return strings.IndexAny(string(s), AsStdString(chars)) }

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func (s String) IndexByte(c byte) int { return strings.IndexByte(string(s), c) }

// IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do.
func (s String) IndexFunc(f func(rune) bool) int { return strings.IndexFunc(string(s), f) }

// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
func (s String) IndexRune(r rune) int { return strings.IndexRune(string(s), r) }

// Join concatenates the elements of array to create a single string. The string
// object is placed between elements in the resulting string.
func (s String) Join(array ...IString) String { return AsList(array).Join(s) }

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func (s String) LastIndex(substr IString) int {
	return strings.LastIndex(string(s), AsStdString(substr))
}

// LastIndexAny returns the index of the last instance of any Unicode code point from chars in s, or -1
// if no Unicode code point from chars is present in s.
func (s String) LastIndexAny(chars IString) int {
	return strings.LastIndexAny(string(s), AsStdString(chars))
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func (s String) LastIndexByte(c byte) int { return strings.LastIndexByte(string(s), c) }

// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func (s String) LastIndexFunc(f func(rune) bool) int { return strings.LastIndexFunc(string(s), f) }

// Lines splits up s into a StringArray using the newline as character separator
func (s String) Lines() StringArray { return s.Split("\n") }

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func (s String) Map(mapping func(rune) rune) String { return String(strings.Map(mapping, string(s))) }

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if the result of (len(s) * count) overflows.
func (s String) Repeat(count int) String { return String(strings.Repeat(string(s), count)) }

// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
func (s String) Split(sep IString) StringArray {
	return AsList(strings.Split(string(s), AsStdString(sep))).Strings()
}

// Split2 returns left and right part of a split.
func (s String) Split2(sep IString) (left, right String) {
	split := s.SplitN(sep, 2).Strings()
	left = split[0]
	if len(split) > 1 {
		right = split[1]
	}
	return
}

// SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter returns a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to SplitAfterN with a count of -1.
func (s String) SplitAfter(sep IString) StringArray {
	return AsList(strings.SplitAfter(string(s), AsStdString(sep))).Strings()
}

// SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for SplitAfter.
func (s String) SplitAfterN(sep IString, n int) StringArray {
	return AsList(strings.SplitAfterN(string(s), AsStdString(sep), n)).Strings()
}

// SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.
//
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for Split.
func (s String) SplitN(sep IString, n int) StringArray {
	return AsList(strings.SplitN(string(s), AsStdString(sep), n)).Strings()
}

// Title returns a copy of the string s with all Unicode letters that begin words mapped to their title case.
//
// BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
func (s String) Title() String { return String(strings.Title(string(s))) }

// ToLower returns a copy of the string s with all Unicode letters mapped to their lower case.
func (s String) ToLower() String { return String(strings.ToLower(string(s))) }

// ToTitle returns a copy of the string s with all Unicode letters mapped to their title case.
func (s String) ToTitle() String { return String(strings.ToTitle(string(s))) }

// ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.
func (s String) ToUpper() String { return String(strings.ToUpper(string(s))) }

// Trim returns a slice of the string s with all leading and // trailing Unicode code points contained in cutset removed.
func (s String) Trim(cutset IString) String {
	return String(strings.Trim(string(s), AsStdString(cutset)))
}

// TrimFunc returns a slice of the string s with all leading and trailing Unicode code points c satisfying f(c) removed.
func (s String) TrimFunc(f func(rune) bool) String { return String(strings.TrimFunc(string(s), f)) }

// TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed.
func (s String) TrimLeft(cutset IString) String {
	return String(strings.TrimLeft(string(s), AsStdString(cutset)))
}

// TrimLeftFunc returns a slice of the string s with all leading Unicode code points c satisfying f(c) removed.
func (s String) TrimLeftFunc(f func(rune) bool) String {
	return String(strings.TrimLeftFunc(string(s), f))
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func (s String) TrimPrefix(prefix IString) String {
	return String(strings.TrimPrefix(string(s), AsStdString(prefix)))
}

// TrimRight returns a slice of the string s, with all trailing Unicode code points contained in cutset removed.
func (s String) TrimRight(cutset IString) String {
	return String(strings.TrimRight(string(s), AsStdString(cutset)))
}

// TrimRightFunc returns a slice of the string s with all trailing Unicode code points c satisfying f(c) removed.
func (s String) TrimRightFunc(f func(rune) bool) String {
	return String(strings.TrimRightFunc(string(s), f))
}

// TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
func (s String) TrimSpace() String { return String(strings.TrimSpace(string(s))) }

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func (s String) TrimSuffix(suffix IString) String {
	return String(strings.TrimSuffix(string(s), AsStdString(suffix)))
}

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case, giving priority to the special casing rules.
func (s String) ToLowerSpecial(c unicode.SpecialCase) String {
	return String(strings.ToLowerSpecial(c, string(s)))
}

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// title case, giving priority to the special casing rules.
func (s String) ToTitleSpecial(c unicode.SpecialCase) String {
	return String(strings.ToTitleSpecial(c, string(s)))
}

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case, giving priority to the special casing rules.
func (s String) ToUpperSpecial(c unicode.SpecialCase) String {
	return String(strings.ToUpperSpecial(c, string(s)))
}

// -------------------------------------------------------------------------------------------------------------------
// The following functions are extension or variation of the standard go string library

// String simply convert a String object into a regular string.
func (s String) String() string { return string(s) }

// Str is an abbreviation of String.
func (s String) Str() string { return string(s) }

// Len returns the length of the string.
func (s String) Len() int { return len(s) }

// Format returns the string on which the format has been applied.
func (s String) Format(format interface{}, args ...interface{}) String {
	return FormatMessage(AsList(format).Append(s).Append(args...).AsArray()...)
}

// Quote returns the string between quotes.
func (s String) Quote() String { return String(fmt.Sprintf("%q", s)) }

// Escape returns the representation of the string with escape characters.
func (s String) Escape() String {
	q := s.Quote()
	return String(q[1 : len(q)-1])
}

// FieldsID splits the string s at character that is not a valid identifier character (letter, digit or underscore).
func (s String) FieldsID() StringArray {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '_'
	}
	return s.FieldsFunc(f)
}

// Center returns the string centered within the specified width.
func (s String) Center(width int) String {
	l := utf8.RuneCountInString(string(s))
	if l > width {
		return s
	}
	left := (width - l) / 2
	right := width - left - l
	return s.Format("%[3]*[2]s%[1]s%[4]*[2]s", "", left, right)
}

// Wrap returns the string wrapped with newline when exceeding the specified width.
func (s String) Wrap(width int) String {
	lines := s.Lines().Strings()
	result := CreateList(0, len(lines))

	for i := range lines {
		words := lines[i].TrimRight("\t ").Fields().StdStrings()
		start, length := 0, 0
		for j := range words {
			if length > 0 && length+len(words[j]) > width {
				result = result.Append(strings.Join(words[start:j], " "))
				start, length = j, 0
			}
			length += len(words[j])
			if j-start > 0 {
				length++
			}
		}
		result = result.Append(strings.Join(words[start:len(words)], " "))
	}

	return result.Join("\n")
}

// Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to
// k+1 replacements for a k-rune string.
func (s String) Replace(old, new IString) String {
	return String(strings.Replace(string(s), AsStdString(old), AsStdString(new), -1))
}

// ReplaceN returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to
// If n < 0, there is no limit on the number of replacements.
func (s String) ReplaceN(old, new IString, n int) String {
	return String(strings.Replace(string(s), AsStdString(old), AsStdString(new), n))
}

// Indent returns the indented version of the supplied string (indent represents the string used to indent the lines).
func (s String) Indent(indent IString) String {
	split, spacer := s.Lines().Strings(), AsString(indent)

	for i := range split {
		split[i] = spacer + split[i]
	}
	return split.Join("\n")
}

// IndentN returns the indented version of the supplied string (indent represents the number of spaces used to indent the lines).
func (s String) IndentN(indent int) String { return s.Indent(String(" ").Repeat(indent)) }

// UnIndent returns the string unindented.
func (s String) UnIndent() String {
	lines := s.Lines().Strings()
	if len(lines) <= 1 {
		return s
	}

	var spaces *String
	for i, line := range lines {
		if spaces == nil {
			if line.TrimSpace() == "" {
				// We do not consider empty lines
				continue
			}
			trimmed := line.TrimLeftFunc(unicode.IsSpace)
			trimmed = lines[i][:len(lines[i])-len(trimmed)]
			spaces = &trimmed
		}
		if !line.HasPrefix(*spaces) && line.TrimSpace() != "" {
			return s
		}
		lines[i] = line.TrimPrefix(*spaces)
	}

	return AsList(lines).Join("\n")
}

// GetWordAtPosition returns the selected word and the start position from the specified position.
func (s String) GetWordAtPosition(pos int, accept ...string) (String, int) {
	if pos < 0 || pos >= len(s) {
		return "", -1
	}

	acceptChars := strings.Join(accept, "")
	isBreak := func(c rune) bool {
		return unicode.IsSpace(c) || unicode.IsPunct(c) && !strings.ContainsRune(acceptChars, c)
	}
	begin, end := pos, pos
	for begin >= 0 && !isBreak(rune(s[begin])) {
		begin--
	}
	for end < len(s) && !isBreak(rune(s[end])) {
		end++
	}
	if begin != end {
		begin++
	}
	return s[begin:end], begin
}

// SelectWord returns the selected word from the specified position.
func (s String) SelectWord(pos int, accept ...string) String {
	result, _ := s.GetWordAtPosition(pos, accept...)
	return result
}

// IndexAll returns all positions where substring is found within s.
func (s String) IndexAll(substr string) (result []int) {
	if substr == "" || s == "" {
		return nil
	}
	start, lenSubstr := 0, len(substr)
	for pos := s.Index(substr); pos >= 0; pos = s[start:].Index(substr) {
		result = append(result, start+pos)
		start += pos + lenSubstr
	}
	return
}

// GetContextAtPosition tries to extend the context from the specified position within specified boundaries.
func (s String) GetContextAtPosition(pos int, left, right string) (a String, b int) {
	if pos < 0 || pos >= len(s) {
		// Trying to select context out of bound
		return "", -1
	}

	findLeft := func(s String, pos int) int {
		if left == "" {
			return pos
		}
		return s[:pos].LastIndex(left)
	}
	begin, end, lenLeft, lenRight := findLeft(s, pos), pos+1, len(left), len(right)

	if begin >= 0 && lenRight > 0 {
		if end = s[pos:].Index(right); end >= 0 {
			end += pos + lenRight
			back := findLeft(s[:end], end-lenRight)
			if left != "" && back != begin {
				// There is at least one enclosed start, we must find the corresponding end
				pos = begin + lenLeft
				lefts := s[pos:].IndexAll(left)
				rights := s[pos:].IndexAll(right)
				for i := range lefts {
					if i == len(rights) {
						return "", -1
					}
					if lefts[i] > rights[i] {
						return s[begin : rights[i]+pos+lenRight], begin
					}
				}
				if len(rights) > len(lefts) {
					return s[begin : rights[len(lefts)]+pos+lenRight], begin
				}
				end = -1
			}
		}
	}
	if begin < 0 || end < 0 {
		return "", -1
	}
	return s[begin:end], begin
}

// SelectContext returns the selected word from the specified position.
func (s String) SelectContext(pos int, left, right string) String {
	result, _ := s.GetContextAtPosition(pos, left, right)
	return result
}

// Protect returns a string with all included strings replaced by a token and an array of all replaced strings.
// The function is able to detect strings enclosed between quotes "" or backtick `` and it detects escaped characters.
func (s String) Protect() (result String, array StringArray) {
	defer func() { result += s }()

	escaped := func(from int) bool {
		// Determine if the previous characters are escaping the current value
		count := 0
		for ; s[from-count] == '\\'; count++ {
		}
		return count%2 == 1
	}

	for end := 0; end >= 0; {
		pos := s.IndexAny("`\"")
		if pos < 0 {
			break
		}

		for end = s[pos+1:].IndexRune(rune(s[pos])); end >= 0; {
			end += pos + 1
			if s[end] == '"' && escaped(end-1) {
				// The quote has been escaped, so we find the next one
				newEnd := s[end+1:].IndexRune('"')
				if newEnd < 0 {
					return
				}
				end += newEnd - pos
			} else {
				array = append(array, s[pos:end+1])
				result += s[:pos] + String(fmt.Sprintf(replacementFormat, array.Count()-1))
				s = s[end+1:]
				break
			}
		}
	}
	return
}

// RestoreProtected restores a string transformed by ProtectString to its original value.
func (s String) RestoreProtected(array StringArray) String {
	return String(replacementRegex.ReplaceAllStringFunc(s.Str(), func(match string) string {
		index := Must(strconv.Atoi(replacementRegex.FindStringSubmatch(match)[1])).(int)
		return array[index].Str()
	}))
}

const replacementFormat = `"♠%d"`

var replacementRegex = regexp.MustCompile(`"♠(\d+)"`)

// AddLineNumber adds line number to a string
func (s String) AddLineNumber(space int) String {
	lines := s.Lines().AsArray()
	if space <= 0 {
		space = int(math.Log10(float64(len(lines)))) + 1
	}

	for i := range lines {
		lines[i] = String(fmt.Sprintf("%*d %s", space, i+1, lines[i]))
	}
	return AsList(lines).Join("\n")
}

// ParseBool returns true if variable exist and is not clearly a false value
//    i.e. empty, 0, Off, No, n, false, f
func (s String) ParseBool() bool {
	// We first try with the strconv library
	if result, err := strconv.ParseBool(s.Str()); err == nil {
		return result
	}
	switch s.ToUpper() {
	case "", "N", "NO", "OFF":
		return false
	default:
		// Any other value is considered as true
		return true
	}
}
