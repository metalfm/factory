package fc_entity

import (
	"factory/internal/tests/entity"
)

var User = NewFactoryUser(
	entity.User{
		Location: "Cyprus",
	}).
	SeqInt(func(e *entity.User, n int) {
		e.ID = n
	}).
	OnFriend(func(e *entity.User) {
		friend := NewFactoryUser(
			entity.User{
				Name: "Alice", // everyone wants to be friends with Alice
			}).
			MustBuild()

		e.Friend = &friend
	})
