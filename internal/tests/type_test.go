package tests_test

import (
	"context"
	"errors"
	"testing"

	"factory/internal/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fc_tests "factory/internal/tests/fc"
)

func TestTypes(t *testing.T) {
	e := expected()

	got := fc_tests.NewFactoryTypes(e).MustBuild()
	require.Equal(t, e, got)

	got = fc_tests.NewFactoryTypes(e).MustBuildCtx(context.Background())
	require.Equal(t, e, got)
}

func TestTypesSetters(t *testing.T) {
	factory := fc_tests.NewFactoryTypes(tests.Types{}).
		Any("any").
		Bool(true).
		String("str").
		Int(1).
		Int8(2).
		Int16(3).
		Int32(4).
		Int64(5).
		Uint(6).
		Uint8(7).
		Uint16(8).
		Uint32(9).
		Uint64(10).
		Uintptr(11).
		Byte(12).
		Rune(13).
		Float32(14).
		Float64(15).
		Complex64(16).
		Complex128(17).
		Str0("str0").
		Str2("str2").
		Err(errors.New("my error"))

	assert.Equal(t, expected(), factory.MustBuild())
	assert.Equal(t, expected(), factory.MustBuildCtx(context.Background()))
}

func TestTypesSettersFunc(t *testing.T) {
	e := tests.Types{
		Any:        "any_2",
		Bool:       false,
		String:     "str_2",
		Int:        3,
		Int8:       3,
		Int16:      3,
		Int32:      3,
		Int64:      3,
		Uint:       3,
		Uint8:      3,
		Uint16:     3,
		Uint32:     3,
		Uint64:     3,
		Uintptr:    3,
		Byte:       3,
		Rune:       3,
		Float32:    3.0,
		Float64:    3.0,
		Complex64:  3.0,
		Complex128: 3.0,
		Err:        errors.New("err2"),
		Str0:       "str_2",
		Str2:       "str_2",
	}

	got := fc_tests.
		NewFactoryTypes(expected()).
		Any("any_2").
		OnAny(func(e *tests.Types) { e.Any = "any_1" }).
		Bool(false).
		OnBool(func(e *tests.Types) { e.Bool = true }).
		String("str_2").
		OnString(func(e *tests.Types) { e.String = "str_1" }).
		Int(3).
		OnInt(func(e *tests.Types) { e.Int = 2 }).
		Int8(3).
		OnInt8(func(e *tests.Types) { e.Int8 = 2 }).
		Int16(3).
		OnInt16(func(e *tests.Types) { e.Int16 = 2 }).
		Int32(3).
		OnInt32(func(e *tests.Types) { e.Int32 = 2 }).
		Int64(3).
		OnInt64(func(e *tests.Types) { e.Int64 = 2 }).
		Uint(3).
		OnUint(func(e *tests.Types) { e.Uint = 2 }).
		Uint8(3).
		OnUint8(func(e *tests.Types) { e.Uint8 = 2 }).
		Uint16(3).
		OnUint16(func(e *tests.Types) { e.Uint16 = 2 }).
		Uint32(3).
		OnUint32(func(e *tests.Types) { e.Uint32 = 2 }).
		Uint64(3).
		OnUint64(func(e *tests.Types) { e.Uint64 = 2 }).
		Uintptr(3).
		OnUintptr(func(e *tests.Types) { e.Uintptr = 2 }).
		Byte(3).
		OnByte(func(e *tests.Types) { e.Byte = 2 }).
		Rune(3).
		OnRune(func(e *tests.Types) { e.Rune = 2 }).
		Float32(3).
		OnFloat32(func(e *tests.Types) { e.Float32 = 2 }).
		Float64(3).
		OnFloat64(func(e *tests.Types) { e.Float64 = 2 }).
		Complex64(3).
		OnComplex64(func(e *tests.Types) { e.Complex64 = 2 }).
		Complex128(3).
		OnComplex128(func(e *tests.Types) { e.Complex128 = 2 }).
		Err(errors.New("err2")).
		OnErr(func(e *tests.Types) { e.Err = errors.New("err1") }).
		Str0("str_2").
		OnStr0(func(e *tests.Types) { e.Str0 = "str_0" }).
		Str2("str_2").
		OnStr2(func(e *tests.Types) { e.Str2 = "str_0" }).
		MustBuild()

	assert.Equal(t, e, got)
}

func TestSequence(t *testing.T) {
	factory := fc_tests.
		NewFactoryTypes(tests.Types{}).
		Seq(func(e *tests.Types) {
			e.Byte = 1
		}).
		SeqInt(func(e *tests.Types, n int) {
			e.Int = n
		}).
		SeqInt64(func(e *tests.Types, n int64) {
			e.Int64 = n
		})

	e := tests.Types{
		Byte:  1,
		Int:   1,
		Int64: 1,
	}
	assert.Equal(t, e, factory.MustBuild())

	e = tests.Types{
		Byte:  1,
		Int:   2,
		Int64: 2,
	}
	assert.Equal(t, e, factory.MustBuild())
}

