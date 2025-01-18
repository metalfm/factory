package tests_test

import (
	"context"
	"testing"

	"factory/internal/tests"
	fc_tests "factory/internal/tests/fc"

	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	factory := fc_tests.NewFactoryInterface(tests.Interface{
		Field0: nil,
		Field1: nil,
		Field2: nil,
		Field3: nil,
		Field4: nil,
		Field5: nil,
	})

	assert.Equal(t, tests.Interface{}, factory.MustBuild())
	assert.Equal(t, tests.Interface{}, factory.MustBuildCtx(context.Background()))
}

func TestInterfaceEmpty(t *testing.T) {
	factory := fc_tests.NewFactoryInterfaceEmpty(tests.InterfaceEmpty{})

	assert.Equal(t, tests.InterfaceEmpty{}, factory.MustBuild())
	assert.Equal(t, tests.InterfaceEmpty{}, factory.MustBuildCtx(context.Background()))
}
