package stream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/typedarray"
)

var singletonWritableStreamDefault sync.Once

var writeablestreamdefaultinterface js.Value

// GetWritableStreamDefaultWriterInterface
func GetWritableStreamDefaultWriterInterface() js.Value {
	singletonReadableStreamDefault.Do(func() {

		var err error
		if writeablestreamdefaultinterface, err = baseobject.Get(js.Global(), "WritableStreamDefaultWriter"); err != nil {
			writeablestreamdefaultinterface = js.Undefined()
		}
	})

	return writeablestreamdefaultinterface
}

type WritableStreamDefaultWriter struct {
	baseobject.BaseObject
}

type WritableStreamDefaultWriterFrom interface {
	WritableStreamDefaultWriter_() WritableStreamDefaultWriter
}

func (w WritableStreamDefaultWriter) WritableStreamDefaultWriter_() WritableStreamDefaultWriter {
	return w
}

func NewWritableStreamDefaultWriterFromJSObject(obj js.Value) (WritableStreamDefaultWriter, error) {
	var w WritableStreamDefaultWriter

	if rsi := GetWritableStreamDefaultWriterInterface(); !rsi.IsUndefined() {
		if obj.InstanceOf(rsi) {
			w.BaseObject = w.SetObject(obj)
			return w, nil

		}
	}

	return w, ErrNotAWritableStream
}

func (w WritableStreamDefaultWriter) Closed() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = w.Get("closed"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStreamDefaultWriter) DesiredSize() (int, error) {
	return w.GetAttributeInt("desiredSize")
}

func (w WritableStreamDefaultWriter) Ready() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = w.Get("ready"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStreamDefaultWriter) Abort(reason string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = w.Call("abort", js.ValueOf(reason)); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStreamDefaultWriter) Close() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = w.Call("close"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WritableStreamDefaultWriter) ReleaseLock() error {

	_, err := w.Call("releaseLock")
	return err
}

func (w WritableStreamDefaultWriter) Write(chunk typedarray.Uint8Array) (promise.Promise, error) {

	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = w.Call("write", chunk.JSObject()); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}
