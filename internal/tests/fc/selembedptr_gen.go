// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	fmt "fmt"
	tests "github.com/metalfm/factory/internal/tests"
	atomic "sync/atomic"
	time "time"
)

type factorySelEmbedPtr struct {
	entity   *tests.SelEmbedPtr
	seq      func(e *tests.SelEmbedPtr)
	seqInt   func(e *tests.SelEmbedPtr)
	seqInt64 func(e *tests.SelEmbedPtr)
	onCreate func(ctx context.Context, e *tests.SelEmbedPtr) error
	onTime   func(e *tests.SelEmbedPtr)
	setTime  bool
	onTime0  func(e *tests.SelEmbedPtr)
	setTime0 bool
}

func NewFactorySelEmbedPtr(e tests.SelEmbedPtr) *factorySelEmbedPtr {
	return &factorySelEmbedPtr{
		entity:   &e,
		seq:      func(e *tests.SelEmbedPtr) {},
		seqInt:   func(e *tests.SelEmbedPtr) {},
		seqInt64: func(e *tests.SelEmbedPtr) {},
		onCreate: func(ctx context.Context, e *tests.SelEmbedPtr) error { return nil },
		onTime:   func(e *tests.SelEmbedPtr) {},
		onTime0:  func(e *tests.SelEmbedPtr) {},
	}
}

func (slf *factorySelEmbedPtr) Time(v *time.Time) *factorySelEmbedPtr {
	slf.entity.Time = v
	slf.setTime = true
	return slf
}

func (slf *factorySelEmbedPtr) OnTime(fn func(e *tests.SelEmbedPtr)) *factorySelEmbedPtr {
	slf.onTime = fn
	return slf
}

func (slf *factorySelEmbedPtr) Time0(v ****time.Time) *factorySelEmbedPtr {
	slf.entity.Time0 = v
	slf.setTime0 = true
	return slf
}

func (slf *factorySelEmbedPtr) OnTime0(fn func(e *tests.SelEmbedPtr)) *factorySelEmbedPtr {
	slf.onTime0 = fn
	return slf
}

func (slf *factorySelEmbedPtr) Seq(fn func(e *tests.SelEmbedPtr)) *factorySelEmbedPtr {
	slf.seq = fn
	return slf
}

func (slf *factorySelEmbedPtr) SeqInt(fn func(e *tests.SelEmbedPtr, n int)) *factorySelEmbedPtr {
	var seq int32
	slf.seqInt = func(e *tests.SelEmbedPtr) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factorySelEmbedPtr) SeqInt64(fn func(e *tests.SelEmbedPtr, n int64)) *factorySelEmbedPtr {
	var seq int64
	slf.seqInt64 = func(e *tests.SelEmbedPtr) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factorySelEmbedPtr) OnCreate(fn func(ctx context.Context, e *tests.SelEmbedPtr) error) *factorySelEmbedPtr {
	slf.onCreate = fn
	return slf
}

func (slf *factorySelEmbedPtr) MustBuild() tests.SelEmbedPtr {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factorySelEmbedPtr) MustBuildCtx(ctx context.Context) tests.SelEmbedPtr {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setTime {
		slf.onTime(slf.entity)
	}

	if !slf.setTime0 {
		slf.onTime0(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}
