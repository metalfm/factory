package fc

import (
	"go/ast"
)

// File holds a single parsed file and associated data.
type File struct {
	Pkg     *Package  // Package to which this file belongs.
	Ast     *ast.File // Parsed AST.
	Imports map[string]string

	// These fields are reset for each type being generated.
	//typeName string // Name of the constant type.
	//values   []Value // Accumulator for constant values of that type.
}
