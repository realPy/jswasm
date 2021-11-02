package history

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

// https://developer.mozilla.org/fr/docs/Web/API/History_API

var singleton sync.Once

// var historyinterface *JSInterface

var historyinterface js.Value

//JSInterface JSInterface struct
// type JSInterface struct {
// 	objectInterface js.Value
// }

//History struct
type History struct {
	baseobject.BaseObject
}

type HistoryFrom interface {
	History_() History
}

func (h History) History_() History {
	return h
}

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if historyinterface, err = baseobject.Get(js.Global(), "History"); err != nil {
			historyinterface = js.Undefined()
		}
		baseobject.Register(historyinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return historyinterface
}

func NewFromJSObject(obj js.Value) (History, error) {
	var h History
	var err error
	if hci := GetInterface(); !hci.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(hci) {

				h.BaseObject = h.SetObject(obj)

			} else {
				err = ErrNotAnHistory
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return h, err
}

func (h History) Forward() error {
	var err error
	_, err = h.Call("forward")
	return err
}

func (h History) Back() error {
	var err error
	_, err = h.Call("back")
	return err
}

func (h History) Go(position int) error {
	var err error
	_, err = h.Call("go", js.ValueOf(position))

	return err
}

func (h History) Length() (int, error) {
	var err error

	obj, err := h.Get("length")

	return obj.Int(), err
}

func (h History) PushState(obj interface{}, name string, page string) error {
	var err error
	_, err = h.Call("pushState", js.ValueOf(obj), js.ValueOf(name), js.ValueOf(page))

	return err
}

func (h History) ReplaceState(obj interface{}, name string, page string) error {
	var err error
	_, err = h.Call("replaceState", js.ValueOf(obj), js.ValueOf(name), js.ValueOf(page))

	return err
}

func (h History) State() (interface{}, error) {
	var err error
	var obj js.Value

	obj, err = h.Get("state")

	return obj, err
}
