package tests_test

import (
	"context"
	"testing"

	"factory/internal/tests"
	fc_tests "factory/internal/tests/fc"

	"github.com/stretchr/testify/assert"
)

func TestChan(t *testing.T) {
	factory := fc_tests.NewFactoryChan(tests.Chan{
		Chan0:  nil,
		Chan1:  nil,
		Chan2:  nil,
		Chan3:  nil,
		Chan4:  nil,
		Chan5:  nil,
		Chan6:  nil,
		Chan7:  nil,
		Chan8:  nil,
		Chan9:  nil,
		Chan10: nil,
		Chan11: nil,
		Chan12: nil,
	})

	assert.Equal(t, tests.Chan{}, factory.MustBuild())
	assert.Equal(t, tests.Chan{}, factory.MustBuildCtx(context.Background()))
}

func TestChanUnExport(t *testing.T) {
	factory := fc_tests.NewFactoryChanUnExport(tests.ChanUnExport{
		Chan0: nil,
		Chan1: nil,
	})

	assert.Equal(t, tests.ChanUnExport{}, factory.MustBuild())
	assert.Equal(t, tests.ChanUnExport{}, factory.MustBuildCtx(context.Background()))
}

func TestChanEmbed(t *testing.T) {
	factory := fc_tests.NewFactoryChanEmbed(tests.ChanEmbed{
		ChanMsg: nil,
	})

	assert.Equal(t, tests.ChanEmbed{}, factory.MustBuild())
	assert.Equal(t, tests.ChanEmbed{}, factory.MustBuildCtx(context.Background()))
}
