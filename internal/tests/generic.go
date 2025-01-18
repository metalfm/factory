package tests

import (
	"factory/internal/tests/entity"
)

//go:generate go run ../../fcgen.go -type=Generic

type Gen[t any] interface {
	Generate(t) t
}

type Payloader[E any] interface{}
type payloader interface{}

type Generic[
	T0,
	T1 any,
	E entity.Document[P],
	P entity.Payloader,
	G Gen[entity.Document[P]],
] struct {
	Field0    T1
	Document0 E
	Document1 entity.Document[P]
	Document2 entity.Document[entity.PayloadCreditNote]

	P0 payloader
	P1 Payloader[any]
	P2 Payloader[**any]
	P3 Payloader[entity.PayloadCreditNote]
	P4 Payloader[T0]
	P5 Payloader[P]

	G0 G
}
