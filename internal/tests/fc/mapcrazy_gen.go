// Code generated by fcgen; DO NOT EDIT.

package fc_tests

import (
	context "context"
	fmt "fmt"
	tests "github.com/metalfm/factory/internal/tests"
	entity "github.com/metalfm/factory/internal/tests/entity"
	ast "go/ast"
	atomic "sync/atomic"
	time "time"
	time1 "time"
)

type factoryMapCrazy struct {
	entity   *tests.MapCrazy
	seq      func(e *tests.MapCrazy)
	seqInt   func(e *tests.MapCrazy)
	seqInt64 func(e *tests.MapCrazy)
	onCreate func(ctx context.Context, e *tests.MapCrazy) error
	onMap0   func(e *tests.MapCrazy)
	setMap0  bool
	onMap1   func(e *tests.MapCrazy)
	setMap1  bool
	onMap2   func(e *tests.MapCrazy)
	setMap2  bool
	onMap3   func(e *tests.MapCrazy)
	setMap3  bool
	onMap4   func(e *tests.MapCrazy)
	setMap4  bool
	onMap5   func(e *tests.MapCrazy)
	setMap5  bool
	onMap6   func(e *tests.MapCrazy)
	setMap6  bool
	onMap7   func(e *tests.MapCrazy)
	setMap7  bool
}

func NewFactoryMapCrazy(e tests.MapCrazy) *factoryMapCrazy {
	return &factoryMapCrazy{
		entity:   &e,
		seq:      func(e *tests.MapCrazy) {},
		seqInt:   func(e *tests.MapCrazy) {},
		seqInt64: func(e *tests.MapCrazy) {},
		onCreate: func(ctx context.Context, e *tests.MapCrazy) error { return nil },
		onMap0:   func(e *tests.MapCrazy) {},
		onMap1:   func(e *tests.MapCrazy) {},
		onMap2:   func(e *tests.MapCrazy) {},
		onMap3:   func(e *tests.MapCrazy) {},
		onMap4:   func(e *tests.MapCrazy) {},
		onMap5:   func(e *tests.MapCrazy) {},
		onMap6:   func(e *tests.MapCrazy) {},
		onMap7:   func(e *tests.MapCrazy) {},
	}
}

func (slf *factoryMapCrazy) Map0(v map[string][1][2][3]any) *factoryMapCrazy {
	slf.entity.Map0 = v
	slf.setMap0 = true
	return slf
}

func (slf *factoryMapCrazy) OnMap0(fn func(e *tests.MapCrazy)) *factoryMapCrazy {
	slf.onMap0 = fn
	return slf
}

func (slf *factoryMapCrazy) Map1(v map[*map[*map[any]time.Time]int][2]any) *factoryMapCrazy {
	slf.entity.Map1 = v
	slf.setMap1 = true
	return slf
}

func (slf *factoryMapCrazy) OnMap1(fn func(e *tests.MapCrazy)) *factoryMapCrazy {
	slf.onMap1 = fn
	return slf
}

func (slf *factoryMapCrazy) Map2(v map[*[]entity.User]any) *factoryMapCrazy {
	slf.entity.Map2 = v
	slf.setMap2 = true
	return slf
}

func (slf *factoryMapCrazy) OnMap2(fn func(e *tests.MapCrazy)) *factoryMapCrazy {
	slf.onMap2 = fn
	return slf
}

func (slf *factoryMapCrazy) Map3(v **map[**map[**string]**string]**map[**any]time1.Time) *factoryMapCrazy {
	slf.entity.Map3 = v
	slf.setMap3 = true
	return slf
}

func (slf *factoryMapCrazy) OnMap3(fn func(e *tests.MapCrazy)) *factoryMapCrazy {
	slf.onMap3 = fn
	return slf
}

func (slf *factoryMapCrazy) Map4(v map[any][2]**int) *factoryMapCrazy {
	slf.entity.Map4 = v
	slf.setMap4 = true
	return slf
}

func (slf *factoryMapCrazy) OnMap4(fn func(e *tests.MapCrazy)) *factoryMapCrazy {
	slf.onMap4 = fn
	return slf
}

func (slf *factoryMapCrazy) Map5(v map[any][2]any) *factoryMapCrazy {
	slf.entity.Map5 = v
	slf.setMap5 = true
	return slf
}

func (slf *factoryMapCrazy) OnMap5(fn func(e *tests.MapCrazy)) *factoryMapCrazy {
	slf.onMap5 = fn
	return slf
}

func (slf *factoryMapCrazy) Map6(v map[*map[*ast.ObjKind]**string]**entity.User) *factoryMapCrazy {
	slf.entity.Map6 = v
	slf.setMap6 = true
	return slf
}

func (slf *factoryMapCrazy) OnMap6(fn func(e *tests.MapCrazy)) *factoryMapCrazy {
	slf.onMap6 = fn
	return slf
}

func (slf *factoryMapCrazy) Map7(v map[*entity.User]*entity.User) *factoryMapCrazy {
	slf.entity.Map7 = v
	slf.setMap7 = true
	return slf
}

func (slf *factoryMapCrazy) OnMap7(fn func(e *tests.MapCrazy)) *factoryMapCrazy {
	slf.onMap7 = fn
	return slf
}

func (slf *factoryMapCrazy) Seq(fn func(e *tests.MapCrazy)) *factoryMapCrazy {
	slf.seq = fn
	return slf
}

func (slf *factoryMapCrazy) SeqInt(fn func(e *tests.MapCrazy, n int)) *factoryMapCrazy {
	var seq int32
	slf.seqInt = func(e *tests.MapCrazy) {
		v := atomic.AddInt32(&seq, 1)
		fn(e, int(v))
	}
	return slf
}

func (slf *factoryMapCrazy) SeqInt64(fn func(e *tests.MapCrazy, n int64)) *factoryMapCrazy {
	var seq int64
	slf.seqInt64 = func(e *tests.MapCrazy) {
		v := atomic.AddInt64(&seq, 1)
		fn(e, v)
	}
	return slf
}

func (slf *factoryMapCrazy) OnCreate(fn func(ctx context.Context, e *tests.MapCrazy) error) *factoryMapCrazy {
	slf.onCreate = fn
	return slf
}

func (slf *factoryMapCrazy) MustBuild() tests.MapCrazy {
	return slf.MustBuildCtx(context.Background())
}

func (slf *factoryMapCrazy) MustBuildCtx(ctx context.Context) tests.MapCrazy {
	slf.seq(slf.entity)
	slf.seqInt(slf.entity)
	slf.seqInt64(slf.entity)

	if !slf.setMap0 {
		slf.onMap0(slf.entity)
	}

	if !slf.setMap1 {
		slf.onMap1(slf.entity)
	}

	if !slf.setMap2 {
		slf.onMap2(slf.entity)
	}

	if !slf.setMap3 {
		slf.onMap3(slf.entity)
	}

	if !slf.setMap4 {
		slf.onMap4(slf.entity)
	}

	if !slf.setMap5 {
		slf.onMap5(slf.entity)
	}

	if !slf.setMap6 {
		slf.onMap6(slf.entity)
	}

	if !slf.setMap7 {
		slf.onMap7(slf.entity)
	}

	err := slf.onCreate(ctx, slf.entity)
	if err != nil {
		panic(fmt.Errorf("factory.OnCreate: %w", err))
	}

	return *slf.entity
}
