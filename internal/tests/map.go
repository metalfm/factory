package tests

import (
	"go/ast"
	"net/http"
	"time"
	time1 "time"

	"factory/internal/tests/entity"
)

//go:generate go run ../../fcgen.go -type=Map,MapCrazy

//nolint:unused //need to test unexported fields
type Map struct {
	_        map[string]string
	unExport map[string]string

	Map0 map[string]string
	Map1 map[string]*string
	Map2 map[*string]string
	Map3 map[*string]*string
	Map4 map[*any]*any
	Map5 **map[**string]**any
	Map6 map[interface{}]interface{}

	Map7 map[time.Time]time1.Time
	Map8 map[**time.Time]**time1.Time
	Map9 **map[**time.Time]**time1.Time

	Map10 map[ast.ObjKind]http.HandlerFunc
	Map11 **map[*ast.ObjKind]*http.HandlerFunc

	Map12 map[[2]int][2]any
}

type MapCrazy struct {
	Map0 map[string][1][2][3]any
	Map1 map[*map[*map[any]time.Time]int][2]any
	Map2 map[*[]entity.User]any
	Map3 **map[**map[**string]**string]**map[**any]time1.Time
	Map4 map[any][2]**int
	Map5 map[any][2]any
	Map6 map[*map[*ast.ObjKind]**string]**entity.User
	Map7 map[*entity.User]*entity.User
}
