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

type factoryChan struct {
	entity    *tests.Chan
	seq       func(e *tests.Chan)
	seqInt    func(e *tests.Chan)
	seqInt64  func(e *tests.Chan)
	onCreate  func(ctx context.Context, e *tests.Chan) error
	onChan0   func(e *tests.Chan)
	setChan0  bool
	onChan1   func(e *tests.Chan)
	setChan1  bool
	onChan2   func(e *tests.Chan)
	setChan2  bool
	onChan3   func(e *tests.Chan)
	setChan3  bool
	onChan4   func(e *tests.Chan)
	setChan4  bool
	onChan5   func(e *tests.Chan)
	setChan5  bool
	onChan6   func(e *tests.Chan)
	setChan6  bool
	onChan7   func(e *tests.Chan)
	setChan7  bool
	onChan8   func(e *tests.Chan)
	setChan8  bool
	onChan9   func(e *tests.Chan)
	setChan9  bool
	onChan10  func(e *tests.Chan)
	setChan10 bool
	onChan11  func(e *tests.Chan)
	setChan11 bool
	onChan12  func(e *tests.Chan)
	setChan12 bool
}

func NewFactoryChan(e tests.Chan) *factoryChan {
	return &factoryChan{
		entity:   &e,
		seq:      func(e *tests.Chan) {},
		seqInt:   func(e *tests.Chan) {},
		seqInt64: func(e *tests.Chan) {},
		onCreate: func(ctx context.Context, e *tests.Chan) error { return nil },
		onChan0:  func(e *tests.Chan) {},
		onChan1:  func(e *tests.Chan) {},
		onChan2:  func(e *tests.Chan) {},
		onChan3:  func(e *tests.Chan) {},
		onChan4:  func(e *tests.Chan) {},
		onChan5:  func(e *tests.Chan) {},
		onChan6:  func(e *tests.Chan) {},
		onChan7:  func(e *tests.Chan) {},
		onChan8:  func(e *tests.Chan) {},
		onChan9:  func(e *tests.Chan) {},
		onChan10: func(e *tests.Chan) {},
		onChan11: func(e *tests.Chan) {},
		onChan12: func(e *tests.Chan) {},
	}
}

func (slf *factoryChan) Chan0(v chan int) *factoryChan {
	slf.entity.Chan0 = v
	slf.setChan0 = true
	return slf
}

func (slf *factoryChan) OnChan0(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan0 = fn
	return slf
}

func (slf *factoryChan) Chan1(v chan<- int) *factoryChan {
	slf.entity.Chan1 = v
	slf.setChan1 = true
	return slf
}

func (slf *factoryChan) OnChan1(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan1 = fn
	return slf
}

func (slf *factoryChan) Chan2(v <-chan int) *factoryChan {
	slf.entity.Chan2 = v
	slf.setChan2 = true
	return slf
}

func (slf *factoryChan) OnChan2(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan2 = fn
	return slf
}

func (slf *factoryChan) Chan3(v chan *int) *factoryChan {
	slf.entity.Chan3 = v
	slf.setChan3 = true
	return slf
}

func (slf *factoryChan) OnChan3(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan3 = fn
	return slf
}

func (slf *factoryChan) Chan4(v chan<- *any) *factoryChan {
	slf.entity.Chan4 = v
	slf.setChan4 = true
	return slf
}

func (slf *factoryChan) OnChan4(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan4 = fn
	return slf
}

func (slf *factoryChan) Chan5(v <-chan *any) *factoryChan {
	slf.entity.Chan5 = v
	slf.setChan5 = true
	return slf
}

func (slf *factoryChan) OnChan5(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan5 = fn
	return slf
}

func (slf *factoryChan) Chan6(v chan time.Time) *factoryChan {
	slf.entity.Chan6 = v
	slf.setChan6 = true
	return slf
}

func (slf *factoryChan) OnChan6(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan6 = fn
	return slf
}

func (slf *factoryChan) Chan7(v chan *time1.Time) *factoryChan {
	slf.entity.Chan7 = v
	slf.setChan7 = true
	return slf
}

func (slf *factoryChan) OnChan7(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan7 = fn
	return slf
}

func (slf *factoryChan) Chan8(v **chan **int) *factoryChan {
	slf.entity.Chan8 = v
	slf.setChan8 = true
	return slf
}

func (slf *factoryChan) OnChan8(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan8 = fn
	return slf
}

func (slf *factoryChan) Chan9(v **<-chan **int) *factoryChan {
	slf.entity.Chan9 = v
	slf.setChan9 = true
	return slf
}

func (slf *factoryChan) OnChan9(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan9 = fn
	return slf
}

func (slf *factoryChan) Chan10(v **chan<- **int) *factoryChan {
	slf.entity.Chan10 = v
	slf.setChan10 = true
	return slf
}

func (slf *factoryChan) OnChan10(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan10 = fn
	return slf
}

func (slf *factoryChan) Chan11(v **<-chan **[5][]func(a, b **any) (c, d **string)) *factoryChan {
	slf.entity.Chan11 = v
	slf.setChan11 = true
	return slf
}

func (slf *factoryChan) OnChan11(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan11 = fn
	return slf
}

func (slf *factoryChan) Chan12(v **<-chan **[5][]func(a, b **any) (c, d **string)) *factoryChan {
	slf.entity.Chan12 = v
	slf.setChan12 = true
	return slf
}

func (slf *factoryChan) OnChan12(fn func(e *tests.Chan)) *factoryChan {
	slf.onChan12 = fn
	return slf
}

func (slf *factoryChan) Seq(fn func(e *tests.Chan)) *factoryChan {
	slf.seq = fn
	return slf
}

func (slf *factoryChan) SeqInt(fn func(e *tests.Chan, n int)) *factoryChan {
	var seq int32
	slf.seqInt = func(e *tests.Chan) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factoryChan) SeqInt64(fn func(e *tests.Chan, n int64)) *factoryChan {
	var seq int64
	slf.seqInt64 = func(e *tests.Chan) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factoryChan) OnCreate(fn func(ctx context.Context, e *tests.Chan) error) *factoryChan {
	slf.onCreate = fn
	return slf
}

func (slf *factoryChan) MustBuild() tests.Chan {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factoryChan) MustBuildCtx(ctx context.Context) tests.Chan {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setChan0 {
		slf.onChan0(slf.entity)
	}

	if !slf.setChan1 {
		slf.onChan1(slf.entity)
	}

	if !slf.setChan2 {
		slf.onChan2(slf.entity)
	}

	if !slf.setChan3 {
		slf.onChan3(slf.entity)
	}

	if !slf.setChan4 {
		slf.onChan4(slf.entity)
	}

	if !slf.setChan5 {
		slf.onChan5(slf.entity)
	}

	if !slf.setChan6 {
		slf.onChan6(slf.entity)
	}

	if !slf.setChan7 {
		slf.onChan7(slf.entity)
	}

	if !slf.setChan8 {
		slf.onChan8(slf.entity)
	}

	if !slf.setChan9 {
		slf.onChan9(slf.entity)
	}

	if !slf.setChan10 {
		slf.onChan10(slf.entity)
	}

	if !slf.setChan11 {
		slf.onChan11(slf.entity)
	}

	if !slf.setChan12 {
		slf.onChan12(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}
