package gen

import (
	"go/ast"
)

type FieldInterface struct{}

func newFieldInterface() *FieldInterface {
	return &FieldInterface{}
}

func (slf *FieldInterface) Generate(
	fieldNames []*ast.Ident,
	fieldType ast.Expr,
) (
	fields []field,
	imports map[string]string,
) {
	v := deref[*ast.InterfaceType](fieldType)
	if v == nil {
		return
	}

	if len(v.expr.Methods.List) > 0 {
		return
	}

	for _, fieldName := range fieldNames {
		if !ast.IsExported(fieldName.Name) {
			continue
		}

		fields = append(fields, field{
			Name: fieldName.Name,
			Type: v.String(),
		})
	}

	return
}
