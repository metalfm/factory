package tests

import (
	"net/http"
)

//go:generate go run ../../fcgen.go -type=Obj,ObjUnExported,ObjEmbed,ObjEmbedPtr,ObjExternalInterface

type Status string
type status string

type ObjStruct struct{}
type objStruct struct{}

type ObjInterface interface{}
type objInterface interface{}

type Obj struct {
	Status0 Status
	Status1 *Status

	Status2, Status3 *Status

	Struct0 ObjStruct
	Struct1 *ObjStruct

	Struct2, Struct3, struct4 ObjStruct //nolint:unused //need to test unexported fields

	Interface0 ObjInterface
	Sel0       Sel
}

//nolint:unused //need to test unexported fields
type ObjUnExported struct {
	_ Status
	_ status

	unExport0 Status
	unExport1 status

	Struct0    objStruct
	Interface0 objInterface
}

//nolint:unused //need to test unexported fields
type ObjEmbed struct {
	status
	Status

	objStruct
	ObjStruct

	objInterface
	ObjInterface
}

type ObjEmbedPtr struct {
	*Status
	*ObjStruct

	Status0 ******Status
}

type ObjExternalInterface struct {
	http.HandlerFunc

	Handler0 http.HandlerFunc
	Handler1 *http.HandlerFunc
	Handler2 **http.HandlerFunc
}
