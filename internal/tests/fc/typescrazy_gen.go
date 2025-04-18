// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	fmt "fmt"
	tests "github.com/metalfm/factory/internal/tests"
	atomic "sync/atomic"
)

type factoryTypesCrazy struct {
	entity   *tests.TypesCrazy
	seq      func(e *tests.TypesCrazy)
	seqInt   func(e *tests.TypesCrazy)
	seqInt64 func(e *tests.TypesCrazy)
	onCreate func(ctx context.Context, e *tests.TypesCrazy) error
	onInt    func(e *tests.TypesCrazy)
	setInt   bool
	onStr    func(e *tests.TypesCrazy)
	setStr   bool
}

func NewFactoryTypesCrazy(e tests.TypesCrazy) *factoryTypesCrazy {
	return &factoryTypesCrazy{
		entity:   &e,
		seq:      func(e *tests.TypesCrazy) {},
		seqInt:   func(e *tests.TypesCrazy) {},
		seqInt64: func(e *tests.TypesCrazy) {},
		onCreate: func(ctx context.Context, e *tests.TypesCrazy) error { return nil },
		onInt:    func(e *tests.TypesCrazy) {},
		onStr:    func(e *tests.TypesCrazy) {},
	}
}

func (slf *factoryTypesCrazy) Int(v *****int) *factoryTypesCrazy {
	slf.entity.Int = v
	slf.setInt = true
	return slf
}

func (slf *factoryTypesCrazy) OnInt(fn func(e *tests.TypesCrazy)) *factoryTypesCrazy {
	slf.onInt = fn
	return slf
}

func (slf *factoryTypesCrazy) Str(v *****string) *factoryTypesCrazy {
	slf.entity.Str = v
	slf.setStr = true
	return slf
}

func (slf *factoryTypesCrazy) OnStr(fn func(e *tests.TypesCrazy)) *factoryTypesCrazy {
	slf.onStr = fn
	return slf
}

func (slf *factoryTypesCrazy) Seq(fn func(e *tests.TypesCrazy)) *factoryTypesCrazy {
	slf.seq = fn
	return slf
}

func (slf *factoryTypesCrazy) SeqInt(fn func(e *tests.TypesCrazy, n int)) *factoryTypesCrazy {
	var seq int32
	slf.seqInt = func(e *tests.TypesCrazy) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factoryTypesCrazy) SeqInt64(fn func(e *tests.TypesCrazy, n int64)) *factoryTypesCrazy {
	var seq int64
	slf.seqInt64 = func(e *tests.TypesCrazy) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factoryTypesCrazy) OnCreate(fn func(ctx context.Context, e *tests.TypesCrazy) error) *factoryTypesCrazy {
	slf.onCreate = fn
	return slf
}

func (slf *factoryTypesCrazy) MustBuild() tests.TypesCrazy {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factoryTypesCrazy) MustBuildCtx(ctx context.Context) tests.TypesCrazy {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setInt {
		slf.onInt(slf.entity)
	}

	if !slf.setStr {
		slf.onStr(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}
