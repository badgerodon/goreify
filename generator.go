package main

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	"path"
	"reflect"
	"strings"

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
		switch t := decl.(type) {
		case *ast.FuncDecl:
			isGenericFunction := false
			isMethod := false
			isGlobalFunction := false
			if t.Name.Name == entity {
				isGenericFunction = true
			} else {
				if t.Recv != nil && len(t.Recv.List) > 0 {
					for _, ident := range g.getIdentifiers(t.Recv) {
						if ident.Name == entity {
							isMethod = true
							break
						}
					}
				}

				if t.Type.Results != nil && len(t.Type.Results.List) > 0 {
					for _, ident := range g.getIdentifiers(t.Type.Results) {
						if ident.Name == entity {
							isGlobalFunction = true
							break
						}
					}
				}
			}

			if isGenericFunction || isMethod || isGlobalFunction {
				err := g.GenerateFromFunction(t, entity, cfg,
					// we update the types inside the function if this is a
					// method or a global function
					isMethod || isGlobalFunction,
					// we rename the function if its generic or uses the generic
					// type
					(isGenericFunction || isGlobalFunction) && !isMethod)
				if err != nil {
					return err
				}
			}
		case *ast.GenDecl:
			if t.Tok == token.TYPE {
				for _, s := range t.Specs {
					spec := s.(*ast.TypeSpec)
					if spec.Name.Name != entity {
						continue
					}

					err := g.GenerateFromType(file, spec, cfg)
					if err != nil {
						return err
					}
				}
			}
		default:
			//log.Printf("unhandled decl: %T, %v", decl, decl)
		}
	}
	return nil
}

// GenerateFromFunction generates definitions from the given function
func (g *Generator) GenerateFromFunction(
	f *ast.FuncDecl, entity string, cfg *ReifyConfig,
	updateTypes, updateName bool,
) error {
	for _, reified := range cfg.Permutations() {
		if updateTypes {
			ids := g.getIdentifiers(f)
			for i := range ids {
				id := ids[i]
				if id.Name == entity {
					originalName := id.Name
					id.Name = originalName + reified.NameExtension()
					g.todo.Append(func() {
						id.Name = originalName
					})
				}
			}
		}

		if updateName {
			originalName := f.Name.Name
			f.Name.Name = originalName + reified.NameExtension()
			g.todo.Append(func() {
				f.Name.Name = originalName
			})
			if f.Doc != nil && len(f.Doc.List) > 0 {
				originalDoc := f.Doc.List[0].Text
				idx := strings.Index(originalDoc, originalName)
				if idx >= 0 {
					f.Doc.List[0].Text = originalDoc[:idx] +
						f.Name.Name +
						originalDoc[idx+len(originalName):]
					g.todo.Append(func() {
						f.Doc.List[0].Text = originalDoc
					})
				}
			}
		}

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

// GenerateFromType reifies a type definition
func (g *Generator) GenerateFromType(file *ast.File, spec *ast.TypeSpec, cfg *ReifyConfig) error {
	for _, reified := range cfg.Permutations() {
		// generate the type
		originalName := spec.Name.Name
		spec.Name.Name = spec.Name.Name + reified.NameExtension()
		g.todo.Append(func() {
			spec.Name.Name = originalName
		})

		var nodestack []ast.Node
		ast.Inspect(spec, func(n ast.Node) bool {
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

		printer.Fprint(&g.code, g.fset, []ast.Decl{
			&ast.GenDecl{
				Tok:   token.TYPE,
				Specs: []ast.Spec{spec},
			},
		})
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

func (g *Generator) getIdentifiers(node ast.Node) []*ast.Ident {
	if node == nil {
		return nil
	}
	var ids []*ast.Ident
	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		if id, ok := n.(*ast.Ident); ok {
			ids = append(ids, id)
		}
		return true
	})
	return ids
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
