package gen

import (
	"go/ast"
)

type expr interface {
	*ast.Ident | *ast.InterfaceType | *ast.StarExpr | *ast.ArrayType | *ast.SelectorExpr | *ast.MapType | *ast.FuncType | *ast.ChanType | *ast.IndexExpr | *ast.UnaryExpr
}

type pointer struct {
	next *pointer
}

func (slf pointer) String() string {
	p := ""
	if slf.next != nil {
		p = slf.next.String()
	}

	return "*" + p
}

type value[T expr] struct {
	name string
	ptr  *pointer
	expr T
}

func (slf value[T]) String() string {
	return slf.StringPtr() + slf.name
}

func (slf value[T]) StringPtr() string {
	p := ""
	if slf.ptr != nil {
		p = slf.ptr.String()
	}

	return p
}

func deref[T expr](expr ast.Expr) *value[T] {
	switch e := expr.(type) {
	case *ast.StarExpr:
		v := deref[T](e.X)
		if v == nil {
			return nil
		}

		ex := cast[T](v.expr)
		if ex != nil {
			return &value[T]{
				name: v.name,
				ptr:  &pointer{next: v.ptr},
				expr: ex,
			}
		}
	case *ast.Ident:
		ex := cast[T](e)
		if ex != nil {
			return &value[T]{
				name: e.Name,
				expr: ex,
			}
		}
	case *ast.ArrayType:
		ex := cast[T](e)
		if ex != nil {
			return &value[T]{
				expr: ex,
			}
		}
	case *ast.SelectorExpr:
		ex := cast[T](e)
		if ex != nil {
			v := deref[*ast.Ident](e.X)
			if v != nil {
				return &value[T]{
					name: v.name + "." + e.Sel.Name,
					expr: ex,
				}
			}
		}
	case *ast.InterfaceType:
		ex := cast[T](e)
		if ex != nil {
			return &value[T]{
				name: "interface{}",
				expr: ex,
			}
		}
	case *ast.MapType:
		ex := cast[T](e)
		if ex != nil {
			return &value[T]{
				name: "",
				expr: ex,
			}
		}
	case *ast.FuncType:
		ex := cast[T](e)
		if ex != nil {
			return &value[T]{
				name: "",
				expr: ex,
			}
		}
	case *ast.ChanType:
		ex := cast[T](e)
		if ex != nil {
			return &value[T]{
				name: "",
				expr: ex,
			}
		}
	case *ast.IndexExpr:
		ex := cast[T](e)
		if ex != nil {
			return &value[T]{
				name: "",
				expr: ex,
			}
		}
	case *ast.UnaryExpr:
		ex := cast[T](e)
		if ex != nil {
			return &value[T]{
				name: "",
				expr: ex,
			}
		}
	}

	return nil
}

func cast[T expr](ast any) T {
	v, ok := ast.(T)
	if !ok {
		return nil
	}

	return v
}