func TestOnCreatePanic(t *testing.T) {
	factory := fc_tests.
		NewFactoryTypes(tests.Types{}).
		OnCreate(func(_ context.Context, _ *tests.Types) error {
			return errors.New("my error")
		})

	assert.PanicsWithError(t, "factory.OnCreate: my error", func() { factory.MustBuild() })
}

func TestTypesEmpty(t *testing.T) {
	factory := fc_tests.NewFactoryTypesEmpty(tests.TypesEmpty{})

	assert.Equal(t, tests.TypesEmpty{}, factory.MustBuild())
	assert.Equal(t, tests.TypesEmpty{}, factory.MustBuildCtx(context.Background()))
}

func TestPointer(t *testing.T) {
	e := expectedPtr()

	got := fc_tests.NewFactoryTypesPointer(e).MustBuild()
	require.Equal(t, e, got)

	got = fc_tests.NewFactoryTypesPointer(e).MustBuildCtx(context.Background())
	require.Equal(t, e, got)
}

func TestPointerSetters(t *testing.T) {
	factory := fc_tests.NewFactoryTypesPointer(tests.TypesPointer{}).
		Any(ptr(any("any"))).
		Bool(ptr(true)).
		String(ptr("str")).
		Int(ptr(1)).
		Int8(ptr(int8(2))).
		Int16(ptr(int16(3))).
		Int32(ptr(int32(4))).
		Int64(ptr(int64(5))).
		Uint(ptr(uint(6))).
		Uint8(ptr(uint8(7))).
		Uint16(ptr(uint16(8))).
		Uint32(ptr(uint32(9))).
		Uint64(ptr(uint64(10))).
		Uintptr(ptr(uintptr(11))).
		Byte(ptr(byte(12))).
		Rune(ptr(rune(13))).
		Float32(ptr(float32(14))).
		Float64(ptr(float64(15))).
		Complex64(ptr(complex64(16))).
		Complex128(ptr(complex128(17))).
		Err(ptr(errors.New("error"))).
		Str0(ptr("str_0")).
		Str2(ptr("str_2"))

	assert.Equal(t, expectedPtr(), factory.MustBuild())
	assert.Equal(t, expectedPtr(), factory.MustBuildCtx(context.Background()))
}

func TestTypesCrazy(t *testing.T) {
	factory := fc_tests.NewFactoryTypesCrazy(tests.TypesCrazy{
		Int: nil,
		Str: nil,
	})

	assert.Equal(t, tests.TypesCrazy{}, factory.MustBuild())
	assert.Equal(t, tests.TypesCrazy{}, factory.MustBuildCtx(context.Background()))
}

func expectedPtr() tests.TypesPointer {
	return tests.TypesPointer{
		Any:        ptr(any("any")),
		Bool:       ptr(true),
		String:     ptr("str"),
		Int:        ptr(1),
		Int8:       ptr(int8(2)),
		Int16:      ptr(int16(3)),
		Int32:      ptr(int32(4)),
		Int64:      ptr(int64(5)),
		Uint:       ptr(uint(6)),
		Uint8:      ptr(uint8(7)),
		Uint16:     ptr(uint16(8)),
		Uint32:     ptr(uint32(9)),
		Uint64:     ptr(uint64(10)),
		Uintptr:    ptr(uintptr(11)),
		Byte:       ptr(byte(12)),
		Rune:       ptr(rune(13)),
		Float32:    ptr(float32(14)),
		Float64:    ptr(float64(15)),
		Complex64:  ptr(complex64(16)),
		Complex128: ptr(complex128(17)),
		Err:        ptr(errors.New("error")),
		Str0:       ptr("str_0"),
		Str2:       ptr("str_2"),
	}
}

func expected() tests.Types {
	return tests.Types{
		Any:        "any",
		Bool:       true,
		String:     "str",
		Int:        1,
		Int8:       2,
		Int16:      3,
		Int32:      4,
		Int64:      5,
		Uint:       6,
		Uint8:      7,
		Uint16:     8,
		Uint32:     9,
		Uint64:     10,
		Uintptr:    11,
		Byte:       12,
		Rune:       13,
		Float32:    14,
		Float64:    15,
		Complex64:  16,
		Complex128: 17,
		Err:        errors.New("my error"),
		Str0:       "str0",
		Str2:       "str2",
	}
}

func ptr[T any](v T) *T {
	return &v
}
