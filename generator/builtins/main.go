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
    switch x.(type) {
{{range .Types}}
        case {{.}}:
            return x.({{.}}) {{$.Operator}} y.({{.}})
{{end}}
    }
    panic(fmt.Sprintf("{{.Operator}} not implemented for %T", x))
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
	var defs = []struct {
		Name     string
		Operator string
		Token    string
		Types    []string
	}{
		{
			Name:     "Sum",
			Operator: "+",
			Token:    "ADD",
			Types:    expandTypes("integers", "floats", "complex values", "string"),
		},
		{
			Name:     "Difference",
			Operator: "-",
			Token:    "SUB",
			Types:    expandTypes("integers", "floats", "complex values"),
		},
		{
			Name:     "Product",
			Operator: "*",
			Token:    "MUL",
			Types:    expandTypes("integers", "floats", "complex values"),
		},
	}

	os.Stdout.WriteString("package generics\n\n")
	fmt.Fprintf(os.Stdout, `
import (
	"fmt"
	"go/token"
)
`)

	os.Stdout.WriteString("// BinaryOpTokens are all the binary operations\n")
	os.Stdout.WriteString("var BinaryOpTokens = map[string]token.Token{\n")
	for _, def := range defs {
		fmt.Fprintf(os.Stdout, `"%v": token.%v,`+"\n", def.Name, def.Token)
	}
	os.Stdout.WriteString("}\n")

	for _, def := range defs {
		tpl.ExecuteTemplate(os.Stdout, "binary_op", def)
	}
	os.Stdout.WriteString("\n\n")

}
