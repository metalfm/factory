// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	tests "factory/internal/tests"
	fmt "fmt"
	atomic "sync/atomic"
)

type factorySelUnExport struct {
	entity   *tests.SelUnExport
	seq      func(e *tests.SelUnExport)
	seqInt   func(e *tests.SelUnExport)
	seqInt64 func(e *tests.SelUnExport)
	onCreate func(ctx context.Context, e *tests.SelUnExport) error
}

func NewFactorySelUnExport(e tests.SelUnExport) *factorySelUnExport {
	return &factorySelUnExport{
		entity:   &e,
		seq:      func(e *tests.SelUnExport) {},
		seqInt:   func(e *tests.SelUnExport) {},
		seqInt64: func(e *tests.SelUnExport) {},
		onCreate: func(ctx context.Context, e *tests.SelUnExport) error { return nil },
	}
}

func (slf *factorySelUnExport) Seq(fn func(e *tests.SelUnExport)) *factorySelUnExport {
	slf.seq = fn
	return slf
}

func (slf *factorySelUnExport) SeqInt(fn func(e *tests.SelUnExport, n int)) *factorySelUnExport {
	var seq int32
	slf.seqInt = func(e *tests.SelUnExport) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factorySelUnExport) SeqInt64(fn func(e *tests.SelUnExport, n int64)) *factorySelUnExport {
	var seq int64
	slf.seqInt64 = func(e *tests.SelUnExport) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factorySelUnExport) OnCreate(fn func(ctx context.Context, e *tests.SelUnExport) error) *factorySelUnExport {
	slf.onCreate = fn
	return slf
}

func (slf *factorySelUnExport) MustBuild() tests.SelUnExport {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factorySelUnExport) MustBuildCtx(ctx context.Context) tests.SelUnExport {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}
