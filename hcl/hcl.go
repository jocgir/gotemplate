package hcl

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/coveo/gotemplate/v3/collections"
	"github.com/coveo/gotemplate/v3/collections/implementation"
	"github.com/coveo/gotemplate/v3/errors"
	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/hashicorp/hcl2/hclparse"
)

func (l hclList) String() string {
	result, err := MarshalInternal(l.AsArray())
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (d hclDict) String() string {
	result, err := Marshal(d.AsMap())
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (l hclList) PrettyPrint() string {
	result, _ := MarshalIndent(l.AsArray(), "", "  ")
	return string(result)
}

func (d hclDict) PrettyPrint() string {
	result, _ := MarshalIndent(d.AsMap(), "", "  ")
	return string(result)
}

var _ = func() int {
	collections.TypeConverters["hcl"] = Unmarshal
	return 0
}()

func unmarshal(content, filename string, retryOnError bool, out interface{}) (err error) {
	defer func() { err = errors.Trap(err, recover()) }()
	content = strings.TrimSpace(content)

	result, diag := hclparse.NewParser().ParseHCL([]byte(content), filename)
	if diag.HasErrors() {
		// In case of error, we try to format the code using hcl1 since hcl2 is less flexible regarding
		// formatting than its predecessor
		if formatedContent, err := printer.Format([]byte(content)); err == nil {
			if result2, diag2 := hclparse.NewParser().ParseHCL(formatedContent, filename); !diag2.HasErrors() {
				result = result2
				diag = diag2
			}
		}
		if diag.HasErrors() && retryOnError {
			// If there is still an error, it could be because we try to unmarshal a literal statement
			var temp hclDict
			if unmarshal("_="+content, filename, false, &temp) == nil {
				reflect.ValueOf(out).Elem().Set(reflect.ValueOf(temp["_"]))
				return nil
			}
		}
	}
	if diag.HasErrors() {
		fmt.Println("Error", diag)
		return diag
	}
	// switch body := result.Body.(type) {
	// case *hclsyntax.Body:
	// 	for _, a := range body.Attributes {
	// 		fmt.Println("a1", a.Name, "=", reflect.TypeOf(a.Expr))
	// 		switch a := a.Expr.(type) {
	// 		case *hclsyntax.ObjectConsExpr:
	// 			fmt.Println(a.ExprMap())
	// 			for _, aa := range a.Items {
	// 				fmt.Printf("\t%T %v <==> %T %v : %v\n", aa.KeyExpr, get(aa.KeyExpr), aa.ValueExpr, get(aa.ValueExpr), aa.ValueExpr.StartRange())
	// 			}
	// 		}
	// 	}
	// 	for _, b := range body.Blocks {
	// 		fmt.Println(b.Type, b.Labels, b.Body)
	// 	}
	// }

	// r1, e1 := hclsyntax.ParseConfig(result.Bytes, filename, hcl.Pos{Line: 1, Column: 1})
	// if e1.HasErrors() {
	// 	fmt.Println("r1", r1)
	// 	fmt.Println("e1", e1)
	// }

	fmt.Println("-===================")
	if tokens, diag := hclsyntax.LexConfig(result.Bytes, filename, hcl.Pos{Line: 1, Column: 1}); diag.HasErrors() {
		indent := ""
		for _, x := range tokens {
			switch x.Type {
			case hclsyntax.TokenNewline:
				continue
			case hclsyntax.TokenOBrace:
				indent += "    "
				continue
			case hclsyntax.TokenCBrace:
				indent = indent[0 : len(indent)-4]
				continue
			}
			fmt.Println(indent, x.Type, string(x.Bytes))
		}
	} else {
		fmt.Println(diag)
	}
	// hclsyntax.VisitAll(result.Body.(*hclsyntax.Body).Range(), func(node hclsyntax.Node) hcl.Diagnostics {
	// 	fmt.Printf("++ %[1]T %[1]v\n", node)
	// 	return nil
	// })
	// fmt.Println("----")
	// e2 := hclsyntax.Walk(r1, &testWalker{})
	// fmt.Println("e2", e2)
	// if err = gohcl.DecodeBody(result.Body, nil, out); err != nil {
	// 	switch body := result.Body.(type) {
	// 	case *hclsyntax.Body:
	// 		fmt.Println("attr")
	// 		for key, val := range body.Attributes {
	// 			v, ve := val.Expr.Value(nil)
	// 			fmt.Println(key, "=", v, ":", ve)
	// 		}
	// 		fmt.Println("blocks")
	// 		for i, b := range body.Blocks {
	// 			fmt.Println(i, b)
	// 		}
	// 	}
	// }
	transform(out)
	return
}

func get(e hclsyntax.Expression) string {
	x, _ := e.Value(nil)
	switch e := e.(type) {
	case *hclsyntax.ObjectConsKeyExpr:
		return x.AsString()
	case *hclsyntax.TemplateExpr:
		return x.AsString()
	case *hclsyntax.LiteralValueExpr:
		return fmt.Sprint(e.Val.AsBigFloat())
	default:
		return fmt.Sprintf("%T", e)
	}
}

type testWalker struct{}

func (w *testWalker) Enter(node hclsyntax.Node) hcl.Diagnostics {
	fmt.Printf("Enter %T\n", node)
	switch x := node.(type) {
	case *hclsyntax.LiteralValueExpr:
		fmt.Println(x.Val.AsString())

	case *hclsyntax.TemplateExpr:
		fmt.Println(x.Parts)

	default:
		fmt.Printf("UNKNOWN TYPE %T\n", x)
	}
	return nil
}

func (w *testWalker) Exit(node hclsyntax.Node) hcl.Diagnostics {
	fmt.Printf("Exit %T\n", node)
	return nil
}

// Unmarshal adds support to single array and struct representation
func Unmarshal(bs []byte, out interface{}) error {
	err := unmarshal(string(bs), "inline", true, out)
	return err
}

// Load loads hcl file into variable
func Load(filename string) (result interface{}, err error) {
	var content []byte
	if content, err = ioutil.ReadFile(filename); err == nil {
		err = unmarshal(string(content), filename, false, &result)
	}
	return
}

// Marshal serialize values to hcl format
// func Marshal(value interface{}) (result []byte, err error) {
// 	defer func() { err = errors.Trap(err, recover()) }()
// 	f := hclwrite.NewEmptyFile()
// 	gohcl.EncodeIntoBody(value, f.Body())
// 	return f.Bytes(), nil
// }

// Marshal serialize values to hcl format
func Marshal(value interface{}) ([]byte, error) {
	return MarshalIndent(value, "", "")
}

// MarshalIndent serialize values to hcl format with indentation
func MarshalIndent(value interface{}, prefix, indent string) ([]byte, error) {
	result, err := marshalHCL(collections.ToNativeRepresentation(value), true, true, prefix, indent)
	return []byte(result), err
}

// MarshalInternal serialize values to hcl format for result used in outer hcl struct
func MarshalInternal(value interface{}) ([]byte, error) {
	result, err := marshalHCL(collections.ToNativeRepresentation(value), false, false, "", "")
	return []byte(result), err
}

// MarshalTFVars serialize values to hcl format (without hcl map format)
func MarshalTFVars(value interface{}) ([]byte, error) { return MarshalTFVarsIndent(value, "", "") }

// MarshalTFVarsIndent serialize values to hcl format with indentation (without hcl map format)
func MarshalTFVarsIndent(value interface{}, prefix, indent string) ([]byte, error) {
	result, err := marshalHCL(collections.ToNativeRepresentation(value), false, true, prefix, indent)
	return []byte(result), err
}

// SingleContext converts array of 1 to single object otherwise, let the context unchanged
func SingleContext(context ...interface{}) interface{} {
	if len(context) == 1 {
		return context[0]
	}
	return context
}

type (
	helperBase = implementation.BaseHelper
	helperList = implementation.ListHelper
	helperDict = implementation.DictHelper
)

var needConversionImpl = implementation.NeedConversion

//go:generate genny -pkg=hcl -in=../collections/implementation/generic.go -out=generated_impl.go gen "ListTypeName=List DictTypeName=Dictionary base=hcl"
//go:generate genny -pkg=hcl -in=../collections/implementation/generic_test.go -out=generated_test.go gen "base=hcl"
