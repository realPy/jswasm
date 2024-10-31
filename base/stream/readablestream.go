package stream

// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/promise"
)

func init() {

	initinterface.RegisterInterface(GetRInterface)
	initinterface.RegisterInterface(GetWInterface)
	initinterface.RegisterInterface(GetTInterface)
	initinterface.RegisterInterface(GetReadableStreamDefaultReaderInterface)
	initinterface.RegisterInterface(GetWritableStreamDefaultWriterInterface)
	initinterface.RegisterInterface(GetTransformStreamDefaultControllerInterface)
}

var singletonr sync.Once

var readablestreaminterface js.Value

// GetRInterface get the JS interface ReadableStream.
func GetRInterface() js.Value {

	singletonr.Do(func() {

		var err error
		if readablestreaminterface, err = baseobject.Get(js.Global(), "ReadableStream"); err != nil {
			readablestreaminterface = js.Undefined()
		}
		baseobject.Register(readablestreaminterface, func(v js.Value) (interface{}, error) {
			return NewReadableStreamFromJSObject(v)
		})
	})

	return readablestreaminterface
}

type ReadableStream struct {
	baseobject.BaseObject
}

type ReadableStreamFrom interface {
	ReadableStream_() ReadableStream
}

func (r ReadableStream) ReadableStream_() ReadableStream {
	return r
}

func (r ReadableStream) Locked() (bool, error) {
	return r.GetAttributeBool("locked")
}

// NewReadableStream Create a new ReadableStream
func NewReadableStream() (ReadableStream, error) {
	var r ReadableStream
	var obj js.Value
	var err error
	if ri := GetRInterface(); !ri.IsUndefined() {

		if obj, err = baseobject.New(ri); err == nil {
			r.BaseObject = r.SetObject(obj)
		}

	} else {
		err = ErrNotImplementedReadableStream
	}
	return r, err
}

func NewReadableStreamFromJSObject(obj js.Value) (ReadableStream, error) {
	var r ReadableStream
	var err error
	if rsi := GetRInterface(); !rsi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(rsi) {
				r.BaseObject = r.SetObject(obj)

			} else {
				err = ErrNotAReadableStream
			}
		}
	} else {
		err = ErrNotImplementedReadableStream
	}

	return r, err
}

func (r ReadableStream) Cancel() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = r.Call("cancel"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (r ReadableStream) GetReader() (ReadableStreamDefaultReader, error) {
	var err error
	var obj js.Value

	if obj, err = r.Call("getReader"); err == nil {
		return NewReadableStreamDefaultReaderFromJSObject(obj)

	}
	return ReadableStreamDefaultReader{}, err

}

func (r ReadableStream) Tee() ([]ReadableStream, error) {
	var err error
	var obj js.Value
	var ret []ReadableStream
	var a array.Array

	if obj, err = r.Call("tee"); err == nil {

		if a, err = array.NewFromJSObject(obj); err == nil {

			a.ForEach(func(i interface{}) {

				if r, ok := i.(ReadableStreamFrom); ok {
					ret = append(ret, r.ReadableStream_())
				}

			})
		}

	}
	return ret, err

}

func (r ReadableStream) PipeThrough(t TransformStream, options ...map[string]string) (ReadableStream, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var transformread ReadableStream

	arrayJS = append(arrayJS, t.JSObject())

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj, err = r.Call("pipeThrough", arrayJS...); err == nil {

		transformread, err = NewReadableStreamFromJSObject(obj)

	}

	return transformread, err
}

func (r ReadableStream) PipeTo(w WritableStream, options ...map[string]string) (promise.Promise, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var finalpromise promise.Promise

	arrayJS = append(arrayJS, w.JSObject())

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj, err = r.Call("pipeTo", arrayJS...); err == nil {

		finalpromise, err = promise.NewFromJSObject(obj)

	}

	return finalpromise, err
}
