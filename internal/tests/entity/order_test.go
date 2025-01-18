package entity_test

import (
	"fmt"
	"math/rand"
	"testing"

	fc_entity "factory/internal/tests/entity/fc"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	order := fc_entity.Order.MustBuild()

	fmt.Printf("%+v", order)
	// Output:
	// {ID:1 UUID:3095472579637049805 Status:new Src:src Dst:dst User:<nil> CreatedAt:2025-01-17 15:20:57.248584 +0200 EET m=+0.000842043 UpdateAt:<nil>}
}

func TestOrderWithTheSameUUIDs(t *testing.T) {
	u := rand.Int()

	o0 := fc_entity.Order.UUID(u).MustBuild()
	o1 := fc_entity.Order.UUID(u).MustBuild()

	assert.Equal(t, o0.UUID, o1.UUID)
	assert.NotEqual(t, o0.CreatedAt, o1.CreatedAt)
}

func TestOrderWithUser(t *testing.T) {
	u := fc_entity.User.MustBuild()
	o := fc_entity.Order.User(&u).MustBuild()

	fmt.Printf("%+v", o)
	// Output:
	// {ID:1 UUID:3161475588708285381 Status:new Src:src Dst:dst User:0x1400011b080 CreatedAt:2025-01-17 15:33:01.688783 +0200 EET m=+0.000995543 UpdateAt:<nil>}
}
