package tests

//go:generate go run ../../fcgen.go -type=Types,TypesEmpty,TypesPointer,TypesCrazy

//nolint:unused //need to test unexported fields
type Types struct {
	_        int
	unExport int

	Any any

	Bool   bool
	String string

	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Uintptr uintptr

	Byte byte
	Rune rune

	Float32 float32
	Float64 float64

	Complex64  complex64
	Complex128 complex128

	Err error

	bool
	Str0, str1, Str2, _ string
}

type TypesEmpty struct{}

//nolint:unused //need to test unexported fields
type TypesPointer struct {
	_        *int
	unExport *int

	Any    *any
	Bool   *bool
	String *string

	Str ******string

	Int   *int
	Int8  *int8
	Int16 *int16
	Int32 *int32
	Int64 *int64

	Uint    *uint
	Uint8   *uint8
	Uint16  *uint16
	Uint32  *uint32
	Uint64  *uint64
	Uintptr *uintptr

	Byte *byte
	Rune *rune

	Float32 *float32
	Float64 *float64

	Complex64  *complex64
	Complex128 *complex128

	Err *error

	bool
	Str0, str1, Str2, _ *string
}

type TypesCrazy struct {
	Int *****int
	Str *****string
}
