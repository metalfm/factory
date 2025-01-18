package tests_test

import (
	"context"
	"testing"
	time1 "time"

	"github.com/metalfm/factory/internal/tests"
	fc_tests "github.com/metalfm/factory/internal/tests/fc"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	factory := fc_tests.NewFactoryArray(tests.Array{
		Array0:  [8]int{},
		Array1:  [127]uint{},
		Array2:  [97]string{},
		Array3:  [127]bool{},
		Array5:  [1]any{},
		Array6:  [8]*int{},
		Array7:  [127]*uint{},
		Array8:  [97]*string{},
		Array9:  [127]*bool{},
		Array10: [1]*any{},
		Array11: [1][2]any{},
	})

	assert.Equal(t, tests.Array{}, factory.MustBuild())
	assert.Equal(t, tests.Array{}, factory.MustBuildCtx(context.Background()))
}

func TestArrayPtr(t *testing.T) {
	factory := fc_tests.NewFactoryArrayPtr(tests.ArrayPtr{
		Array0: nil,
		Array1: nil,
		Array2: nil,
		Array3: nil,
		Array4: nil,
		Array5: nil,
		Array6: nil,
		Array7: nil,
		Array8: nil,
	})

	assert.Equal(t, tests.ArrayPtr{}, factory.MustBuild())
	assert.Equal(t, tests.ArrayPtr{}, factory.MustBuildCtx(context.Background()))
}

func TestArraySel(t *testing.T) {
	factory := fc_tests.NewFactoryArraySel(tests.ArraySel{
		Array0: [5]time1.Time{},
		Array1: nil,
		Array2: nil,
		Array3: [5]**time1.Time{},
		Array4: [1][2]**time1.Time{},
	})

	assert.Equal(t, tests.ArraySel{}, factory.MustBuild())
	assert.Equal(t, tests.ArraySel{}, factory.MustBuildCtx(context.Background()))
}

func TestArrayObj(t *testing.T) {
	factory := fc_tests.NewFactoryArrayObj(tests.ArrayObj{
		Array0: [5]tests.ArrayStruct{},
		Array2: nil,
		Array3: [5]**tests.ArrayStruct{},
		Array4: nil,
		Array5: nil,
		Array6: nil,
	})

	assert.Equal(t, tests.ArrayObj{}, factory.MustBuild())
	assert.Equal(t, tests.ArrayObj{}, factory.MustBuildCtx(context.Background()))
}

func TestArrayEmptyInterface(t *testing.T) {
	factory := fc_tests.NewFactoryArrayEmptyInterface(tests.ArrayEmptyInterface{
		Array0: [5]interface{}{},
		Array1: nil,
		Array2: nil,
		Array3: [5]*interface{}{},
		Array4: nil,
		Array5: nil,
	})

	assert.Equal(t, tests.ArrayEmptyInterface{}, factory.MustBuild())
	assert.Equal(t, tests.ArrayEmptyInterface{}, factory.MustBuildCtx(context.Background()))
}
