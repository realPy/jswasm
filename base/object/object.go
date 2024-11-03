package object

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/array"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/objectmap"
)

func init() {
	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var objectinterface js.Value

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if objectinterface, err = baseobject.Get(js.Global(), "Object"); err != nil {
			objectinterface = js.Undefined()
		}
		baseobject.Register(objectinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return objectinterface
}

// Object struct
type Object struct {
	baseobject.BaseObject
}

type ObjectFrom interface {
	Object_() Object
}

func (o Object) Object_() Object {
	return o
}

func New() (Object, error) {
	var o Object
	var err error
	var obj js.Value
	if ai := GetInterface(); !ai.IsUndefined() {

		if obj, err = baseobject.New(ai); err == nil {
			o.BaseObject = o.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func NewFromJSObject(obj js.Value) (Object, error) {
	var o Object
	var err error
	if ai := GetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				o.BaseObject = o.SetObject(obj)

			} else {
				err = ErrNotAnObject
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return o, err
}

func (o Object) Keys() (array.Array, error) {

	var err error
	var obj js.Value
	var newArr array.Array

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj, err = baseobject.Call(ai, "keys", o.JSObject()); err == nil {
			newArr, err = array.NewFromJSObject(obj)

		}

	}

	return newArr, err
}

func (o Object) Values() (array.Array, error) {

	var err error
	var obj js.Value
	var newArr array.Array

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj, err = baseobject.Call(ai, "values", o.JSObject()); err == nil {
			newArr, err = array.NewFromJSObject(obj)

		}

	}

	return newArr, err
}

func (o Object) Map() (objectmap.ObjectMap, error) {
	var err error
	var obj js.Value
	var newMap objectmap.ObjectMap

	if ai := GetInterface(); !ai.IsUndefined() {
		if obj, err = baseobject.Call(ai, "entries", o.JSObject()); err == nil {
			newMap, err = objectmap.New(obj)

		}

	}
	return newMap, err
}

func (o Object) Delete(property string) error {
	_, err := baseobject.Delete(o.JSObject(), property)
	return err
}
