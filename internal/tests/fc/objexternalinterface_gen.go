// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	fmt "fmt"
	tests "github.com/metalfm/factory/internal/tests"
	http "net/http"
	atomic "sync/atomic"
)

type factoryObjExternalInterface struct {
	entity         *tests.ObjExternalInterface
	seq            func(e *tests.ObjExternalInterface)
	seqInt         func(e *tests.ObjExternalInterface)
	seqInt64       func(e *tests.ObjExternalInterface)
	onCreate       func(ctx context.Context, e *tests.ObjExternalInterface) error
	onHandlerFunc  func(e *tests.ObjExternalInterface)
	setHandlerFunc bool
	onHandler0     func(e *tests.ObjExternalInterface)
	setHandler0    bool
	onHandler1     func(e *tests.ObjExternalInterface)
	setHandler1    bool
	onHandler2     func(e *tests.ObjExternalInterface)
	setHandler2    bool
}

func NewFactoryObjExternalInterface(e tests.ObjExternalInterface) *factoryObjExternalInterface {
	return &factoryObjExternalInterface{
		entity:        &e,
		seq:           func(e *tests.ObjExternalInterface) {},
		seqInt:        func(e *tests.ObjExternalInterface) {},
		seqInt64:      func(e *tests.ObjExternalInterface) {},
		onCreate:      func(ctx context.Context, e *tests.ObjExternalInterface) error { return nil },
		onHandlerFunc: func(e *tests.ObjExternalInterface) {},
		onHandler0:    func(e *tests.ObjExternalInterface) {},
		onHandler1:    func(e *tests.ObjExternalInterface) {},
		onHandler2:    func(e *tests.ObjExternalInterface) {},
	}
}

func (slf *factoryObjExternalInterface) HandlerFunc(v http.HandlerFunc) *factoryObjExternalInterface {
	slf.entity.HandlerFunc = v
	slf.setHandlerFunc = true
	return slf
}

func (slf *factoryObjExternalInterface) OnHandlerFunc(fn func(e *tests.ObjExternalInterface)) *factoryObjExternalInterface {
	slf.onHandlerFunc = fn
	return slf
}

func (slf *factoryObjExternalInterface) Handler0(v http.HandlerFunc) *factoryObjExternalInterface {
	slf.entity.Handler0 = v
	slf.setHandler0 = true
	return slf
}

func (slf *factoryObjExternalInterface) OnHandler0(fn func(e *tests.ObjExternalInterface)) *factoryObjExternalInterface {
	slf.onHandler0 = fn
	return slf
}

func (slf *factoryObjExternalInterface) Handler1(v *http.HandlerFunc) *factoryObjExternalInterface {
	slf.entity.Handler1 = v
	slf.setHandler1 = true
	return slf
}

func (slf *factoryObjExternalInterface) OnHandler1(fn func(e *tests.ObjExternalInterface)) *factoryObjExternalInterface {
	slf.onHandler1 = fn
	return slf
}

func (slf *factoryObjExternalInterface) Handler2(v **http.HandlerFunc) *factoryObjExternalInterface {
	slf.entity.Handler2 = v
	slf.setHandler2 = true
	return slf
}

func (slf *factoryObjExternalInterface) OnHandler2(fn func(e *tests.ObjExternalInterface)) *factoryObjExternalInterface {
	slf.onHandler2 = fn
	return slf
}

func (slf *factoryObjExternalInterface) Seq(fn func(e *tests.ObjExternalInterface)) *factoryObjExternalInterface {
	slf.seq = fn
	return slf
}

func (slf *factoryObjExternalInterface) SeqInt(fn func(e *tests.ObjExternalInterface, n int)) *factoryObjExternalInterface {
	var seq int32
	slf.seqInt = func(e *tests.ObjExternalInterface) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factoryObjExternalInterface) SeqInt64(fn func(e *tests.ObjExternalInterface, n int64)) *factoryObjExternalInterface {
	var seq int64
	slf.seqInt64 = func(e *tests.ObjExternalInterface) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factoryObjExternalInterface) OnCreate(fn func(ctx context.Context, e *tests.ObjExternalInterface) error) *factoryObjExternalInterface {
	slf.onCreate = fn
	return slf
}

func (slf *factoryObjExternalInterface) MustBuild() tests.ObjExternalInterface {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factoryObjExternalInterface) MustBuildCtx(ctx context.Context) tests.ObjExternalInterface {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setHandlerFunc {
		slf.onHandlerFunc(slf.entity)
	}

	if !slf.setHandler0 {
		slf.onHandler0(slf.entity)
	}

	if !slf.setHandler1 {
		slf.onHandler1(slf.entity)
	}

	if !slf.setHandler2 {
		slf.onHandler2(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}
