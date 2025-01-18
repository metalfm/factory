package fc

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Package struct {
	Name  string
	Path  string
	Defs  map[*ast.Ident]types.Object
	Files []*File
}

func NewPackages(
	patterns,
	tags []string,
) ([]*Package, error) {
	cfg := &packages.Config{
		Mode:       packages.NeedName | packages.NeedImports | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedFiles,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, ","))},
		Tests:      false,
	}

	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, fmt.Errorf("load packages: %w", err)
	}

	if len(pkgs) == 0 {
		return nil, fmt.Errorf("package not found by '%s': %w", patterns, err)
	}

	out := make([]*Package, 0, len(pkgs))

	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			return nil, fmt.Errorf("load package: %w", pkg.Errors[0])
		}

		p := &Package{
			Name:  pkg.Name,
			Path:  pkg.PkgPath,
			Defs:  pkg.TypesInfo.Defs,
			Files: make([]*File, len(pkg.Syntax)),
		}

		for i, a := range pkg.Syntax {
			imports := make(map[string]string, len(a.Imports))

			for _, spec := range a.Imports {
				var name string

				if spec.Name != nil {
					name = spec.Name.Name
				} else {
					paths := strings.Split(spec.Path.Value, "/")
					name = paths[len(paths)-1]
				}

				switch name {
				case "_", ".":
					continue
				default:
					alias := strings.Replace(name, "\"", "", -1)
					path := strings.Replace(spec.Path.Value, "\"", "", -1)

					imports[alias] = path
				}
			}

			p.Files[i] = &File{
				Ast:     a,
				Pkg:     p,
				Imports: imports,
			}
		}

		out = append(out, p)
	}

	return out, nil
}
