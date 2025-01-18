package tests_test

import (
	"context"
	"testing"

	"factory/internal/tests"
	fc_tests "factory/internal/tests/fc"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	factory := fc_tests.NewFactoryFunc(tests.Func{
		F0: nil,
		F1: nil,
		F2: nil,
		F3: nil,
		F4: nil,
		F5: nil,
		F6: nil,
		F7: nil,
		F8: nil,
		F9: nil,
	})

	assert.Equal(t, tests.Func{}, factory.MustBuild())
	assert.Equal(t, tests.Func{}, factory.MustBuildCtx(context.Background()))
}
