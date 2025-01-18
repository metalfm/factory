package tests

import (
	"time"
	time1 "time"

	"github.com/metalfm/factory/internal/tests/entity"
)

//go:generate go run ../../fcgen.go -type=Func

//nolint:unused //need to test unexported fields
type Func struct {
	_        func()
	unExport func() error

	F0 func()
	F1 func(string, string) (int, int)
	F2 func(str1, str2 string) (k, v int, err error)
	F3 func(str1 string, str2 string) (k int, v int, err error)

	F4 func(*string, *string) (*int, *int)
	F5 func(str1, str2 *string) (k, v *int, err *error)
	F6 func(str1 *string, str2 *string) (k *int, v *int, err *error)

	F7, F8 **func(**int) (**time.Time, **time1.Time)
	F9     **func(a, b map[string]func(a, b time.Time, u **entity.User, c map[any]any)) [5]**int
}
