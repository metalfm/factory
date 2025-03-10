// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	fmt "fmt"
	tests "github.com/metalfm/factory/internal/tests"
	atomic "sync/atomic"
)

type factoryInterfaceEmpty struct {
	entity   *tests.InterfaceEmpty
	seq      func(e *tests.InterfaceEmpty)
	seqInt   func(e *tests.InterfaceEmpty)
	seqInt64 func(e *tests.InterfaceEmpty)
	onCreate func(ctx context.Context, e *tests.InterfaceEmpty) error
	onEmpty  func(e *tests.InterfaceEmpty)
	setEmpty bool
}

func NewFactoryInterfaceEmpty(e tests.InterfaceEmpty) *factoryInterfaceEmpty {
	return &factoryInterfaceEmpty{
		entity:   &e,
		seq:      func(e *tests.InterfaceEmpty) {},
		seqInt:   func(e *tests.InterfaceEmpty) {},
		seqInt64: func(e *tests.InterfaceEmpty) {},
		onCreate: func(ctx context.Context, e *tests.InterfaceEmpty) error { return nil },
		onEmpty:  func(e *tests.InterfaceEmpty) {},
	}
}

func (slf *factoryInterfaceEmpty) Empty(v tests.Empty) *factoryInterfaceEmpty {
	slf.entity.Empty = v
	slf.setEmpty = true
	return slf
}

func (slf *factoryInterfaceEmpty) OnEmpty(fn func(e *tests.InterfaceEmpty)) *factoryInterfaceEmpty {
	slf.onEmpty = fn
	return slf
}

func (slf *factoryInterfaceEmpty) Seq(fn func(e *tests.InterfaceEmpty)) *factoryInterfaceEmpty {
	slf.seq = fn
	return slf
}

func (slf *factoryInterfaceEmpty) SeqInt(fn func(e *tests.InterfaceEmpty, n int)) *factoryInterfaceEmpty {
	var seq int32
	slf.seqInt = func(e *tests.InterfaceEmpty) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factoryInterfaceEmpty) SeqInt64(fn func(e *tests.InterfaceEmpty, n int64)) *factoryInterfaceEmpty {
	var seq int64
	slf.seqInt64 = func(e *tests.InterfaceEmpty) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factoryInterfaceEmpty) OnCreate(fn func(ctx context.Context, e *tests.InterfaceEmpty) error) *factoryInterfaceEmpty {
	slf.onCreate = fn
	return slf
}

func (slf *factoryInterfaceEmpty) MustBuild() tests.InterfaceEmpty {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factoryInterfaceEmpty) MustBuildCtx(ctx context.Context) tests.InterfaceEmpty {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setEmpty {
		slf.onEmpty(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}
