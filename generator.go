package main

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	"path"
	"reflect"

	"github.com/badgerodon/goreify/generics"
)

// A Generator generates code
type Generator struct {
	pkgname string
	imports map[string]string
	code    bytes.Buffer
	fset    *token.FileSet

	todo FunctionStack
}

// NewGenerator creates a new Generator
func NewGenerator(pkgname string, fset *token.FileSet) *Generator {
	return &Generator{
		pkgname: pkgname,
		imports: map[string]string{},
		fset:    fset,
	}
}

// GenerateFromFile generates definitions from the given file
func (g *Generator) GenerateFromFile(file *ast.File, entity string, cfg *ReifyConfig) error {
	for _, i := range file.Imports {
		value := i.Path.Value
		if len(value) > 0 && value[0] == '"' {
			value = value[1:]
		}
		if len(value) > 0 && value[len(value)-1] == '"' {
			value = value[:len(value)-1]
		}
		name := path.Base(value)
		if i.Name != nil {
			name = i.Name.Name
		}
		g.imports[name] = value
	}

	for _, decl := range file.Decls {
		switch decl.(type) {
		case *ast.FuncDecl:
			f := decl.(*ast.FuncDecl)
			if f.Name.Name != entity {
				continue
			}
			err := g.GenerateFromFunction(f, cfg)
			if err != nil {
				return err
			}
		default:
			//log.Printf("unhandled decl: %T, %v", decl, decl)
		}
	}
	return nil
}

// GenerateFromFunction generates definitions from the given function
func (g *Generator) GenerateFromFunction(f *ast.FuncDecl, cfg *ReifyConfig) error {
	for _, reified := range cfg.Permutations() {
		originalName := f.Name.Name
		originalComment := f.Doc
		f.Name.Name = f.Name.Name + "_" + reified.NameExtension()
		f.Doc = &ast.CommentGroup{}
		g.todo.Append(func() {
			f.Name.Name = originalName
			f.Doc = originalComment
		})

		var nodestack []ast.Node
		ast.Inspect(f, func(n ast.Node) bool {
			if n == nil {
				nodestack = nodestack[:len(nodestack)-1]
				return false
			}

			if expr, ok := n.(ast.Expr); ok {
				ne := g.transformExpression(reified, expr)
				if ne != nil {
					parent := nodestack[len(nodestack)-1]
					g.todo.Append(func() {
						g.replaceExpr(parent, ne, expr)
					})
					g.replaceExpr(parent, expr, ne)
				}
			}

			nodestack = append(nodestack, n)
			return true
		})

		printer.Fprint(&g.code, g.fset, []ast.Decl{f})
		g.code.WriteString("\n\n")

		// restore whatever we changed
		g.todo.Run()
	}
	return nil
}

// Export exports the generator to the writer
func (g *Generator) Export(w io.Writer) error {
	io.WriteString(w, "package "+g.pkgname+"\n\n")
	if len(g.imports) > 0 {
		io.WriteString(w, "import (\n")
		for name, path := range g.imports {
			io.WriteString(w, "\t"+name+" \""+path+"\"\n")
		}
		io.WriteString(w, ")\n\n")
	}
	_, err := io.Copy(w, &g.code)
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) transformExpression(reified ReifiedTypes, expr ast.Expr) ast.Expr {
	switch t := expr.(type) {
	case *ast.CallExpr:
		sel, ok := t.Fun.(*ast.SelectorExpr)
		if !ok {
			return nil
		}

		left, ok := sel.X.(*ast.Ident)
		if !ok {
			return nil
		}

		if left.Name != "generics" {
			return nil
		}

		binaryOp, ok := generics.BinaryOpTokens[sel.Sel.Name]
		if ok {
			return &ast.BinaryExpr{
				X:  t.Args[0],
				Y:  t.Args[1],
				Op: binaryOp,
			}
		}
	case *ast.SelectorExpr:
		left, ok := t.X.(*ast.Ident)
		if !ok {
			return nil
		}

		if left.Name != "generics" {
			return nil
		}

		newType, ok := reified[t.Sel.Name]
		if !ok {
			return nil
		}

		return &ast.Ident{
			NamePos: left.NamePos,
			Name:    newType,
		}
	}
	return nil
}

func (g *Generator) replaceExpr(node ast.Node, oldExpr, newExpr ast.Expr) {
	value := reflect.ValueOf(node)
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		var visit []reflect.Value
		if field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				visit = append(visit, field.Index(j))
			}
		} else {
			visit = append(visit, field)
		}

		for _, v := range visit {
			if v.CanSet() &&
				reflect.DeepEqual(v.Interface(), oldExpr) {
				v.Set(reflect.ValueOf(newExpr))
			}
		}
	}
}
