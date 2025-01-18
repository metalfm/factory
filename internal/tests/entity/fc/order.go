package fc_entity

import (
	"context"
	"math/rand"
	"time"

	"factory/internal/tests/entity"
)

var Order = NewFactoryOrder(
	entity.Order{
		Status: entity.StatusNew,
		Src:    "src",
		Dst:    "dst",
	}).
	SeqInt(func(e *entity.Order, n int) {
		e.ID = uint64(n)
	}).
	OnUUID(func(e *entity.Order) {
		e.UUID = rand.Int()
	}).
	OnCreatedAt(func(e *entity.Order) {
		e.CreatedAt = time.Now()
	}).
	OnCreate(func(ctx context.Context, e *entity.Order) error {
		// you can pass any storage instance in ctx
		// retrieve it, for example *sqlx.DB
		// and use it to store entity.Order

		return nil
	})
