package fc_entity

import (
	"time"

	"factory/internal/tests/entity"
)

var DocumentCN = NewFactoryDocument[entity.PayloadCreditNote](
	entity.Document[entity.PayloadCreditNote]{
		Payload: entity.PayloadCreditNote{
			CnField0: uint(1),
		},
	}).
	SeqInt(func(e *entity.Document[entity.PayloadCreditNote], n int) {
		e.ID = uint64(n)
	}).
	OnCreatedAt(func(e *entity.Document[entity.PayloadCreditNote]) {
		e.CreatedAt = time.Now()
	})

var DocumentInvoice = NewFactoryDocument[entity.PayloadInvoice](
	entity.Document[entity.PayloadInvoice]{
		Payload: entity.PayloadInvoice{
			InvField0: "field_0",
			InvField1: (func() *string {
				str := ""
				return &str
			})(),
		},
	}).
	SeqInt(func(e *entity.Document[entity.PayloadInvoice], n int) {
		e.ID = uint64(n)
	}).
	OnCreatedAt(func(e *entity.Document[entity.PayloadInvoice]) {
		e.CreatedAt = time.Now()
	})

func NewDocument[T entity.Payloader](e entity.Document[T]) *factoryDocument[T] {
	return NewFactoryDocument[T](e)
}
