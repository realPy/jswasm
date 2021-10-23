package dragevent

// https://developer.mozilla.org/en-US/docs/Web/API/DragEvent

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/datatransfer"
	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var drageventinterface js.Value

//DragEvent DragEvent struct
type DragEvent struct {
	event.Event
}

type DragEventFrom interface {
	DragEvent_() DragEvent
}

func (d DragEvent) DragEvent_() DragEvent {
	return d
}

//GetInterface get teh JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if drageventinterface, err = js.Global().GetWithErr("DragEvent"); err != nil {
			drageventinterface = js.Undefined()
		}
		baseobject.Register(drageventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return drageventinterface
}

func NewFromJSObject(obj js.Value) (DragEvent, error) {
	var e DragEvent

	if di := GetInterface(); !di.IsUndefined() {
		if obj.InstanceOf(di) {
			e.BaseObject = e.SetObject(obj)
			return e, nil
		}
	}
	return e, ErrNotAnDragEvent
}

func (d DragEvent) DataTransfer() (datatransfer.DataTransfer, error) {

	var err error
	var obj js.Value

	if obj, err = d.JSObject().GetWithErr("dataTransfer"); err == nil {

		return datatransfer.NewFromJSObject(obj)
	}
	return datatransfer.DataTransfer{}, err

}
