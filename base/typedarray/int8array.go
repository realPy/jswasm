package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
)

var singletonint8array sync.Once

var int8arrayinterface js.Value

// Uint8Array struct
type Int8Array struct {
	TypedArray
}

type Int8ArrayFrom interface {
	Uint8Array_() Uint8Array
}

func (a Int8Array) Int8Array_() Int8Array {
	return a
}

// GetInterface get the JS interface
func GetInt8ArrayInterface() js.Value {

	singletonint8array.Do(func() {

		var err error
		if int8arrayinterface, err = baseobject.Get(js.Global(), "Int8Array"); err != nil {
			int8arrayinterface = js.Undefined()
		}
		baseobject.Register(int8arrayinterface, func(v js.Value) (interface{}, error) {
			return NewInt8FromJSObject(v)
		})

	})

	return int8arrayinterface
}

func NewInt8Array(value interface{}) (Int8Array, error) {

	var a Int8Array
	var objnew js.Value
	var err error
	if ai := GetInt8ArrayInterface(); !ai.IsUndefined() {
		if objnew, err = baseobject.New(ai, baseobject.GetJsValueOf(value)); err == nil {
			a.BaseObject = a.SetObject(objnew)
		}

	} else {
		err = ErrNotImplementedInt8Array
	}

	return a, err
}

func NewInt8ArrayFrom(iterable interface{}) (Int8Array, error) {
	arr, err := newTypedArrayFrom(GetInt8ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewInt8FromJSObject(v)
	}, iterable)
	return arr.(Int8Array), err
}

func NewInt8ArrayOf(values ...interface{}) (Int8Array, error) {

	arr, err := newTypedArrayOf(GetInt8ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewInt8FromJSObject(v)
	}, values...)
	return arr.(Int8Array), err
}

func NewInt8FromJSObject(obj js.Value) (Int8Array, error) {
	var u Int8Array
	var err error
	if ui := GetInt8ArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.BaseObject = u.SetObject(obj)
			} else {
				err = ErrNotAInt8Array
			}
		}
	} else {
		err = ErrNotImplementedInt8Array
	}

	return u, err
}
