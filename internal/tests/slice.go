package tests

import (
	"time"
	time1 "time"
)

//go:generate go run ../../fcgen/fcgen.go -type=Slice,SliceSel,SliceFunc

//nolint:unused //need to test unexported fields
type Slice struct {
	_        []string
	unExport []string

	Slice0 []string
	Slice1 []int
	Slice2 []bool
	Slice3 []*string
	Slice4 []*int
	Slice5 []*bool
	Slice6 [][]**int
	Slice7 [][]int
	Slice8 **[]**[]**any
}

type SliceSel struct {
	Slice0 **[]**time.Time
	Slice1 [][]time1.Time

	Slice2 **[][]**time.Time
	Slice3 **[]**[]time1.Time
}

type SliceFunc struct {
	Slice0 **[]**func(a, b string) **any
	Slice1 **[]**func(string, string) []func() (**error, **any)
	Slice2 **[]func(a, b **[]**[]func(any) (c, d string)) **[]**[]**any
}
