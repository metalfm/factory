package entity

import (
	"time"
)

type Payloader interface {
	PayloadInvoice | PayloadCreditNote | PayloadReceipt
}

type Document[T Payloader] struct {
	ID        uint64
	Payload   T
	CreatedAt time.Time
}

type PayloadInvoice struct {
	InvField0 string
	InvField1 *string
}

type PayloadCreditNote struct {
	CnField0 uint
}

type PayloadReceipt struct {
	RecField0 []*int
}
