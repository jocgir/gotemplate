package tests

import (
	"github.com/coveo/gotemplate/collections"
	impl "github.com/coveo/gotemplate/collections/implementation"
	"github.com/coveo/gotemplate/hcl"
	"github.com/coveo/gotemplate/json"
	"github.com/coveo/gotemplate/yaml"
)

type (
	dictionary = map[string]interface{}
	str        = collections.String
)

// Imported from parent
var (
	IIf                    = collections.IIf
	IfUndef                = collections.IfUndef
	IsEmptyValue           = collections.IsEmptyValue
	IsExported             = collections.IsExported
	ToNativeRepresentation = collections.ToNativeRepresentation
	genHelper              = impl.DictionaryHelper
	hclHelper              = hcl.DictionaryHelper
	jsonHelper             = json.DictionaryHelper
	yamlHelper             = yaml.DictionaryHelper
)
