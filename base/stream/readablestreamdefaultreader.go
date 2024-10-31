package stream

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/jserror"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/typedarray"
)

var singletonReadableStreamDefault sync.Once

var readablestreamdefaultinterface js.Value

// GetReadableStreamDefaultReaderInterface
func GetReadableStreamDefaultReaderInterface() js.Value {

	singletonReadableStreamDefault.Do(func() {

		var err error
		if readablestreamdefaultinterface, err = baseobject.Get(js.Global(), "ReadableStreamDefaultReader"); err != nil {
			readablestreamdefaultinterface = js.Undefined()
		}
	})

	return readablestreamdefaultinterface
}

type ReadableStreamDefaultReader struct {
	baseobject.BaseObject
}

type ReadableStreamDefaultReaderFrom interface {
	ReadableStreamDefaultReader_() ReadableStreamDefaultReader
}

func (r ReadableStreamDefaultReader) ReadableStreamDefaultReader_() ReadableStreamDefaultReader {
	return r
}

func NewReadableStreamDefaultReaderFromJSObject(obj js.Value) (ReadableStreamDefaultReader, error) {
	var r ReadableStreamDefaultReader

	if rsi := GetReadableStreamDefaultReaderInterface(); !rsi.IsUndefined() {
		if obj.InstanceOf(rsi) {
			r.BaseObject = r.SetObject(obj)
			return r, nil

		}
	}

	return r, ErrNotAReadableStreamDefaultReader
}

func (r ReadableStreamDefaultReader) newRead(data []byte, dataHandle func([]byte, int)) *promise.Promise {
	var pp *promise.Promise
	var err error
	var promiseread js.Value

	if promiseread, err = r.Call("read"); err == nil {
		var p promise.Promise

		if p, err = promise.NewFromJSObject(promiseread); err == nil {

			newpromise, _ := p.Then(func(i interface{}) *promise.Promise {
				var obj js.Value
				if b, ok := i.(baseobject.ObjectFrom); ok {
					obj = b.JSObject()
					var done bool = false
					if obj.Get("done").Bool() == true {
						done = true
					}

					var u8array typedarray.Uint8Array
					var n int
					uint8arrayObject := obj.Get("value")

					if u8array, err = typedarray.NewUint8Array(uint8arrayObject); err == nil {

						if n, err = u8array.CopyBytes(data); err == nil {
							dataHandle(data, n)
						} else {
							rej, _ := promise.Reject(err)
							return &rej
						}

					}

					if done == false {
						return r.newRead(data, dataHandle)
					} else {
						return nil
					}

				}

				return nil
			}, nil)
			pp = &newpromise

		}
	}
	return pp
}
func (r ReadableStreamDefaultReader) AsyncRead(buffersize int, dataHandle func([]byte, int)) (promise.Promise, error) {

	return promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		var data []byte = make([]byte, buffersize)
		var p *promise.Promise

		p = r.newRead(data, dataHandle)

		p.Then(func(i interface{}) *promise.Promise {
			resolvefunc.Invoke(nil)
			return nil
		}, func(e error) {
			if errjs, err := jserror.New(e); err == nil {
				errfunc.Invoke(errjs.JSObject())
			}
		})

		return nil, nil

	})

}

func (r ReadableStreamDefaultReader) Closed() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = r.Get("closed"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (r ReadableStreamDefaultReader) Cancel() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = r.Call("cancel"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (r ReadableStreamDefaultReader) ReleaseLock() error {

	_, err := r.Call("releaseLock")
	return err
}

func (w ReadableStreamDefaultReader) Read() (promise.Promise, error) {

	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = w.Call("read"); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}
