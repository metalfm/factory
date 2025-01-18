package tests

import (
	"time"
	time1 "time"
)

//go:generate go run ../../fcgen/fcgen.go -type=Array,ArrayPtr,ArraySel,ArrayObj,ArrayEmptyInterface

type ArrayStruct struct{}
type arrayStruct struct{}

//nolint:unused //need to test unexported fields
type Array struct {
	_        [5]int64
	unExport [5]int32

	Array0 [8.0]int
	Array1 [0x7f]uint
	Array2 ['a']string
	Array3 ['\x7f']bool
	Array5 [1]any

	Array6  [8.0]*int
	Array7  [0x7f]*uint
	Array8  ['a']*string
	Array9  ['\x7f']*bool
	Array10 [1]*any

	Array11 [1][2]any
	Array12 [1][2]map[string]string
	Array13 [1][2]*map[**string]**string
	Array14 [1][2]*map[**time1.Time]**time.Time
}

//nolint:unused //need to test unexported fields
type ArrayPtr struct {
	_        *[5]int64
	unExport *[5]int32

	Array0 *[8.0]int
	Array1 *[0x7f]uint
	Array2 **['a']string
	Array3 **['\x7f']bool

	Array4 *[8.0]*int
	Array5 *[0x7f]*uint
	Array6 **['a']**string
	Array7 **['\x7f']**bool

	Array8 **[1][2]**any
	Array9 **[1][2]func(a, b any) map[any]**func()
}

type ArraySel struct {
	Array0 [5]time.Time
	Array1 **[5]time.Time
	Array2 **[5]**time.Time
	Array3 [5]**time1.Time
	Array4 [1][2]**time1.Time
}

type ArrayObj struct {
	Array0 [5]ArrayStruct
	Array1 [5]arrayStruct
	Array2 **[5]ArrayStruct
	Array3 [5]**ArrayStruct
	Array4 **[5]**ArrayStruct
	Array5 **[1][2]**ArrayStruct
	Array6 **[1][2][3]**ArrayStruct
}

type ArrayEmptyInterface struct {
	Array0 [5]interface{}
	Array1 *[5]interface{}
	Array2 *[5]*interface{}
	Array3 [5]*interface{}
	Array4 **[5]**interface{}
	Array5 **[1][2]**interface{}
}
