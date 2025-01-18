package gen

import (
	"fmt"
	"go/ast"
	"maps"
	"strings"
)

type FieldComplex struct {
	generators []generator
}

func newFieldComplex(generators []generator) *FieldComplex {
	return &FieldComplex{generators: generators}
}

func (slf *FieldComplex) Generate(
	fieldNames []*ast.Ident,
	fieldType ast.Expr,
) (
	fields []field,
	imports map[string]string,
) {
	imports = make(map[string]string, 10)
	names := guessNames(fieldNames, fieldType)

	for _, fieldName := range names {
		if !ast.IsExported(fieldName.Name) {
			continue
		}

		f, i := slf.generate(fieldName, fieldType)
		if f != nil {
			maps.Copy(imports, i)
			fields = append(fields, field{
				Name: fieldName.String(),
				Type: *f,
			})
		}
	}

	return
}

func (slf *FieldComplex) generate(
	fieldName *ast.Ident,
	fieldType ast.Expr,
) (*string, map[string]string) {
	switch {
	case deref[*ast.FuncType](fieldType) != nil:
		return slf.genFunc(fieldName, fieldType)

	case deref[*ast.MapType](fieldType) != nil:
		return slf.genMap(fieldName, fieldType)

	case deref[*ast.ArrayType](fieldType) != nil:
		return slf.genArray(fieldName, fieldType)

	case deref[*ast.ChanType](fieldType) != nil:
		return slf.genChan(fieldName, fieldType)

	case deref[*ast.IndexExpr](fieldType) != nil:
		return slf.genIndex(fieldName, fieldType)

	case deref[*ast.UnaryExpr](fieldType) != nil:
		return slf.genUnary(fieldName, fieldType)
	}

	out := ""
	imports := make(map[string]string, 10)

	for _, g := range slf.generators {
		fields, imp := g.Generate([]*ast.Ident{fieldName}, fieldType)
		for _, f := range fields {
			out += f.Type
		}
		maps.Copy(imports, imp)
	}
	if out != "" {
		return &out, imports
	}

	return nil, nil
}

func (slf *FieldComplex) genFunc(
	fieldName *ast.Ident,
	fieldType ast.Expr,
) (*string, map[string]string) {
	v := deref[*ast.FuncType](fieldType)
	if v == nil {
		return nil, nil
	}

	inc := make([]string, 0, 10)
	out := make([]string, 0, 10)
	imports := make(map[string]string, 10)

	if v.expr.Params != nil {
		for _, l := range v.expr.Params.List {
			i0, i1 := slf.generate(fieldName, l.Type)
			if i0 == nil {
				continue
			}

			maps.Copy(imports, i1)

			names := make([]string, 0, len(l.Names))
			for _, n := range l.Names {
				names = append(names, n.Name)
			}

			inc = append(inc, fmt.Sprintf("%s %s", strings.Join(names, ", "), *i0))
		}
	}
	if v.expr.Results != nil {
		for _, l := range v.expr.Results.List {
			i0, i1 := slf.generate(fieldName, l.Type)
			if i0 == nil {
				continue
			}

			maps.Copy(imports, i1)

			names := make([]string, 0, len(l.Names))
			for _, n := range l.Names {
				names = append(names, n.Name)
			}

			out = append(out, fmt.Sprintf("%s %s", strings.Join(names, ", "), *i0))
		}
	}

	o := fmt.Sprintf("%sfunc(%s)(%s)", v.StringPtr(), strings.Join(inc, ", "), strings.Join(out, ", "))

	return &o, imports
}

func (slf *FieldComplex) genMap(
	fieldName *ast.Ident,
	fieldType ast.Expr,
) (*string, map[string]string) {
	v := deref[*ast.MapType](fieldType)
	if v == nil {
		return nil, nil
	}

	out := ""
	imports := make(map[string]string, 10)

	vk, i := slf.generate(fieldName, v.expr.Key)
	if vk != nil {
		maps.Copy(imports, i)
		out += fmt.Sprintf("%smap[%s]", v.StringPtr(), *vk)
	}

	vv, i := slf.generate(fieldName, v.expr.Value)
	if vv != nil {
		maps.Copy(imports, i)
		out += *vv
	}

	if vk != nil && vv != nil {
		return &out, imports
	}

	return nil, nil
}

func (slf *FieldComplex) genArray(
	fieldName *ast.Ident,
	fieldType ast.Expr,
) (*string, map[string]string) {
	v := deref[*ast.ArrayType](fieldType)
	if v == nil {
		return nil, nil
	}

	size := sizeValue(v.expr.Len)
	if size == nil {
		return nil, nil
	}

	vv, i := slf.generate(fieldName, v.expr.Elt)
	if vv != nil {
		out := fmt.Sprintf("%s[%s]%s", v.StringPtr(), *size, *vv)

		return &out, i
	}

	return nil, nil
}

func (slf *FieldComplex) genChan(
	fieldName *ast.Ident,
	fieldType ast.Expr,
) (*string, map[string]string) {
	v := deref[*ast.ChanType](fieldType)
	if v == nil {
		return nil, nil
	}

	send := ""
	recv := ""

	switch v.expr.Dir {
	case ast.SEND:
		send = "<-"
	case ast.RECV:
		recv = "<-"
	}

	vv, i := slf.generate(fieldName, v.expr.Value)
	if vv != nil {
		out := fmt.Sprintf("%s%schan%s %s", v.StringPtr(), recv, send, *vv)

		return &out, i
	}

	return nil, nil
}

func (slf *FieldComplex) genIndex(
	fieldName *ast.Ident,
	fieldType ast.Expr,
) (*string, map[string]string) {
	v := deref[*ast.IndexExpr](fieldType)
	if v == nil {
		return nil, nil
	}

	out := ""
	imports := make(map[string]string, 10)

	v0, i0 := slf.generate(fieldName, v.expr.X)
	if v0 != nil {
		out += *v0
		maps.Copy(imports, i0)
	}

	v1, i1 := slf.generate(fieldName, v.expr.Index)
	if v1 != nil {
		out += fmt.Sprintf("[%s]", *v1)
		maps.Copy(imports, i1)
	}

	if v0 != nil && v1 != nil {
		return &out, imports
	}

	return nil, nil
}

func (slf *FieldComplex) genUnary(
	fieldName *ast.Ident,
	fieldType ast.Expr,
) (*string, map[string]string) {
	v := deref[*ast.UnaryExpr](fieldType)
	if v == nil {
		return nil, nil
	}

	vv, i := slf.generate(fieldName, v.expr.X)
	if vv != nil {
		out := fmt.Sprintf("~%s", *vv)

		return &out, i
	}

	return nil, nil
}

func sizeValue(len ast.Expr) *string {
	literal, ok := len.(*ast.BasicLit)
	if ok {
		return &literal.Value
	}

	if len == nil {
		s := ""

		return &s
	}

	return nil
}

func guessNames(
	fieldNames []*ast.Ident,
	fieldType ast.Expr,
) []*ast.Ident {
	if len(fieldNames) != 0 {
		return fieldNames
	}

	v0 := deref[*ast.Ident](fieldType)
	if v0 != nil {
		return []*ast.Ident{
			ast.NewIdent(v0.name),
		}
	}

	v1 := deref[*ast.SelectorExpr](fieldType)
	if v1 != nil {
		return []*ast.Ident{
			ast.NewIdent(v1.expr.Sel.Name),
		}
	}

	return nil
}
