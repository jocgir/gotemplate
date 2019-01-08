package template

import (
	"path/filepath"

	"github.com/coveo/gotemplate/collections"
	"github.com/coveo/gotemplate/errors"
	"github.com/coveo/gotemplate/utils"
)

// Imported from subpackages
var (
	AsList        = AsList
	AsStdString   = collections.AsStdString
	AsString      = collections.AsString
	Default       = collections.Default
	FormatMessage = collections.FormatMessage
	IIf           = collections.IIf
	IfUndef       = collections.IfUndef
	Must          = errors.Must
	Split2        = Split2
	Trap          = errors.Trap
)

type (
	// IGenericList imported from collections
	IGenericList = IGenericList

	// IDictionary imported from collections
	IDictionary = collections.IDictionary

	// String imported from collections
	String = collections.String
)

func getTargetFile(targetFile, sourcePath, targetPath string) string {
	if targetPath != "" {
		targetFile = filepath.Join(targetPath, utils.Relative(sourcePath, targetFile))
	}
	return targetFile
}
