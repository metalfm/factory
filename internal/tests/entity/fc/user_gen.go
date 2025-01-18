// Code generated by fcgen; DO NOT EDIT.

package fc_entity

import (
	context "context"
	entity "factory/internal/tests/entity"
	fmt "fmt"
	atomic "sync/atomic"
)

type factoryUser struct {
	entity      *entity.User
	seq         func(e *entity.User)
	seqInt      func(e *entity.User)
	seqInt64    func(e *entity.User)
	onCreate    func(ctx context.Context, e *entity.User) error
	onID        func(e *entity.User)
	setID       bool
	onName      func(e *entity.User)
	setName     bool
	onLocation  func(e *entity.User)
	setLocation bool
	onFriend    func(e *entity.User)
	setFriend   bool
}

func NewFactoryUser(e entity.User) *factoryUser {
	return &factoryUser{
		entity:     &e,
		seq:        func(e *entity.User) {},
		seqInt:     func(e *entity.User) {},
		seqInt64:   func(e *entity.User) {},
		onCreate:   func(ctx context.Context, e *entity.User) error { return nil },
		onID:       func(e *entity.User) {},
		onName:     func(e *entity.User) {},
		onLocation: func(e *entity.User) {},
		onFriend:   func(e *entity.User) {},
	}
}

func (slf *factoryUser) ID(v int) *factoryUser {
	slf.entity.ID = v
	slf.setID = true
	return slf
}

func (slf *factoryUser) OnID(fn func(e *entity.User)) *factoryUser {
	slf.onID = fn
	return slf
}

func (slf *factoryUser) Name(v string) *factoryUser {
	slf.entity.Name = v
	slf.setName = true
	return slf
}

func (slf *factoryUser) OnName(fn func(e *entity.User)) *factoryUser {
	slf.onName = fn
	return slf
}

func (slf *factoryUser) Location(v string) *factoryUser {
	slf.entity.Location = v
	slf.setLocation = true
	return slf
}

func (slf *factoryUser) OnLocation(fn func(e *entity.User)) *factoryUser {
	slf.onLocation = fn
	return slf
}

func (slf *factoryUser) Friend(v *entity.User) *factoryUser {
	slf.entity.Friend = v
	slf.setFriend = true
	return slf
}

func (slf *factoryUser) OnFriend(fn func(e *entity.User)) *factoryUser {
	slf.onFriend = fn
	return slf
}

func (slf *factoryUser) Seq(fn func(e *entity.User)) *factoryUser {
	slf.seq = fn
	return slf
}

func (slf *factoryUser) SeqInt(fn func(e *entity.User, n int)) *factoryUser {
	var seq int32
	slf.seqInt = func(e *entity.User) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factoryUser) SeqInt64(fn func(e *entity.User, n int64)) *factoryUser {
	var seq int64
	slf.seqInt64 = func(e *entity.User) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factoryUser) OnCreate(fn func(ctx context.Context, e *entity.User) error) *factoryUser {
	slf.onCreate = fn
	return slf
}

func (slf *factoryUser) MustBuild() entity.User {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factoryUser) MustBuildCtx(ctx context.Context) entity.User {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setID {
		slf.onID(slf.entity)
	}

	if !slf.setName {
		slf.onName(slf.entity)
	}

	if !slf.setLocation {
		slf.onLocation(slf.entity)
	}

	if !slf.setFriend {
		slf.onFriend(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}