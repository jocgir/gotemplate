// Code generated by "stringer -type=Options -output generated_options.go"; DO NOT EDIT.

package template

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Razor-0]
	_ = x[Extension-1]
	_ = x[Math-2]
	_ = x[Sprig-3]
	_ = x[Data-4]
	_ = x[Logging-5]
	_ = x[Runtime-6]
	_ = x[Utils-7]
	_ = x[Net-8]
	_ = x[OS-9]
	_ = x[Git-10]
	_ = x[OptionOnByDefaultCount-11]
	_ = x[Overwrite-12]
	_ = x[OutputStdout-13]
	_ = x[RenderingDisabled-14]
	_ = x[AcceptNoValue-15]
	_ = x[StrictErrorCheck-16]
	_ = x[RazorCode-17]
}

const _Options_name = "RazorExtensionMathSprigDataLoggingRuntimeUtilsNetOSGitOptionOnByDefaultCountOverwriteOutputStdoutRenderingDisabledAcceptNoValueStrictErrorCheckRazorCode"

var _Options_index = [...]uint8{0, 5, 14, 18, 23, 27, 34, 41, 46, 49, 51, 54, 76, 85, 97, 114, 127, 143, 152}

func (i Options) String() string {
	if i < 0 || i >= Options(len(_Options_index)-1) {
		return "Options(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Options_name[_Options_index[i]:_Options_index[i+1]]
}
