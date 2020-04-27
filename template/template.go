package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/coveooss/gotemplate/v3/collections"
	"github.com/coveooss/gotemplate/v3/utils"
	"github.com/jocgir/template"
)

// IsRazor determines if the supplied code appears to have Razor code (using default delimiters).
func IsRazor(code string) bool { return defaultTemplate.IsRazor(code) }

// IsCode determines if the supplied code appears to have gotemplate code (using default delimiters).
func IsCode(code string) bool { return defaultTemplate.IsCode(code) }

// Must is a helper that wraps a call to a function returning (*Template, error)
// and panics if the error is non-nil. It is intended for use in variable
// initializations such as:
//	var t = template.Must(template.New("name").Parse("text"))
func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

// New allocates a new, undefined template with the given name.
func New(name string) *Template {
	// Set the regular expression replacements
	baseSubstitutesRegex := []string{
		`/(?m)^\s*#!\s*$/`,
		`/"!Q!(?P<content>[\S]+?)!Q!"/${content}`,
		`/"!Q!(?P<content>.*?)!Q!"/"${content}"`,
	}
	if substitutesFromEnv := os.Getenv(EnvSubstitutes); substitutesFromEnv != "" {
		baseSubstitutesRegex = append(baseSubstitutesRegex, strings.Split(substitutesFromEnv, "\n")...)
	}
	t := &Template{
		Template:       template.New(name),
		options:        new(ExtendedOption),
		optionsEnabled: new(ExtendedOption),
		substitutes:    utils.InitReplacers(baseSubstitutesRegex...),
	}
	return t.Delims().SetContext(new(map[string]interface{}))
}

var templateMutex sync.Mutex

// Template let us extend the functionalities of base go template library.
type Template struct {
	*template.Template
	tempFolder     string
	substitutes    []utils.RegexReplacer
	context        interface{}
	constantKeys   []interface{}
	razor          string
	parent         *Template
	folder         string
	children       map[string]*Template
	aliases        funcTableMap
	functions      funcTableMap
	options        *ExtendedOption
	optionsEnabled *ExtendedOption
}

// RazorDelim returns the razor delimiter.
func (t *Template) RazorDelim() string { return t.razor }

// Enabled returns true if the specified options are all set.
func (t *Template) Enabled(o ...ExtendedOption) bool { return t.options.IsSet(o...) }

// Disabled returns true if the specified options are all unset
func (t *Template) Disabled(o ...ExtendedOption) bool { return !t.Enabled(o...) }

// SetContext set the default context for the template ($ variable).
func (t *Template) SetContext(context interface{}) *Template {
	t.context = context
	return t
}

// TempFolder set temporary folder used by this template.
func (t *Template) TempFolder(folder string) *Template {
	t.tempFolder = folder
	return t
}

// Delims allows setting all delimiters (left, right and razor) in one call.
func (t *Template) Delims(delims ...string) *Template {
	switch len(delims) {
	case 0:
		return t.Delims("", "", "")
	case 1:
		return t.Delims(delims[0], "", "")
	case 2:
		return t.Delims(delims[0], delims[1], "")
	case 3:
		t.razor = coalesce(delims[2], t.RazorDelim(), "@").(string)
		t.Template.Delims(
			coalesce(delims[0], t.LeftDelim(), "{{").(string),
			coalesce(delims[1], t.RightDelim(), "{{").(string),
		)
		return t
	default:
		panic(fmt.Errorf("Too many delimiters supplied (max 3): %v", delims))
	}
}

// Option sets options for the template.
func (t *Template) Option(options ...interface{}) *Template {
	for _, opt := range options {
		switch opt := opt.(type) {
		case ExtendedOption:
			t.options.Set(opt, true)
		default:
			t.Template.Option(opt)
		}
	}
	return t
}

// Replacers add a series of regular expressions to apply on template before and after evaluations.
func (t *Template) Replacers(replacers ...string) *Template {
	t.substitutes = append(t.substitutes, utils.InitReplacers(replacers...)...)
	return t
}

// Parse parses text as a template body for t.
// Named template definitions ({{define ...}} or {{block ...}} statements) in text
// define additional templates associated with t and are removed from the
// definition of t itself.
//
// Templates can be redefined in successive calls to Parse.
// A template definition with a body containing only white space and comments
// is considered empty and will not replace an existing template's body.
// This allows using Parse to add new named template definitions without
// overwriting the main template body.
func (t *Template) Parse(text string) (*Template, error) {
	_, err := t.Template.Parse(text)
	return t, err
}

