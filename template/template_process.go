package template

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/coveooss/gotemplate/v3/utils"
	"golang.org/x/crypto/ssh/terminal"
)

var templateExt = []string{".gt", ".template"}

func p(name, expr string) string { return fmt.Sprintf("(?P<%s>%s)", name, expr) }

const (
	noValue      = "<no value>"
	noValueRepl  = "!NO_VALUE!"
	nilValue     = "<nil>"
	nilValueRepl = "!NIL_VALUE!"
	undefError   = `"` + noValue + `"`
	noValueError = "contains undefined value(s)"
	runError     = `"<RUN_ERROR>"`
	tagLine      = "line"
	tagCol       = "column"
	tagCode      = "code"
	tagMsg       = "message"
	tagLocation  = "location"
	tagFile      = "file"
	tagKey       = "key"
	tagErr       = "error"
)

// ProcessContent loads and runs the file template.
func (t *Template) ProcessContent(content, source string) (result string, err error) {
	result, _, err = t.processContentInternal(content, source, nil, 0, true, nil)
	return
}

// ProcessTemplate loads and runs the template if it is a file, otherwise, it simply process the content.
func (t *Template) ProcessTemplate(template, sourceFolder, targetFolder string) (resultFile string, err error) {
	return t.processTemplate(template, sourceFolder, targetFolder, nil)
}

// ProcessTemplates loads and runs the file template or execute the content if it is not a file.
func (t *Template) ProcessTemplates(sourceFolder, targetFolder string, templates ...string) (resultFiles []string, err error) {
	return t.ProcessTemplatesWithHandler(sourceFolder, targetFolder, nil, templates...)
}

func (t *Template) printResult(source, target, result string) (err error) {
	if utils.IsTerraformFile(target) {
		base := filepath.Base(target)
		tempFolder := must(ioutil.TempDir(t.tempFolder, base)).(string)
		tempFile := filepath.Join(tempFolder, base)
		err = ioutil.WriteFile(tempFile, []byte(result), 0644)
		if err != nil {
			return
		}
		err = utils.TerraformFormat(tempFile)
		bytes := must(ioutil.ReadFile(tempFile)).([]byte)
		result = string(bytes)
	}

	if !t.isTemplate(source) && !t.options[Overwrite] {
		source += ".original"
	}

	source = utils.Relative(t.folder, source)
	if relTarget := utils.Relative(t.folder, target); !strings.HasPrefix(relTarget, "../../../") {
		target = relTarget
	}
	if source != target {
		InternalLog.Infof("%s => %s", source, target)
	} else {
		InternalLog.Info(target)
	}
	Print(result)
	if result != "" && terminal.IsTerminal(int(os.Stdout.Fd())) {
		Println()
	}

	return
}

func (t *Template) execute(actual *template.Template, wr io.Writer, data interface{}) error {
	if dict, ok := data.(iDictionary); ok {
		contextStack = append(contextStack, dict)
		defer func() { contextStack = contextStack[:len(contextStack)-1] }()
	}
	return actual.Execute(wr, data)
}

func getContext(n int) iDictionary {
	return contextStack[len(contextStack)-1-n]
}

var contextStack []iDictionary
