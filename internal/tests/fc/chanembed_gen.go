// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	tests "factory/internal/tests"
	fmt "fmt"
	atomic "sync/atomic"
)

type factoryChanEmbed struct {
	entity     *tests.ChanEmbed
	seq        func(e *tests.ChanEmbed)
	seqInt     func(e *tests.ChanEmbed)
	seqInt64   func(e *tests.ChanEmbed)
	onCreate   func(ctx context.Context, e *tests.ChanEmbed) error
	onChanMsg  func(e *tests.ChanEmbed)
	setChanMsg bool
	onChan1    func(e *tests.ChanEmbed)
	setChan1   bool
}

func NewFactoryChanEmbed(e tests.ChanEmbed) *factoryChanEmbed {
	return &factoryChanEmbed{
		entity:    &e,
		seq:       func(e *tests.ChanEmbed) {},
		seqInt:    func(e *tests.ChanEmbed) {},
		seqInt64:  func(e *tests.ChanEmbed) {},
		onCreate:  func(ctx context.Context, e *tests.ChanEmbed) error { return nil },
		onChanMsg: func(e *tests.ChanEmbed) {},
		onChan1:   func(e *tests.ChanEmbed) {},
	}
}

func (slf *factoryChanEmbed) ChanMsg(v tests.ChanMsg) *factoryChanEmbed {
	slf.entity.ChanMsg = v
	slf.setChanMsg = true
	return slf
}

func (slf *factoryChanEmbed) OnChanMsg(fn func(e *tests.ChanEmbed)) *factoryChanEmbed {
	slf.onChanMsg = fn
	return slf
}

func (slf *factoryChanEmbed) Chan1(v tests.ChanMsg) *factoryChanEmbed {
	slf.entity.Chan1 = v
	slf.setChan1 = true
	return slf
}

func (slf *factoryChanEmbed) OnChan1(fn func(e *tests.ChanEmbed)) *factoryChanEmbed {
	slf.onChan1 = fn
	return slf
}

func (slf *factoryChanEmbed) Seq(fn func(e *tests.ChanEmbed)) *factoryChanEmbed {
	slf.seq = fn
	return slf
}

func (slf *factoryChanEmbed) SeqInt(fn func(e *tests.ChanEmbed, n int)) *factoryChanEmbed {
	var seq int32
	slf.seqInt = func(e *tests.ChanEmbed) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factoryChanEmbed) SeqInt64(fn func(e *tests.ChanEmbed, n int64)) *factoryChanEmbed {
	var seq int64
	slf.seqInt64 = func(e *tests.ChanEmbed) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factoryChanEmbed) OnCreate(fn func(ctx context.Context, e *tests.ChanEmbed) error) *factoryChanEmbed {
	slf.onCreate = fn
	return slf
}

func (slf *factoryChanEmbed) MustBuild() tests.ChanEmbed {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factoryChanEmbed) MustBuildCtx(ctx context.Context) tests.ChanEmbed {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setChanMsg {
		slf.onChanMsg(slf.entity)
	}

	if !slf.setChan1 {
		slf.onChan1(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}
