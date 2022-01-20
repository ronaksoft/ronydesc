package codeparse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func ExtractMessages(filename string) (messages, services []string, err error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return nil, nil, err
	}

	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue
			}
			for _, s := range d.Specs {
				switch s := s.(type) {
				case *ast.TypeSpec:
					_ = s.Name.Name
					switch t := s.Type.(type) {
					case *ast.StructType:
						if len(t.Fields.List) == 0 {
							continue
						}
						if x, ok := t.Fields.List[0].Type.(*ast.SelectorExpr); ok {
							if z, ok := x.X.(*ast.Ident); ok {
								if z.Name == "desc" && x.Sel.Name == "Message" {
									messages = append(messages, fmt.Sprintf("%s{}", s.Name.Name))
								} else if z.Name == "desc" && x.Sel.Name == "Service" {
									services = append(services, fmt.Sprintf("%s{}", s.Name.Name))
								}
							}
						}
					}
				}
			}
		}
	}

	return
}
