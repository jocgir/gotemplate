package template

import (
	"os"

	"github.com/coveooss/gotemplate/v3/collections"
	multicolor "github.com/coveooss/multilogger/color"
	"github.com/jocgir/template"
)

// Types aliases
type (
	// String is an alias to collections.String
	String = collections.String

	// Context is an alias for template.Context
	Context = template.Context
	// ContextSource is an alias for template.ContextSource
	ContextSource = template.ContextSource
	// ErrorAction is an alias for template.ErrorAction
	ErrorAction = template.ErrorAction
	// ErrorHandler is an alias for template.ErrorHandler
	ErrorHandler = template.ErrorHandler
	// ErrorManager is an alias for template.ErrorManager
	ErrorManager = template.ErrorManager
	// ErrorManagers is an alias for template.ErrorManagers
	ErrorManagers = template.ErrorManagers
	// ExecError is an alias for template.ExecError
	ExecError = template.ExecError
	// FuncMap is an alias for template.FuncMap
	FuncMap = template.FuncMap
	// Map is an alias for template.Map
	Map = template.Map
	// MissingAction is an alias for template.MissingAction
	MissingAction = template.MissingAction
	// Option is an alias for template.Option
	Option = template.Option
	// StackCall is an alias for template.StackCall
	StackCall = template.StackCall
)

// Aliases for imported functions
var (
	toStrings     = collections.ToStrings
	acceptNoValue = String(os.Getenv(EnvAcceptNoValue)).ParseBool()
	strictError   = String(os.Getenv(EnvStrictErrorCheck)).ParseBool()

	// Function aliases from github.com/coveooss/multilogger/color
	Print      = multicolor.Print
	Printf     = multicolor.Printf
	Println    = multicolor.Println
	ErrPrintf  = multicolor.ErrorPrintf
	ErrPrintln = multicolor.ErrorPrintln
	ErrPrint   = multicolor.ErrorPrint

	// Function aliases from github.com/jocgir/template
	HTMLEscape       = template.HTMLEscape
	HTMLEscapeString = template.HTMLEscapeString
	HTMLEscaper      = template.HTMLEscaper
	IsTrue           = template.IsTrue
	JSEscape         = template.JSEscape
	JSEscapeString   = template.JSEscapeString
	JSEscaper        = template.JSEscaper
	NewErrorManager  = template.NewErrorManager
	URLQueryEscaper  = template.URLQueryEscaper
)

// Aliases for imported contants
const (
	// Default delimiters
	DefaultLeftDelim  = template.DefaultLeftDelim
	DefaultRightDelim = template.DefaultRightDelim
	DefaultRazorDelim = "@"

	// Constants aliases from github.com/jocgir/template Option
	AllNativeOptions     = template.AllOptions
	Eval                 = template.Eval
	FlowControl          = template.FlowControl
	FunctionsAsMethods   = template.FunctionsAsMethods
	FunctionsWithContext = template.FunctionsWithContext
	NonStandardResults   = template.NonStandardResults
	Trap                 = template.Trap

	// Constants aliases from github.com/jocgir/template MissingAction
	Default   = template.Default
	Error     = template.Error
	Invalid   = template.Invalid
	ZeroValue = template.ZeroValue

	// Constants aliases from github.com/jocgir/template ErrorAction
	NoReplace      = template.NoReplace
	ResultAsArray  = template.ResultAsArray
	ResultReplaced = template.ResultReplaced

	// Constants aliases from github.com/jocgir/template ContextSource
	Call         = template.Call
	CallContext  = template.CallContext
	CallError    = template.CallError
	FieldError   = template.FieldError
	PrintContext = template.Print

	// Constants aliases from github.com/jocgir/template
	CallFailID       = template.CallFailID
	ContextID        = template.ContextID
	FuncsAsMethodsID = template.FuncsAsMethodsID
	NoValue          = template.NoValue
)

// Environment variables that could be defined to override default behaviors.
const (
	EnvAcceptNoValue    = "GOTEMPLATE_NO_VALUE"
	EnvStrictErrorCheck = "GOTEMPLATE_STRICT_ERROR"
	EnvSubstitutes      = "GOTEMPLATE_SUBSTITUTES"
	EnvDebug            = "GOTEMPLATE_DEBUG"
	EnvExtensionPath    = "GOTEMPLATE_PATH"

	// Enable/Disable gotemplate keywords.
	noGoTemplate       = "no-gotemplate!"
	noRazor            = "no-razor!"
	explicitGoTemplate = "force-gotemplate!"

	// Pause/Resume gotemplate/Razor
	pauseGoTemplate  = "gotemplate-pause!"
	resumeGoTemplate = "gotemplate-resume!"
	pauseRazor       = "razor-pause!"
	resumeRazor      = "razor-resume!"
)

// Common variables
var (
	// ExtensionDepth the depth level of search of gotemplate extension from the current directory (default = 2).
	ExtensionDepth  = 2
	defaultTemplate = New("")
)
