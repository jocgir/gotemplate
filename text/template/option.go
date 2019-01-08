// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the code to handle template options.

package template

import (
	"reflect"
	"sort"
	"strings"

	"github.com/coveo/gotemplate/text/template/parse"
)

// missingKeyAction defines how to respond to indexing a map with a key that is not present.
type missingKeyAction int

const (
	mapInvalid   missingKeyAction = iota // Return an invalid reflect.Value.
	mapZeroValue                         // Return the zero value for the map element.
	mapError                             // Error out
)

// MissingHandler represents the function type used to try to recover missing key during the template evaluation.
type MissingHandler func(fieldName string, receiver reflect.Value, node parse.Node) (interface{}, MissingAction)

// MissingAction defines the action done by external handler when managing missing key.
type MissingAction int

const (
	// NoReplace is returned if the external handler has not been able to fix the missing key
	NoReplace MissingAction = iota

	// MissingReplaced is returned if the external handler returned a valid replacement for the missing key
	MissingReplaced

	// ReceiverIsArray is returned if the external handler returned an array on which we should apply the missing key
	ReceiverIsArray
)

type option struct {
	missingKey          missingKeyAction
	missingHandlers     map[string]MissingHandler
	missingHandlersKeys []string
}

// NoValue is the rendered string representation of invalid value if missingkey is set to invalid or left to default.
const NoValue = "<no value>"

// Option sets options for the template. Options are described by
// strings, either a simple string or "key=value". There can be at
// most one equals sign in an option string. If the option string
// is unrecognized or otherwise invalid, Option panics.
//
// Known options:
//
// missingkey: Control the behavior during execution if a map is
// indexed with a key that is not present in the map.
//	"missingkey=default" or "missingkey=invalid"
//		The default behavior: Do nothing and continue execution.
//		If printed, the result of the index operation is the string
//		"<no value>".
//	"missingkey=zero"
//		The operation returns the zero value for the map type's element.
//	"missingkey=error"
//		Execution stops immediately with an error.
//
func (t *Template) Option(opt ...string) *Template {
	t.init()
	for _, s := range opt {
		t.setOption(s)
	}
	return t
}

func (t *Template) setOption(opt string) {
	if opt == "" {
		panic("empty option string")
	}
	elems := strings.Split(opt, "=")
	switch len(elems) {
	case 2:
		// key=value
		switch elems[0] {
		case "missingkey":
			switch elems[1] {
			case "invalid", "default":
				t.option.missingKey = mapInvalid
				return
			case "zero":
				t.option.missingKey = mapZeroValue
				return
			case "error":
				t.option.missingKey = mapError
				return
			}
		}
	}
	panic("unrecognized option: " + opt)
}

// AddMissingHandler allows registration of function that could handle missing keys.
// This give the caller opportunity to recover errors and/or add custom values other that that default actions.
func (t *Template) AddMissingHandler(name string, handler MissingHandler) {
	t.option.addMissingHandler(name, handler)
}

func (o *option) addMissingHandler(name string, handler MissingHandler) {
	if o.missingHandlers == nil {
		o.missingHandlers = make(map[string]MissingHandler)
	}
	if handler == nil {
		delete(o.missingHandlers, name)
	} else {
		o.missingHandlers[name] = handler
	}

	o.missingHandlersKeys = make([]string, 0, len(o.missingHandlers))
	for key := range o.missingHandlers {
		o.missingHandlersKeys = append(o.missingHandlersKeys, key)
	}
	sort.Strings(o.missingHandlersKeys)
}

func (o *option) invoke(fieldName string, receiver reflect.Value, node parse.Node) (result reflect.Value, action MissingAction) {
	for _, key := range o.missingHandlersKeys {
		if value, missedAction := o.missingHandlers[key](fieldName, receiver, node); missedAction != NoReplace {
			return reflect.ValueOf(value), missedAction
		}
	}
	return
}