// ParseFiles parses the named files and associates the resulting templates with
// t. If an error occurs, parsing stops and the returned template is nil;
// otherwise it is t. There must be at least one file.
// Since the templates created by ParseFiles are named by the base
// names of the argument files, t should usually have the name of one
// of the (base) names of the files. If it does not, depending on t's
// contents before calling ParseFiles, t.Execute may fail. In that
// case use t.ExecuteTemplate to execute a valid template.
//
// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
	_, err := t.Template.ParseFiles(filenames...)
	return t, err
}

// NewTemplate creates an Template object with default initialization.
// func NewTemplate(folder string, context interface{}, delimiters string, options ExtendedOption, substitutes ...string) (result *Template, err error) {
// 	defer func() {
// 		if rec := recover(); rec != nil {
// 			result, err = nil, fmt.Errorf("%v", rec)
// 		}
// 	}()
// 	if options == 0 {
// 		options = DefaultOptions
// 	}

// 	t := Must(New("Main").Parse(""))
// 	if acceptNoValue {
// 		t.Option(AcceptNoValue)
// 	}
// 	if strictError {
// 		t.Option(StrictErrorCheck)
// 	}
// 	t.folder, _ = filepath.Abs(iif(folder != "", folder, utils.Pwd()).(string))
// 	t.context = iif(context != nil, context, collections.CreateDictionary())
// 	t.aliases = make(funcTableMap)

// 	// Set the regular expression replacements
// 	baseSubstitutesRegex := []string{`/(?m)^\s*#!\s*$/`}
// 	if substitutesFromEnv := os.Getenv(EnvSubstitutes); substitutesFromEnv != "" {
// 		baseSubstitutesRegex = append(baseSubstitutesRegex, strings.Split(substitutesFromEnv, "\n")...)
// 	}
// 	t.substitutes = utils.InitReplacers(append(baseSubstitutesRegex, substitutes...)...)

// 	if t.Enabled(Extension) {
// 		t.initExtension()
// 	}

// 	// Set the options supplied by caller
// 	t.init("")
// 	if delimiters != "" {
// 		t.Delims(strings.Split(delimiters, ",")...)
// 	}
// 	return t, nil
// }

func (t *Template) init2() {
}

// Initialize a new template with same attributes as the current context.
func (t *Template) init(folder string) {
	if folder != "" {
		t.folder, _ = filepath.Abs(folder)
	}
	if *t.options == 0 {
		t.Option(DefaultOptions)
	}
	if t.Options() == 0 {
		t.Option(AllOptions)
	}
	t.addFuncs()
	t.children = make(map[string]*Template)
	t.setConstant(false, "\n", "NL", "CR", "NEWLINE")
	t.setConstant(false, true, "true")
	t.setConstant(false, false, "false")
	t.setConstant(false, nil, "null")
}

// GetNewContext returns a distinct context for each folder.
func (t *Template) GetNewContext(folder string, useCache bool) *Template {
	folder = iif(folder != "", folder, t.folder).(string)
	if context, found := t.children[folder]; useCache && found {
		return context
	}

	newTemplate := *t
	newTemplate.Template = template.New(folder)
	newTemplate.addFunctions(t.functions)
	newTemplate.addFunctions(t.aliases)
	newTemplate.init(folder)
	newTemplate.parent = t
	newTemplate.importTemplates(t)
	newTemplate.options = new(ExtendedOption)
	*newTemplate.options = *t.options
	if dict := t.Context(); dict.Len() > 0 {
		newTemplate.context = dict.Clone()
	}

	if !useCache {
		return &newTemplate
	}
	// We register the new template as a child of the main template
	t.children[folder] = &newTemplate
	return t.children[folder]
}

// IsCode determines if the supplied code appears to have gotemplate code.
func (t *Template) IsCode(code string) bool {
	return !strings.Contains(code, noGoTemplate) && (t.IsRazor(code) || strings.Contains(code, t.LeftDelim()) && strings.Contains(code, t.RightDelim()))
}

// IsRazor determines if the supplied code appears to have Razor code.
func (t *Template) IsRazor(code string) bool {
	return strings.Contains(code, t.RazorDelim()) && !strings.Contains(code, noGoTemplate) && !strings.Contains(code, noRazor)
}

func (t *Template) isTemplate(file string) bool {
	for i := range templateExt {
		if strings.HasSuffix(file, templateExt[i]) {
			return true
		}
	}
	return false
}

