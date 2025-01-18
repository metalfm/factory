package entity_test

import (
	"fmt"
	"testing"

	"factory/internal/tests/entity"
	fc_entity "factory/internal/tests/entity/fc"
)

func TestCreditNote(t *testing.T) {
	factory := fc_entity.DocumentCN

	fmt.Printf("%+v", factory.MustBuild())
	// Output:
	// {ID:1 Payload:{CnField0:1} CreatedAt:2025-01-17 15:41:50.523847 +0200 EET m=+0.000986792}

	factory.Payload(entity.PayloadCreditNote{CnField0: uint(3)})

	fmt.Printf("%+v", factory.MustBuild())
	// Output:
	// {ID:2 Payload:{CnField0:3} CreatedAt:2025-01-17 15:41:50.524283 +0200 EET m=+0.001423292}
}

func TestInvoice(t *testing.T) {
	factory := fc_entity.DocumentInvoice

	fmt.Printf("%+v", factory.MustBuild())
	// Output:
	// {ID:1 Payload:{InvField0:field_0 InvField1:0x140001040a0} CreatedAt:2025-01-17 15:42:36.495835 +0200 EET m=+0.000896501}

	factory.Payload(entity.PayloadInvoice{
		InvField0: "my_field_value",
	})

	fmt.Printf("%+v", factory.MustBuild())
	// Output:
	// {ID:2 Payload:{InvField0:my_field_value InvField1:<nil>} CreatedAt:2025-01-17 15:42:50.466779 +0200 EET m=+0.001241251}
}

func TestDocumentAny(t *testing.T) {
	d0 := entity.Document[entity.PayloadReceipt]{}
	f0 := fc_entity.NewDocument[entity.PayloadReceipt](d0)

	d1 := entity.Document[entity.PayloadInvoice]{}
	f1 := fc_entity.NewDocument[entity.PayloadInvoice](d1)

	fmt.Printf("%+v", f0.MustBuild())
	// Output:
	// {ID:0 Payload:{RecField0:[]} CreatedAt:0001-01-01 00:00:00 +0000 UTC}

	fmt.Printf("%+v", f1.MustBuild())
	// Output:
	// {ID:0 Payload:{InvField0: InvField1:<nil>} CreatedAt:0001-01-01 00:00:00 +0000 UTC}
}
