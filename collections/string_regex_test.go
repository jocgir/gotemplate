package collections

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildBatchReplacer(t *testing.T) {
	type build struct {
		pairs   []ReplacementPair
		want    *BatchRegex
		wantErr string
	}
	type test struct {
		reversible bool
		text       String
		want       String
		wantList   []Replacement
	}
	tests := []struct {
		name string
		b    build
		t    test
	}{
		{"Empty", build{nil, nil, "There must be at least 1 pair"}, test{}},
		{"With elements",
			build{[]ReplacementPair{{"a", "b"}, {"[bcd]", "."}}, &BatchRegex{regexp.MustCompile("(a|[bcd])"), map[string]string{"a": "b"}, []*regexp.Regexp{regexp.MustCompile("[bcd]")}, []string{"."}}, ""},
			test{true, "abcdef", "b...ef", []Replacement{{"a", "b", 0}, {"b", ".", 1}, {"c", ".", 2}, {"d", ".", 3}}}},
		{"With range",
			build{[]ReplacementPair{{"a", "*"}, {"[b-e]", "-"}}, &BatchRegex{regexp.MustCompile("(a|[b-e])"), map[string]string{"a": "*"}, []*regexp.Regexp{regexp.MustCompile("[b-e]")}, []string{"-"}}, ""},
			test{true, "abcdef", "*----f", []Replacement{{"a", "*", 0}, {"b", "-", 1}, {"c", "-", 2}, {"d", "-", 3}, {"e", "-", 4}}}},
		{"No match",
			build{[]ReplacementPair{{"a", "*"}, {"[b-e]", "-"}}, &BatchRegex{regexp.MustCompile("(a|[b-e])"), map[string]string{"a": "*"}, []*regexp.Regexp{regexp.MustCompile("[b-e]")}, []string{"-"}}, ""},
			test{false, "0123456789", "0123456789", nil}},
		{"With transformation",
			build{[]ReplacementPair{{"a", "*"}, {"[b-e]", "-$0"}}, &BatchRegex{regexp.MustCompile("(a|[b-e])"), map[string]string{"a": "*"}, []*regexp.Regexp{regexp.MustCompile("[b-e]")}, []string{"-$0"}}, ""},
			test{true, "abcdef", "*-b-c-d-ef", []Replacement{{"a", "*", 0}, {"b", "-b", 1}, {"c", "-c", 2}, {"d", "-d", 3}, {"e", "-e", 4}}}},
		{"With group transformation",
			build{[]ReplacementPair{{"a", "*"}, {"(?P<letter>[b-d])", ".${letter}"}}, &BatchRegex{regexp.MustCompile("(a|(?P<letter>[b-d]))"), map[string]string{"a": "*"}, []*regexp.Regexp{regexp.MustCompile("(?P<letter>[b-d])")}, []string{".${letter}"}}, ""},
			test{true, "abcdef", "*.b.c.def", []Replacement{{"a", "*", 0}, {"b", ".b", 1}, {"c", ".c", 2}, {"d", ".d", 3}}}},
		{"With group error",
			build{[]ReplacementPair{{"a", "*"}, {"[b-d)", ".${letter}"}}, nil, "error parsing regexp: missing closing ]: `[b-d)`"},
			test{false, "abcdef", "", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildBatchRegex(tt.b.pairs...)
			var result String
			var list []Replacement
			if tt.b.wantErr == "" {
				assert.NoError(t, err)
				if err == nil {
					if tt.t.reversible {
						result, list = tt.t.text.BatchReplaceReversible(got)
						assert.Equal(t, tt.t.text, result.BatchReplaceRevert(list))
					} else {
						result = tt.t.text.BatchReplace(got)
					}
				}
			} else {
				assert.EqualError(t, err, tt.b.wantErr)
			}
			assert.Equal(t, tt.b.want, got)
			assert.Equal(t, tt.t.want, result)
			assert.Equal(t, tt.t.wantList, list)
		})
	}
}
