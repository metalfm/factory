package tests_test

import (
	"context"
	"testing"

	"github.com/metalfm/factory/internal/tests"
	fc_tests "github.com/metalfm/factory/internal/tests/fc"

	"github.com/stretchr/testify/assert"
)

func TestObj(t *testing.T) {
	factory := fc_tests.NewFactoryObj(tests.Obj{
		Status0:    "",
		Status1:    nil,
		Status2:    nil,
		Status3:    nil,
		Struct0:    tests.ObjStruct{},
		Struct1:    nil,
		Struct2:    tests.ObjStruct{},
		Struct3:    tests.ObjStruct{},
		Interface0: nil,
		Sel0:       tests.Sel{},
	})

	assert.Equal(t, tests.Obj{}, factory.MustBuild())
	assert.Equal(t, tests.Obj{}, factory.MustBuildCtx(context.Background()))
}

func TestObjUnExport(t *testing.T) {
	factory := fc_tests.NewFactoryObjUnExported(tests.ObjUnExported{})

	assert.Equal(t, tests.ObjUnExported{}, factory.MustBuild())
	assert.Equal(t, tests.ObjUnExported{}, factory.MustBuildCtx(context.Background()))
}

func TestObjEmbed(t *testing.T) {
	factory := fc_tests.NewFactoryObjEmbed(tests.ObjEmbed{
		Status:       "",
		ObjStruct:    tests.ObjStruct{},
		ObjInterface: nil,
	})

	assert.Equal(t, tests.ObjEmbed{}, factory.MustBuild())
	assert.Equal(t, tests.ObjEmbed{}, factory.MustBuildCtx(context.Background()))
}

func TestObjEmbedPtr(t *testing.T) {
	factory := fc_tests.NewFactoryObjEmbedPtr(tests.ObjEmbedPtr{
		Status:    nil,
		ObjStruct: nil,
		Status0:   nil,
	})

	assert.Equal(t, tests.ObjEmbedPtr{}, factory.MustBuild())
	assert.Equal(t, tests.ObjEmbedPtr{}, factory.MustBuildCtx(context.Background()))
}

func TestObjExternalInterface(t *testing.T) {
	factory := fc_tests.NewFactoryObjExternalInterface(tests.ObjExternalInterface{
		HandlerFunc: nil,
		Handler0:    nil,
		Handler1:    nil,
		Handler2:    nil,
	})

	assert.Equal(t, tests.ObjExternalInterface{}, factory.MustBuild())
	assert.Equal(t, tests.ObjExternalInterface{}, factory.MustBuildCtx(context.Background()))
}
