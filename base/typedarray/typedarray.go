package typedarray

import (
	"errors"
	"syscall/js"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/arraybuffer"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {
	initinterface.RegisterInterface(GetFloat32ArrayInterface)
	initinterface.RegisterInterface(GetFloat64ArrayInterface)
	initinterface.RegisterInterface(GetInt8ArrayInterface)
	initinterface.RegisterInterface(GetInt16ArrayInterface)
	initinterface.RegisterInterface(GetInt32ArrayInterface)
	initinterface.RegisterInterface(GetUint8ArrayInterface)
	initinterface.RegisterInterface(GetUint8ClampedArrayInterface)
	initinterface.RegisterInterface(GetUint16ArrayInterface)
	initinterface.RegisterInterface(GetUint32ArrayInterface)

}

// TypedArray struct
type TypedArray struct {
	array.Array
}

type TypedArrayFrom interface {
	TypedArray_() TypedArray
}

func (t TypedArray) TypedArray_() TypedArray {
	return t
}

func (t TypedArray) Bytes() ([]byte, error) {
	var err error
	var buffer []byte
	var l int
	if l, err = t.Length(); err == nil {
		buffer = make([]byte, l)
		if _, err = baseobject.CopyBytesToGo(buffer, t.JSObject()); err == nil {
			return buffer, nil
		}
	}

	return buffer, err
}

func (t TypedArray) CopyBytes(buffer []byte) (int, error) {

	var err error
	var l int
	if l, err = t.Length(); err == nil {
		if len(buffer) < l {
			return 0, errors.New("Increase your buffer size")
		}

	} else {
		return 0, err
	}

	return baseobject.CopyBytesToGo(buffer, t.JSObject())

}

func (t TypedArray) CopyFromBytes(buffer []byte) (int, error) {

	var err error
	var l int
	if l, err = t.Length(); err == nil {
		if len(buffer) < l {
			return 0, errors.New("Increase your buffer size")
		}

	} else {
		return 0, err
	}

	return baseobject.CopyBytesToJS(t.JSObject(), buffer)
}

func (t TypedArray) Buffer() (arraybuffer.ArrayBuffer, error) {

	var err error
	var obj js.Value

	if obj, err = t.Get("buffer"); err == nil {

		return arraybuffer.NewFromJSObject(obj)

	}

	return arraybuffer.ArrayBuffer{}, err
}

func (t TypedArray) ByteLength() (int64, error) {

	return t.GetAttributeInt64("byteLength")
}

func (t TypedArray) ByteOffset() (int64, error) {

	return t.GetAttributeInt64("byteOffset")
}

func (t TypedArray) BYTES_PER_ELEMENT() (int, error) {

	return t.GetAttributeInt("BYTES_PER_ELEMENT")
}

func (t TypedArray) Subarray(opts ...int) (interface{}, error) {

	var err error
	var arrayJS []interface{}
	var obj js.Value
	var newArr interface{}

	if len(opts) < 3 {
		for _, opt := range opts {
			arrayJS = append(arrayJS, js.ValueOf(opt))
		}
	}

	if obj, err = t.Call("subarray", arrayJS...); err == nil {
		newArr, err = baseobject.Discover(obj)

	}

	return newArr, err
}

func newTypedArrayFrom(interfacejs js.Value, f func(js.Value) (interface{}, error), iterable interface{}) (interface{}, error) {
	var obj js.Value
	var err error
	var newArr interface{}
	if obj, err = baseobject.Call(interfacejs, "from", baseobject.GetJsValueOf(iterable)); err == nil {
		newArr, err = f(obj)
	}
	return newArr, err
}

func newTypedArrayOf(interfacejs js.Value, f func(js.Value) (interface{}, error), values ...interface{}) (interface{}, error) {

	var newArr interface{}
	var arrayJS []interface{}
	var obj js.Value
	var err error
	for _, value := range values {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}
	if obj, err = baseobject.Call(interfacejs, "of", arrayJS...); err == nil {
		newArr, err = f(obj)
	}
	return newArr, err
}
