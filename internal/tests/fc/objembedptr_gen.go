// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	tests "factory/internal/tests"
	fmt "fmt"
	atomic "sync/atomic"
)

type factoryObjEmbedPtr struct {
	entity       *tests.ObjEmbedPtr
	seq          func(e *tests.ObjEmbedPtr)
	seqInt       func(e *tests.ObjEmbedPtr)
	seqInt64     func(e *tests.ObjEmbedPtr)
	onCreate     func(ctx context.Context, e *tests.ObjEmbedPtr) error
	onStatus     func(e *tests.ObjEmbedPtr)
	setStatus    bool
	onObjStruct  func(e *tests.ObjEmbedPtr)
	setObjStruct bool
	onStatus0    func(e *tests.ObjEmbedPtr)
	setStatus0   bool
}

func NewFactoryObjEmbedPtr(e tests.ObjEmbedPtr) *factoryObjEmbedPtr {
	return &factoryObjEmbedPtr{
		entity:      &e,
		seq:         func(e *tests.ObjEmbedPtr) {},
		seqInt:      func(e *tests.ObjEmbedPtr) {},
		seqInt64:    func(e *tests.ObjEmbedPtr) {},
		onCreate:    func(ctx context.Context, e *tests.ObjEmbedPtr) error { return nil },
		onStatus:    func(e *tests.ObjEmbedPtr) {},
		onObjStruct: func(e *tests.ObjEmbedPtr) {},
		onStatus0:   func(e *tests.ObjEmbedPtr) {},
	}
}

func (slf *factoryObjEmbedPtr) Status(v *tests.Status) *factoryObjEmbedPtr {
	slf.entity.Status = v
	slf.setStatus = true
	return slf
}

func (slf *factoryObjEmbedPtr) OnStatus(fn func(e *tests.ObjEmbedPtr)) *factoryObjEmbedPtr {
	slf.onStatus = fn
	return slf
}

func (slf *factoryObjEmbedPtr) ObjStruct(v *tests.ObjStruct) *factoryObjEmbedPtr {
	slf.entity.ObjStruct = v
	slf.setObjStruct = true
	return slf
}

func (slf *factoryObjEmbedPtr) OnObjStruct(fn func(e *tests.ObjEmbedPtr)) *factoryObjEmbedPtr {
	slf.onObjStruct = fn
	return slf
}

func (slf *factoryObjEmbedPtr) Status0(v ******tests.Status) *factoryObjEmbedPtr {
	slf.entity.Status0 = v
	slf.setStatus0 = true
	return slf
}

func (slf *factoryObjEmbedPtr) OnStatus0(fn func(e *tests.ObjEmbedPtr)) *factoryObjEmbedPtr {
	slf.onStatus0 = fn
	return slf
}

func (slf *factoryObjEmbedPtr) Seq(fn func(e *tests.ObjEmbedPtr)) *factoryObjEmbedPtr {
	slf.seq = fn
	return slf
}

func (slf *factoryObjEmbedPtr) SeqInt(fn func(e *tests.ObjEmbedPtr, n int)) *factoryObjEmbedPtr {
	var seq int32
	slf.seqInt = func(e *tests.ObjEmbedPtr) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factoryObjEmbedPtr) SeqInt64(fn func(e *tests.ObjEmbedPtr, n int64)) *factoryObjEmbedPtr {
	var seq int64
	slf.seqInt64 = func(e *tests.ObjEmbedPtr) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factoryObjEmbedPtr) OnCreate(fn func(ctx context.Context, e *tests.ObjEmbedPtr) error) *factoryObjEmbedPtr {
	slf.onCreate = fn
	return slf
}

func (slf *factoryObjEmbedPtr) MustBuild() tests.ObjEmbedPtr {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factoryObjEmbedPtr) MustBuildCtx(ctx context.Context) tests.ObjEmbedPtr {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setStatus {
		slf.onStatus(slf.entity)
	}

	if !slf.setObjStruct {
		slf.onObjStruct(slf.entity)
	}

	if !slf.setStatus0 {
		slf.onStatus0(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}