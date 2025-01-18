package fc

import (
	"go/ast"
)

type Struct struct {
	Name   string
	Type   *ast.StructType
	Params *ast.FieldList
}

func CollectDefs(node ast.Node) map[string]struct{} {
	defs := make(map[string]struct{}, 10)

	ast.Inspect(node, func(n ast.Node) bool {
		x, ok := n.(*ast.TypeSpec)
		if ok {
			defs[x.Name.Name] = struct{}{}
		}

		return true
	})

	return defs
}

func CollectStructs(node ast.Node) []Struct {
	structs := make([]Struct, 0, 10)

	ast.Inspect(node, func(n ast.Node) bool {
		var expr ast.Expr
		var stName string
		var params *ast.FieldList

		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Type == nil {
				return true
			}

			stName = x.Name.Name
			expr = x.Type
			params = x.TypeParams

		case *ast.CompositeLit:
			expr = x.Type
		case *ast.ValueSpec:
			stName = x.Names[0].Name
			expr = x.Type
		}

		// if expression is in form "*T" or "[]T", dereference to check if "T"
		// contains a struct expression
		expr = deref(expr)

		x, ok := expr.(*ast.StructType)
		if !ok {
			return true
		}

		if !ast.IsExported(stName) {
			return true
		}

		structs = append(structs, Struct{
			Name:   stName,
			Type:   x,
			Params: params,
		})

		return false
	})

	return structs
}

func deref(expr ast.Expr) ast.Expr {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return deref(t.X)
	case *ast.ArrayType:
		return deref(t.Elt)
	}

	return expr
}
