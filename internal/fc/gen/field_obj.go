package gen

import (
	"fmt"
	"go/ast"
	"strings"
)

type FieldObj struct {
	pkgName  string
	pkgPath  string
	generics map[string]struct{}
	defs     map[string]struct{}
}

func newFieldObj(
	pkgPath string,
	params *ast.FieldList,
	defs map[string]struct{},
) *FieldObj {
	v := strings.Split(pkgPath, "/")
	gens := make(map[string]struct{}, 10)

	if params != nil {
		for _, l := range params.List {
			for _, n := range l.Names {
				gens[n.Name] = struct{}{}
			}
		}
	}

	return &FieldObj{
		pkgName:  v[len(v)-1],
		pkgPath:  pkgPath,
		generics: gens,
		defs:     defs,
	}
}

func (slf *FieldObj) Generate(
	fieldNames []*ast.Ident,
	fieldType ast.Expr,
) (
	fields []field,
	imports map[string]string,
) {
	v := deref[*ast.Ident](fieldType)
	if v == nil {
		return
	}

	if v.expr.Obj == nil {
		_, ok := slf.defs[v.name]
		if !ok {
			return
		}

		v.expr.Obj = ast.NewObj(ast.Typ, v.name)
	}

	if !ast.IsExported(v.expr.Name) {
		return
	}

	if v.expr.Obj.Kind != ast.Typ {
		return
	}

	imports = make(map[string]string, 10)

	for _, fieldName := range fieldNames {
		if !ast.IsExported(fieldName.Name) {
			continue
		}

		typ := fmt.Sprintf("%s%s.%s", v.StringPtr(), slf.pkgName, v.name)

		_, ok := slf.generics[v.name]
		if ok {
			typ = fmt.Sprintf("%s%s", v.StringPtr(), v.name)
		}

		imports[slf.pkgName] = slf.pkgPath
		fields = append(fields, field{
			Name: fieldName.Name,
			Type: typ,
		})
	}

	return
}
