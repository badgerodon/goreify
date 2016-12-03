package main

import (
	"fmt"
	"os"
	"sort"
	"text/template"
)

var tpl = template.Must(template.New("").Parse(`
{{define "binary_op"}}
// {{.Name}} calls the '{{.Operator}}' operator
func {{.Name}}(x, y interface{}) interface{} {
    switch t := x.(type) {
{{range .Types}}
        case {{.}}:
            return t {{$.Operator}} {{if $.Shift}}To_uint(y){{else}}To_{{.}}(y){{end}}
{{end}}
    }
    panic(fmt.Sprintf("{{.Operator}} not implemented for %T", x))
}
{{end}}

{{define "comparison_op"}}
func {{.Name}}(x, y interface{}) bool {
	switch t := x.(type) {
{{range .Types}}
	case {{.}}:
		return t {{$.Operator}} To_{{.}}(y)
{{end}}
	}
    panic(fmt.Sprintf("{{.Operator}} not implemented for %T", x))
}
{{end}}

{{define "converter"}}
// To_{{.Name}} converts anything to a {{.Name}}
func To_{{.Name}}(x interface{}) {{.Name}} {
	switch t := x.(type) {
{{range .Types}}
		case {{.}}:
			{{if $.Name | eq "complex128"}}
				{{if . | eq "complex128"}}
					return t
				{{else if . | eq "complex64"}}
					return complex(float64(real(t)), float64(imag(t)))
				{{else}}
					return complex(To_float64(t), 0)
				{{end}}
			{{else if $.Name | eq "complex64"}}
				{{if . | eq "complex64"}}
					return t
				{{else if . | eq "complex128"}}
					return complex(float32(real(t)), float32(imag(t)))
				{{else}}
					return complex(To_float32(t), 0)
				{{end}}
			{{else}}
				{{if . | eq "complex128"}}
					return {{$.Name}}(real(t))
				{{else if . | eq "complex64"}}
					return {{$.Name}}(real(t))
				{{else}}
					return {{$.Name}}(t)
				{{end}}
			{{end}}
{{end}}
	}
	var o {{.Name}}
	fmt.Fscanln(strings.NewReader(fmt.Sprint(x)), &o)
	return o
}
{{end}}
`))

func expandTypes(types ...string) []string {
	var all []string
	for _, typ := range types {
		switch typ {
		case "integers":
			all = append(all,
				"int", "int8", "int16", "int32",
				"uint", "uint8", "uint16", "uint64",
				"uintptr")
		case "floats":
			all = append(all,
				"float32", "float64")
		case "complex values":
			all = append(all,
				"complex64", "complex128")
		case "string":
			all = append(all, "string")
		default:
			all = append(all, typ)
		}
	}
	sort.Strings(all)
	return all
}

func main() {
	var binaryDefs = []struct {
		Name       string
		Operator   string
		Token      string
		Types      []string
		Shift      bool
		Comparison bool
	}{
		{
			Name:     "Add",
			Operator: "+",
			Token:    "ADD",
			Types:    expandTypes("integers", "floats", "complex values", "string"),
		},
		{
			Name:     "Subtract",
			Operator: "-",
			Token:    "SUB",
			Types:    expandTypes("integers", "floats", "complex values"),
		},
		{
			Name:     "Multiply",
			Operator: "*",
			Token:    "MUL",
			Types:    expandTypes("integers", "floats", "complex values"),
		},
		{
			Name:     "Divide",
			Operator: "/",
			Token:    "QUO",
			Types:    expandTypes("integers", "floats", "complex values"),
		},
		{
			Name:     "Remainder",
			Operator: "%",
			Token:    "REM",
			Types:    expandTypes("integers"),
		},
		{
			Name:     "BitwiseAnd",
			Operator: "&",
			Token:    "AND",
			Types:    expandTypes("integers"),
		},
		{
			Name:     "BitwiseOr",
			Operator: "|",
			Token:    "OR",
			Types:    expandTypes("integers"),
		},
		{
			Name:     "BitwiseXor",
			Operator: "^",
			Token:    "XOR",
			Types:    expandTypes("integers"),
		},
		{
			Name:     "BitClear",
			Operator: "%^",
			Token:    "AND_NOT",
			Types:    expandTypes("integers"),
		},

		// shifts
		{
			Name:     "LeftShift",
			Operator: "<<",
			Token:    "SHL",
			Types:    expandTypes("integers"),
			Shift:    true,
		},
		{
			Name:     "RightShift",
			Operator: ">>",
			Token:    "SHR",
			Types:    expandTypes("integers"),
			Shift:    true,
		},

		// comparisons
		{
			Name:       "Less",
			Operator:   "<",
			Token:      "LSS",
			Types:      expandTypes("integers", "floats", "string"),
			Comparison: true,
		},
		{
			Name:       "LessOrEqual",
			Operator:   "<=",
			Token:      "LEQ",
			Types:      expandTypes("integers", "floats", "string"),
			Comparison: true,
		},
		{
			Name:       "Greater",
			Operator:   ">",
			Token:      "GTR",
			Types:      expandTypes("integers", "floats", "string"),
			Comparison: true,
		},
		{
			Name:       "GreaterOrEqual",
			Operator:   ">=",
			Token:      "GEQ",
			Types:      expandTypes("integers", "floats", "string"),
			Comparison: true,
		},
	}

	os.Stdout.WriteString("package generics\n\n")
	fmt.Fprintf(os.Stdout, `
import (
	"fmt"
	"go/token"
	"reflect"
	"strings"
)

func To_string(x interface{}) string {
	return fmt.Sprint(x)
}

func Equal(x, y interface{}) bool {
	return reflect.DeepEqual(x, y)
}

func NotEqual(x, y interface{}) bool {
	return !reflect.DeepEqual(x, y)
}
`)

	os.Stdout.WriteString("// BinaryOpTokens are all the binary operations\n")
	os.Stdout.WriteString("var BinaryOpTokens = map[string]token.Token{\n")
	os.Stdout.WriteString(`	"Equal": token.EQL,` + "\n")
	os.Stdout.WriteString(`	"NotEqual": token.NEQ,` + "\n")
	for _, def := range binaryDefs {
		fmt.Fprintf(os.Stdout, `	"%v": token.%v,`+"\n", def.Name, def.Token)
	}
	os.Stdout.WriteString("}\n")

	for _, def := range binaryDefs {
		tplnm := "binary_op"
		if def.Comparison {
			tplnm = "comparison_op"
		}
		tpl.ExecuteTemplate(os.Stdout, tplnm, def)
	}
	os.Stdout.WriteString("\n\n")

	type Converter struct {
		Name  string
		Types []string
	}
	var converters []Converter
	for _, name := range expandTypes("integers", "floats", "complex values") {
		converters = append(converters, Converter{
			Name:  name,
			Types: expandTypes("integers", "floats", "complex values"),
		})
	}
	for _, converter := range converters {
		tpl.ExecuteTemplate(os.Stdout, "converter", converter)
	}
	os.Stdout.WriteString("\n\n")

}
