package stream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
)

var singletonTransformStreamDefaultController sync.Once

var transformstreamdefaultcontrollerinterface js.Value

// GetTransformStreamDefaultControllerInterface
func GetTransformStreamDefaultControllerInterface() js.Value {
	singletonTransformStreamDefaultController.Do(func() {

		var err error
		if transformstreamdefaultcontrollerinterface, err = baseobject.Get(js.Global(), "TransformStreamDefaultController"); err != nil {
			transformstreamdefaultcontrollerinterface = js.Undefined()
		}

		baseobject.Register(transformstreamdefaultcontrollerinterface, func(v js.Value) (interface{}, error) {
			return NewTransformStreamDefaultControllerFromJSObject(v)
		})
	})

	return transformstreamdefaultcontrollerinterface
}

type TransformStreamDefaultController struct {
	baseobject.BaseObject
}

type TransformStreamDefaultControllerFrom interface {
	TransformStreamDefaultController_() TransformStreamDefaultController
}

func (t TransformStreamDefaultController) TransformStreamDefaultController_() TransformStreamDefaultController {
	return t
}

func NewTransformStreamDefaultControllerFromJSObject(obj js.Value) (TransformStreamDefaultController, error) {
	var t TransformStreamDefaultController
	var err error
	if wsi := GetTransformStreamDefaultControllerInterface(); !wsi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(wsi) {
				t.BaseObject = t.SetObject(obj)

			} else {
				err = ErrNotATransformStreamDefaultController
			}
		}
	} else {
		err = ErrNotImplementedTransformStream
	}

	return t, err
}

func (t TransformStreamDefaultController) Enqueue(chunk baseobject.BaseObject) error {
	var err error

	if _, err = t.Call("enqueue", chunk.JSObject()); err != nil {
		return err

	}

	return nil
}

func (t TransformStreamDefaultController) Terminate() error {
	var err error

	if _, err = t.Call("terminate"); err != nil {
		return err

	}

	return nil
}
