package tests_test

import (
	"context"
	"testing"
	"time"
	time1 "time"

	"github.com/metalfm/factory/internal/tests"
	"github.com/metalfm/factory/internal/tests/entity"
	fc_tests "github.com/metalfm/factory/internal/tests/fc"

	"github.com/stretchr/testify/assert"
)

func TestStd(t *testing.T) {
	factory := fc_tests.NewFactorySel(tests.Sel{
		Time0:    time1.Time{},
		Time1:    time1.Time{},
		Time2:    nil,
		Time3:    nil,
		Kind0:    0,
		Kind1:    nil,
		Handler0: nil,
		Handler1: nil,
		User0:    entity.User{},
		User1:    nil,
		User2:    nil,
	})

	assert.Equal(t, tests.Sel{}, factory.MustBuild())
	assert.Equal(t, tests.Sel{}, factory.MustBuildCtx(context.Background()))
}

func TestStdUnExport(t *testing.T) {
	factory := fc_tests.NewFactorySelUnExport(tests.SelUnExport{})

	assert.Equal(t, tests.SelUnExport{}, factory.MustBuild())
	assert.Equal(t, tests.SelUnExport{}, factory.MustBuildCtx(context.Background()))
}

func TestStdEmbedAlias(t *testing.T) {
	factory := fc_tests.NewFactorySelEmbedAlias(tests.SelEmbedAlias{
		Time: time1.Time{},
	})

	assert.Equal(t, tests.SelEmbedAlias{}, factory.MustBuild())
	assert.Equal(t, tests.SelEmbedAlias{}, factory.MustBuildCtx(context.Background()))
}

func TestStdEmbed(t *testing.T) {
	factory := fc_tests.NewFactorySelEmbed(tests.SelEmbed{
		Time: time.Time{},
	})

	assert.Equal(t, tests.SelEmbed{}, factory.MustBuild())
	assert.Equal(t, tests.SelEmbed{}, factory.MustBuildCtx(context.Background()))
}

func TestStdEmbedPtr(t *testing.T) {
	factory := fc_tests.NewFactorySelEmbedPtr(tests.SelEmbedPtr{
		Time:  nil,
		Time0: nil,
	})

	assert.Equal(t, tests.SelEmbedPtr{}, factory.MustBuild())
	assert.Equal(t, tests.SelEmbedPtr{}, factory.MustBuildCtx(context.Background()))
}
