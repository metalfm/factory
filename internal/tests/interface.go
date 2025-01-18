package tests

//go:generate go run ../../fcgen.go -type=Interface,InterfaceEmpty

//nolint:unused //need to test unexported fields
type empty interface{}
type Empty interface{}

type Interface struct {
	Field0 interface{}
	Field1 interface{ Method() }
	Field2 *interface{}
	Field3 *interface{ Method() }

	Field4 *interface{}
	Field5 **interface{}

	_        interface{}
	unExport interface{} //nolint:unused //need to test unexported fields
}

type InterfaceEmpty struct {
	empty //nolint:unused //need to test unexported fields
	Empty
}
