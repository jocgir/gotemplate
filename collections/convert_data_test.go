package collections

import (
	"testing"
)

func Test_quote(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg  string
		want String
	}{
		{"Simple value", "Foo", "Foo"},
		{"Simple value", "Foo Bar", `"Foo Bar"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quote(tt.arg); got != tt.want {
				t.Errorf("quote() = %v, want %v", got, tt.want)
			}
		})
	}
}
