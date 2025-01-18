// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	tests "factory/internal/tests"
	fmt "fmt"
	atomic "sync/atomic"
)

type factoryObj struct {
	entity        *tests.Obj
	seq           func(e *tests.Obj)
	seqInt        func(e *tests.Obj)
	seqInt64      func(e *tests.Obj)
	onCreate      func(ctx context.Context, e *tests.Obj) error
	onStatus0     func(e *tests.Obj)
	setStatus0    bool
	onStatus1     func(e *tests.Obj)
	setStatus1    bool
	onStatus2     func(e *tests.Obj)
	setStatus2    bool
	onStatus3     func(e *tests.Obj)
	setStatus3    bool
	onStruct0     func(e *tests.Obj)
	setStruct0    bool
	onStruct1     func(e *tests.Obj)
	setStruct1    bool
	onStruct2     func(e *tests.Obj)
	setStruct2    bool
	onStruct3     func(e *tests.Obj)
	setStruct3    bool
	onInterface0  func(e *tests.Obj)
	setInterface0 bool
	onSel0        func(e *tests.Obj)
	setSel0       bool
}

func NewFactoryObj(e tests.Obj) *factoryObj {
	return &factoryObj{
		entity:       &e,
		seq:          func(e *tests.Obj) {},
		seqInt:       func(e *tests.Obj) {},
		seqInt64:     func(e *tests.Obj) {},
		onCreate:     func(ctx context.Context, e *tests.Obj) error { return nil },
		onStatus0:    func(e *tests.Obj) {},
		onStatus1:    func(e *tests.Obj) {},
		onStatus2:    func(e *tests.Obj) {},
		onStatus3:    func(e *tests.Obj) {},
		onStruct0:    func(e *tests.Obj) {},
		onStruct1:    func(e *tests.Obj) {},
		onStruct2:    func(e *tests.Obj) {},
		onStruct3:    func(e *tests.Obj) {},
		onInterface0: func(e *tests.Obj) {},
		onSel0:       func(e *tests.Obj) {},
	}
}

func (slf *factoryObj) Status0(v tests.Status) *factoryObj {
	slf.entity.Status0 = v
	slf.setStatus0 = true
	return slf
}

func (slf *factoryObj) OnStatus0(fn func(e *tests.Obj)) *factoryObj {
	slf.onStatus0 = fn
	return slf
}

func (slf *factoryObj) Status1(v *tests.Status) *factoryObj {
	slf.entity.Status1 = v
	slf.setStatus1 = true
	return slf
}

func (slf *factoryObj) OnStatus1(fn func(e *tests.Obj)) *factoryObj {
	slf.onStatus1 = fn
	return slf
}

func (slf *factoryObj) Status2(v *tests.Status) *factoryObj {
	slf.entity.Status2 = v
	slf.setStatus2 = true
	return slf
}

func (slf *factoryObj) OnStatus2(fn func(e *tests.Obj)) *factoryObj {
	slf.onStatus2 = fn
	return slf
}

func (slf *factoryObj) Status3(v *tests.Status) *factoryObj {
	slf.entity.Status3 = v
	slf.setStatus3 = true
	return slf
}

func (slf *factoryObj) OnStatus3(fn func(e *tests.Obj)) *factoryObj {
	slf.onStatus3 = fn
	return slf
}

func (slf *factoryObj) Struct0(v tests.ObjStruct) *factoryObj {
	slf.entity.Struct0 = v
	slf.setStruct0 = true
	return slf
}

func (slf *factoryObj) OnStruct0(fn func(e *tests.Obj)) *factoryObj {
	slf.onStruct0 = fn
	return slf
}

func (slf *factoryObj) Struct1(v *tests.ObjStruct) *factoryObj {
	slf.entity.Struct1 = v
	slf.setStruct1 = true
	return slf
}

func (slf *factoryObj) OnStruct1(fn func(e *tests.Obj)) *factoryObj {
	slf.onStruct1 = fn
	return slf
}

func (slf *factoryObj) Struct2(v tests.ObjStruct) *factoryObj {
	slf.entity.Struct2 = v
	slf.setStruct2 = true
	return slf
}

func (slf *factoryObj) OnStruct2(fn func(e *tests.Obj)) *factoryObj {
	slf.onStruct2 = fn
	return slf
}

func (slf *factoryObj) Struct3(v tests.ObjStruct) *factoryObj {
	slf.entity.Struct3 = v
	slf.setStruct3 = true
	return slf
}

func (slf *factoryObj) OnStruct3(fn func(e *tests.Obj)) *factoryObj {
	slf.onStruct3 = fn
	return slf
}

func (slf *factoryObj) Interface0(v tests.ObjInterface) *factoryObj {
	slf.entity.Interface0 = v
	slf.setInterface0 = true
	return slf
}

func (slf *factoryObj) OnInterface0(fn func(e *tests.Obj)) *factoryObj {
	slf.onInterface0 = fn
	return slf
}

func (slf *factoryObj) Sel0(v tests.Sel) *factoryObj {
	slf.entity.Sel0 = v
	slf.setSel0 = true
	return slf
}

func (slf *factoryObj) OnSel0(fn func(e *tests.Obj)) *factoryObj {
	slf.onSel0 = fn
	return slf
}

func (slf *factoryObj) Seq(fn func(e *tests.Obj)) *factoryObj {
	slf.seq = fn
	return slf
}

func (slf *factoryObj) SeqInt(fn func(e *tests.Obj, n int)) *factoryObj {
	var seq int32
	slf.seqInt = func(e *tests.Obj) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factoryObj) SeqInt64(fn func(e *tests.Obj, n int64)) *factoryObj {
	var seq int64
	slf.seqInt64 = func(e *tests.Obj) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factoryObj) OnCreate(fn func(ctx context.Context, e *tests.Obj) error) *factoryObj {
	slf.onCreate = fn
	return slf
}

func (slf *factoryObj) MustBuild() tests.Obj {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factoryObj) MustBuildCtx(ctx context.Context) tests.Obj {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setStatus0 {
		slf.onStatus0(slf.entity)
	}

	if !slf.setStatus1 {
		slf.onStatus1(slf.entity)
	}

	if !slf.setStatus2 {
		slf.onStatus2(slf.entity)
	}

	if !slf.setStatus3 {
		slf.onStatus3(slf.entity)
	}

	if !slf.setStruct0 {
		slf.onStruct0(slf.entity)
	}

	if !slf.setStruct1 {
		slf.onStruct1(slf.entity)
	}

	if !slf.setStruct2 {
		slf.onStruct2(slf.entity)
	}

	if !slf.setStruct3 {
		slf.onStruct3(slf.entity)
	}

	if !slf.setInterface0 {
		slf.onInterface0(slf.entity)
	}

	if !slf.setSel0 {
		slf.onSel0(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}
