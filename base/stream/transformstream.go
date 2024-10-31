package stream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
)

type TransformStream struct {
	baseobject.BaseObject
	start, transform, flush js.Func
}

var singletont sync.Once

var transformstreaminterface js.Value

// GetRInterface get the JS interface ReadableStream.
func GetTInterface() js.Value {

	singletont.Do(func() {

		var err error
		if transformstreaminterface, err = baseobject.Get(js.Global(), "TransformStream"); err != nil {
			transformstreaminterface = js.Undefined()
		}
		baseobject.Register(transformstreaminterface, func(v js.Value) (interface{}, error) {
			return NewTransformStreamFromJSObject(v)
		})
	})

	return transformstreaminterface
}

// NewReadableStream Create a new ReadableStream
func NewTransformStream(
	start func(controller TransformStreamDefaultController),
	transform func(chunk interface{}, controller TransformStreamDefaultController),
	flush func(controller TransformStreamDefaultController),
) (TransformStream, error) {
	var t TransformStream
	var obj js.Value
	var err error
	if ri := GetTInterface(); !ri.IsUndefined() {
		ob, _ := baseobject.NewEmptyObject()
		t.start = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if c, err := baseobject.Discover(args[0]); err == nil {
				if ctrl, ok := c.(TransformStreamDefaultController); ok && start != nil {
					start(ctrl)
				}
			}
			return nil
		})
		t.transform = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if c, err := baseobject.Discover(args[1]); err == nil {
				if ctrl, ok := c.(TransformStreamDefaultController); ok && transform != nil {
					chunk, _ := baseobject.Discover(args[0])
					transform(chunk, ctrl)

				}
			}

			return nil
		})

		t.flush = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if c, err := baseobject.Discover(args[0]); err == nil {
				if ctrl, ok := c.(TransformStreamDefaultController); ok && flush != nil {
					flush(ctrl)

				}
			}
			return nil
		})

		ob.Set("start", t.start)
		ob.Set("transform", t.transform)
		ob.Set("flush", t.flush)

		if obj, err = baseobject.New(ri, ob.JSObject()); err == nil {
			t.BaseObject = t.BaseObject.SetObject(obj)

		}

	} else {
		err = ErrNotImplementedReadableStream
	}
	return t, err
}

func NewTransformStreamFromJSObject(obj js.Value) (TransformStream, error) {
	var t TransformStream
	var err error
	if tsi := GetTInterface(); !tsi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(tsi) {
				t.BaseObject = t.SetObject(obj)

			} else {
				err = ErrNotATransformStream
			}
		}
	} else {
		err = ErrNotImplementedTransformStream
	}

	return t, err
}

func (t *TransformStream) Release() {

	t.start.Release()
	t.transform.Release()
	t.flush.Release()
}

func TransfertToTransformStream(b baseobject.BaseObject) TransformStream {
	var t TransformStream
	t.BaseObject = t.SetObject(b.JSObject())
	return t
}
