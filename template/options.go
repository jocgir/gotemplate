package template

// ExtendedOption represents the additional options enabled by this template package.
type ExtendedOption uint32
type ExtendedOption2 uint32

// Options values
const (
	Razor ExtendedOption = 1 << iota
	Extension
	Math
	Sprig
	Data
	Logging
	Runtime
	Utils
	Net
	OS
	Git
	Overwrite
	OutputStdout
	RenderingDisabled
	AcceptNoValue
	StrictErrorCheck

	DefaultOptions = ExtendedOption(Git<<1 - 1)
	AllOptions     = ExtendedOption(StrictErrorCheck<<1 - 1)

	xxx ExtendedOption2 = iota
	yyy
	zzz
)

//go:generate bitenum -output generated_types.go -exclude FuncOptionsSet
