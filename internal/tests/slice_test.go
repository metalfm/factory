package tests_test

import (
	"context"
	"testing"

	"github.com/metalfm/factory/internal/tests"
	fc_tests "github.com/metalfm/factory/internal/tests/fc"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	factory := fc_tests.NewFactorySlice(tests.Slice{
		Slice0: nil,
		Slice1: nil,
		Slice2: nil,
		Slice3: nil,
		Slice4: nil,
		Slice5: nil,
		Slice6: nil,
		Slice7: nil,
		Slice8: nil,
	})

	assert.Equal(t, tests.Slice{}, factory.MustBuild())
	assert.Equal(t, tests.Slice{}, factory.MustBuildCtx(context.Background()))
}

func TestSliceSel(t *testing.T) {
	factory := fc_tests.NewFactorySliceSel(tests.SliceSel{
		Slice0: nil,
		Slice1: nil,
		Slice2: nil,
		Slice3: nil,
	})

	assert.Equal(t, tests.SliceSel{}, factory.MustBuild())
	assert.Equal(t, tests.SliceSel{}, factory.MustBuildCtx(context.Background()))
}

func TestSliceFunc(t *testing.T) {
	factory := fc_tests.NewFactorySliceFunc(tests.SliceFunc{
		Slice0: nil,
		Slice1: nil,
		Slice2: nil,
	})

	assert.Equal(t, tests.SliceFunc{}, factory.MustBuild())
	assert.Equal(t, tests.SliceFunc{}, factory.MustBuildCtx(context.Background()))
}
