// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	tests "factory/internal/tests"
	fmt "fmt"
	atomic "sync/atomic"
	time "time"
	time1 "time"
)

type factorySliceSel struct {
	entity    *tests.SliceSel
	seq       func(e *tests.SliceSel)
	seqInt    func(e *tests.SliceSel)
	seqInt64  func(e *tests.SliceSel)
	onCreate  func(ctx context.Context, e *tests.SliceSel) error
	onSlice0  func(e *tests.SliceSel)
	setSlice0 bool
	onSlice1  func(e *tests.SliceSel)
	setSlice1 bool
	onSlice2  func(e *tests.SliceSel)
	setSlice2 bool
	onSlice3  func(e *tests.SliceSel)
	setSlice3 bool
}

func NewFactorySliceSel(e tests.SliceSel) *factorySliceSel {
	return &factorySliceSel{
		entity:   &e,
		seq:      func(e *tests.SliceSel) {},
		seqInt:   func(e *tests.SliceSel) {},
		seqInt64: func(e *tests.SliceSel) {},
		onCreate: func(ctx context.Context, e *tests.SliceSel) error { return nil },
		onSlice0: func(e *tests.SliceSel) {},
		onSlice1: func(e *tests.SliceSel) {},
		onSlice2: func(e *tests.SliceSel) {},
		onSlice3: func(e *tests.SliceSel) {},
	}
}

func (slf *factorySliceSel) Slice0(v **[]**time.Time) *factorySliceSel {
	slf.entity.Slice0 = v
	slf.setSlice0 = true
	return slf
}

func (slf *factorySliceSel) OnSlice0(fn func(e *tests.SliceSel)) *factorySliceSel {
	slf.onSlice0 = fn
	return slf
}

func (slf *factorySliceSel) Slice1(v [][]time1.Time) *factorySliceSel {
	slf.entity.Slice1 = v
	slf.setSlice1 = true
	return slf
}

func (slf *factorySliceSel) OnSlice1(fn func(e *tests.SliceSel)) *factorySliceSel {
	slf.onSlice1 = fn
	return slf
}

func (slf *factorySliceSel) Slice2(v **[][]**time.Time) *factorySliceSel {
	slf.entity.Slice2 = v
	slf.setSlice2 = true
	return slf
}

func (slf *factorySliceSel) OnSlice2(fn func(e *tests.SliceSel)) *factorySliceSel {
	slf.onSlice2 = fn
	return slf
}

func (slf *factorySliceSel) Slice3(v **[]**[]time1.Time) *factorySliceSel {
	slf.entity.Slice3 = v
	slf.setSlice3 = true
	return slf
}

func (slf *factorySliceSel) OnSlice3(fn func(e *tests.SliceSel)) *factorySliceSel {
	slf.onSlice3 = fn
	return slf
}

func (slf *factorySliceSel) Seq(fn func(e *tests.SliceSel)) *factorySliceSel {
	slf.seq = fn
	return slf
}

func (slf *factorySliceSel) SeqInt(fn func(e *tests.SliceSel, n int)) *factorySliceSel {
	var seq int32
	slf.seqInt = func(e *tests.SliceSel) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factorySliceSel) SeqInt64(fn func(e *tests.SliceSel, n int64)) *factorySliceSel {
	var seq int64
	slf.seqInt64 = func(e *tests.SliceSel) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factorySliceSel) OnCreate(fn func(ctx context.Context, e *tests.SliceSel) error) *factorySliceSel {
	slf.onCreate = fn
	return slf
}

func (slf *factorySliceSel) MustBuild() tests.SliceSel {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factorySliceSel) MustBuildCtx(ctx context.Context) tests.SliceSel {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setSlice0 {
		slf.onSlice0(slf.entity)
	}

	if !slf.setSlice1 {
		slf.onSlice1(slf.entity)
	}

	if !slf.setSlice2 {
		slf.onSlice2(slf.entity)
	}

	if !slf.setSlice3 {
		slf.onSlice3(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}