package data

import (
	"github.com/coveo/gotemplate/v3/collections"
	"github.com/coveo/gotemplate/v3/errors"
	"github.com/coveo/gotemplate/v3/hcl"
	"github.com/coveo/gotemplate/v3/json"
	"github.com/coveo/gotemplate/v3/yaml"
)

// Various import from other modules
var (
	must                   = errors.Must
	TryAsList              = collections.TryAsList
	TryAsDictionary        = collections.TryAsDictionary
	AsDictionary           = collections.AsDictionary
	ToNativeRepresentation = collections.ToNativeRepresentation
)

// Import type libraries
type (
	hclDict  = hcl.Dictionary
	hclList  = hcl.List
	jsonDict = json.Dictionary
	jsonList = json.List
	yamlDict = yaml.Dictionary
	yamlList = yaml.List
)
