package tests

import (
	"go/ast"
	"net/http"
	"time"
	time1 "time"

	"github.com/metalfm/factory/internal/tests/entity"
)

//go:generate go run ../../fcgen/fcgen.go -type=Sel,SelUnExport,SelEmbedAlias,SelEmbed,SelEmbedPtr

type Sel struct {
	Time0 time.Time
	Time1 time1.Time

	Time2 *time.Time
	Time3 *time1.Time

	Kind0 ast.ObjKind
	Kind1 *ast.ObjKind

	Handler0 http.HandlerFunc
	Handler1 *http.HandlerFunc

	User0 entity.User
	User1 *entity.User
	User2 **entity.User
}

//nolint:unused //need to test unexported fields
type SelUnExport struct {
	_        time.Time
	unExport time1.Time
}

type SelEmbedAlias struct {
	_ time1.Time
	time1.Time
}

type SelEmbed struct {
	time.Time
}

type SelEmbedPtr struct {
	*time.Time
	Time0 ****time.Time
}
