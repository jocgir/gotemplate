package template

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionTemplating(t *testing.T) {
	t.Parallel()
	template := New("Test")
	template.completeExamples()

	for _, functionName := range template.getFunctions() {
		funcInfo := template.getFunction(functionName)
		for i, test := range funcInfo.examples {
			example := test
			t.Run(fmt.Sprintf("%s_#%d", funcInfo.name, i), func(t *testing.T) {
				if example.Razor != "" {
					appliedRazor, changed := template.applyRazor([]byte(example.Razor))
					assert.Equal(t, example.Template, string(appliedRazor), "Razor wasn't resolved correctly")
					assert.True(t, changed)
				}
				if example.Template != "" {
					appliedTemplate, err := template.ProcessContent(example.Template, ".")
					assert.Equal(t, example.Result, appliedTemplate, "Template wasn't resolved correctly")
					assert.NoError(t, err)
				}
			})
		}
	}
}
