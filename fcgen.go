package main

import (
	"flag"
	"fmt"
	"log"
	"maps"
	"os"
	"path"
	"strings"

	"factory/internal/fc"
	"factory/internal/fc/gen"
)

var (
	flagTypes = flag.String("type", "", "comma-separated list of type names; must be set")
	flagDir   = flag.String("out", "", "output directory; default fc/<type>_gen.go for each type")
	flagTags  = flag.String("tags", "", "comma-separated list of build tags to apply")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of fcgen:\n")
	fmt.Fprintf(os.Stderr, "\tfcgen [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "Arguments:\n")
	fmt.Fprintf(os.Stderr, "\t[directory] if no set, '.' is uses\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("fcgen: ")

	flag.Usage = Usage
	flag.Parse()

	if len(*flagTypes) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	types := strings.Split(*flagTypes, ",")

	var tags []string
	if len(*flagTags) > 0 {
		tags = strings.Split(*flagTags, ",")
	}

	args := []string{"."}
	if len(flag.Args()) > 0 {
		args = flag.Args()
	}

	pkgs, err := fc.NewPackages(args, tags)
	if err != nil {
		log.Fatalf("new package: %v", err)
	}

	dir := "fc"
	if flagDir != nil && *flagDir != "" {
		dir = *flagDir
	}

	founds := make(map[string]struct{}, 10)
	leftovers := make(map[string]struct{}, 10)

	for _, pkg := range pkgs {
		for _, typeName := range types {
			defs := make(map[string]struct{}, len(pkg.Files))
			for _, file := range pkg.Files {
				maps.Copy(defs, fc.CollectDefs(file.Ast))
			}

			for _, file := range pkg.Files {
				var st *fc.Struct
				structs := fc.CollectStructs(file.Ast)

				for _, s := range structs {
					if s.Name == typeName {
						st = &s
						break
					}
				}

				if st != nil {
					g := gen.
						NewGenerator(pkg, "fc_"+pkg.Name, file.Imports, st.Params, defs).
						Generate(*st)

					write(g, dir, typeName)

					founds[typeName] = struct{}{}
					delete(leftovers, typeName)

					continue
				}

				_, ok := founds[typeName]
				if !ok {
					leftovers[typeName] = struct{}{}
				}
			}
		}
	}

	if len(leftovers) > 0 {
		str := make([]string, 0, len(leftovers))
		for t := range leftovers {
			str = append(str, fmt.Sprintf("'%s'", t))
		}

		log.Fatalf("those types not found %s", strings.Join(str, ", "))
	}
}

func write(
	gen interface {
		Format() ([]byte, error)
	},
	dir,
	typeName string,
) {
	file := path.Join(dir, strings.ToLower(typeName)+"_gen.go")

	bytes, err := gen.Format()
	if err != nil {
		log.Fatalf("format code: %v", err)
	}

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("creating directory %q: %v", dir, err)
	}

	err = os.WriteFile(file, bytes, 0644)
	if err != nil {
		log.Fatalf("creating file %q: %v", file, err)
	}
}
