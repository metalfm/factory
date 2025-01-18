package gen

import (
	"go/ast"
)

type FieldSimple struct{}

func newFieldSimple() *FieldSimple {
	return &FieldSimple{}
}

func (slf *FieldSimple) Generate(
	fieldNames []*ast.Ident,
	fieldType ast.Expr,
) (
	fields []field,
	imports map[string]string,
) {
	v := deref[*ast.Ident](fieldType)
	if v == nil {
		return nil, nil
	}

	for _, fieldName := range fieldNames {
		if !ast.IsExported(fieldName.Name) {
			continue
		}

		if v.expr.Obj != nil {
			continue
		}

		fields = append(fields, field{
			Name: fieldName.Name,
			Type: v.String(),
		})
	}

	return
}
