package tests_test

import (
	"context"
	"testing"

	"factory/internal/tests"
	fc_tests "factory/internal/tests/fc"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	factory := fc_tests.NewFactoryMap(tests.Map{
		Map0:  nil,
		Map1:  nil,
		Map2:  nil,
		Map3:  nil,
		Map4:  nil,
		Map5:  nil,
		Map6:  nil,
		Map7:  nil,
		Map8:  nil,
		Map9:  nil,
		Map10: nil,
		Map11: nil,
		Map12: nil,
	})

	assert.Equal(t, tests.Map{}, factory.MustBuild())
	assert.Equal(t, tests.Map{}, factory.MustBuildCtx(context.Background()))
}

func TestMapCrazy(t *testing.T) {
	factory := fc_tests.NewFactoryMapCrazy(tests.MapCrazy{
		Map0: nil,
		Map1: nil,
		Map2: nil,
		Map3: nil,
		Map4: nil,
		Map5: nil,
		Map6: nil,
		Map7: nil,
	})

	assert.Equal(t, tests.MapCrazy{}, factory.MustBuild())
	assert.Equal(t, tests.MapCrazy{}, factory.MustBuildCtx(context.Background()))
}
