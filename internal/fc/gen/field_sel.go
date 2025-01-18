package gen

import (
	"go/ast"
)

type FieldSel struct {
	imports map[string]string
}

func newFieldSel(imports map[string]string) *FieldSel {
	return &FieldSel{imports}
}

func (slf *FieldSel) Generate(
	fieldNames []*ast.Ident,
	fieldType ast.Expr,
) (
	fields []field,
	imports map[string]string,
) {
	v := deref[*ast.SelectorExpr](fieldType)
	if v == nil {
		return
	}

	ident, ok := v.expr.X.(*ast.Ident)
	if !ok {
		return
	}

	imports = make(map[string]string, 10)

	for _, fieldName := range fieldNames {
		if !ast.IsExported(fieldName.Name) {
			continue
		}

		i, ok := slf.imports[ident.String()]
		if !ok {
			continue
		}

		imports[ident.String()] = i
		fields = append(fields, field{
			Name: fieldName.Name,
			Type: v.String(),
		})
	}

	return
}
