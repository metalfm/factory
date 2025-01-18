package tests_test

import (
	"context"
	"testing"

	"factory/internal/tests"
	"factory/internal/tests/entity"
	fc_tests "factory/internal/tests/fc"
)

func TestGeneric(t *testing.T) {
	factory := fc_tests.NewFactoryGeneric(tests.Generic[
		any,
		any,
		entity.Document[entity.PayloadInvoice],
		entity.PayloadInvoice,
		tests.Gen[entity.Document[entity.PayloadInvoice]],
	]{
		Field0:    nil,
		Document0: entity.Document[entity.PayloadInvoice]{},
		Document1: entity.Document[entity.PayloadInvoice]{},
		Document2: entity.Document[entity.PayloadCreditNote]{},
		P0:        nil,
		P1:        nil,
		P2:        nil,
		P3:        nil,
		P4:        nil,
		P5:        nil,
		G0:        nil,
	})

	factory.MustBuild()
	factory.MustBuildCtx(context.Background())
}
