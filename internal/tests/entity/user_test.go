package entity_test

import (
	"fmt"
	"testing"

	fc_entity "github.com/metalfm/factory/internal/tests/entity/fc"
)

func TestUser(t *testing.T) {
	u := fc_entity.User.MustBuild()

	fmt.Printf("%+v", u)
	// Output:
	// {ID:1 Name: Location:Cyprus Friend:0x1400011b020}
}

func TestWithoutFriendship(t *testing.T) {
	u := fc_entity.User.Friend(nil).MustBuild()

	fmt.Printf("%+v", u)
	// Output:
	// {ID:1 Name: Location:Cyprus Friend:<nil>}
}

func TestMyFriendIsBob(t *testing.T) {
	bob := fc_entity.User.
		Name("Bob").
		Friend(nil).
		MustBuild()

	me := fc_entity.User.
		Name("me").
		Friend(&bob).
		MustBuild()

	fmt.Printf("%+v %+v", me, *me.Friend)
	// Output:
	// {ID:2 Name:me Location:Cyprus Friend:0x14000109080} {ID:1 Name:Bob Location:Cyprus Friend:<nil>}
}