func (t *Template) initExtension() {
	ext := t.GetNewContext("", false)
	t.constantKeys = ext.constantKeys
	*ext.options = DefaultOptions

	var extensionfiles []string
	if extensionFolders := strings.TrimSpace(os.Getenv(EnvExtensionPath)); extensionFolders != "" {
		for _, path := range strings.Split(extensionFolders, string(os.PathListSeparator)) {
			if path != "" {
				files, _ := utils.FindFilesMaxDepth(path, ExtensionDepth, false, "*.gte")
				extensionfiles = append(extensionfiles, files...)
			}
		}
	}
	extensionfiles = append(extensionfiles, utils.MustFindFilesMaxDepth(ext.folder, ExtensionDepth, false, "*.gte")...)

	// Retrieve the template extension files
	for _, file := range extensionfiles {
		// We just load all the template files available to ensure that all template definition are loaded
		// We do not use ParseFiles because it names the template with the base name of the file
		// which result in overriding templates with the same base name in different folders.
		content := string(must(ioutil.ReadFile(file)).([]byte))

		// We execute the content, but we ignore errors. The goal is only to register the sub templates and aliases properly
		// We also do not ask to clone the context as we wish to let extension to be able to alter the supplied context
		if _, _, err := ext.processContentInternal(content, file, nil, 0, false, nil); err != nil {
			InternalLog.Error(err)
		}
	}

	// Add the children contexts to the main context
	for _, context := range ext.children {
		t.importTemplates(context)
	}

	// We reset the list of templates
	t.children = make(map[string]*Template)
}

func (t *Template) setConstant(stopOnFirst bool, value interface{}, names ...string) {
	c, err := collections.TryAsDictionary(t.context)
	if err != nil {
		return
	}

	context := c.AsMap()
	for i := range names {
		if val, isSet := context[names[i]]; !isSet {
			context[names[i]] = value
			t.constantKeys = append(t.constantKeys, names[i])
			if stopOnFirst {
				return
			}
		} else if isSet && reflect.DeepEqual(value, val) {
			return
		}
	}
}

// Import templates from another template.
func (t *Template) importTemplates(source *Template) {
	for _, subTemplate := range source.Templates() {
		if subTemplate.Name() != subTemplate.ParseName {
			t.AddParseTree(subTemplate.Name(), subTemplate.Tree)
		}
	}
}

// Add allows adding a value to the template context.
// The context must be a dictionnary to use that function, otherwise, it will panic.
func (t *Template) Add(key string, value interface{}) {
	collections.AsDictionary(t.context).Add(key, value)
}

// Merge allows adding multiple values to the template context.
// The context and values must both be dictionnary to use that function, otherwise, it will panic.
func (t *Template) Merge(values interface{}) {
	collections.AsDictionary(t.context).Add(key, collections.AsDictionary(values))
}

// Context returns the template context as a dictionnary if possible, otherwise, it returns null.
func (t *Template) Context() (result collections.IDictionary) {
	if result, _ = collections.TryAsDictionary(t.context); result == nil {
		result = collections.CreateDictionary()
	}
	return
}

// ParseGlob parses the template definitions in the files identified by the
// pattern and associates the resulting templates with t. The files are matched
// according to the semantics of filepath.Match, and the pattern must match at
// least one file. ParseGlob is equivalent to calling t.ParseFiles with the
// list of files matched by the pattern.
//
// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
func ParseGlob(pattern string) (*Template, error) {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(filenames) == 0 {
		return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
	}
	return ParseFiles(filenames...)
}

// ParseFiles parses the named files and associates the resulting templates with
// t. If an error occurs, parsing stops and the returned template is nil;
// otherwise it is t. There must be at least one file.
// Since the templates created by ParseFiles are named by the base
// names of the argument files, t should usually have the name of one
// of the (base) names of the files. If it does not, depending on t's
// contents before calling ParseFiles, t.Execute may fail. In that
// case use t.ExecuteTemplate to execute a valid template.
//
// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
func ParseFiles(filenames ...string) (t *Template, err error) {
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("template: no files named in call to ParseFiles")
	}
	defer func() {
		switch rec := recover().(type) {
		case nil:
		case error:
			err = rec
		default:
			err = fmt.Errorf("%v", rec)
		}
	}()
	t, err = New(filepath.Base(filenames[0])).ParseFiles(filenames...)
	return
}
