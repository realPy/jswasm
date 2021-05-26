package iterator

import (
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

//Iterator iterator
type Iterator struct {
	baseobject.BaseObject
}

func NewFromJSObject(obj js.Value) Iterator {
	var i Iterator
	i.BaseObject = i.SetObject(obj)
	return i
}

func pairValues(obj js.Value) (int, interface{}) {

	var value interface{}
	var index int = -1

	if obj.Type() == js.TypeObject {
		if obj.Length() == 2 {

			index = obj.Index(0).Int()

			value = baseobject.GoValue(obj.Index(1))

		}

	}
	return index, value
}

/* Parse using

for index, value, err := it.Next(); err == nil; index, value, err = it.Next() {

}
*/

func (i Iterator) Next() (int, interface{}, error) {

	var err error
	var done bool = true
	var obj, doneobj, valueobj js.Value
	var index int = -1
	var value interface{}

	if obj, err = i.JSObject().CallWithErr("next"); err == nil {

		if doneobj, err = obj.GetWithErr("done"); err == nil {
			if doneobj.Type() == js.TypeBoolean {
				done = doneobj.Bool()
			} else {
				err = baseobject.ErrObjectNotBool
			}
		}
		if done {
			err = EOI

		} else {

			if valueobj, err = obj.GetWithErr("value"); err == nil {
				if valueobj.Type() == js.TypeObject {
					index, value = pairValues(valueobj)
				} else {
					value = baseobject.GoValue(valueobj)
				}

			}
		}

	}
	return index, value, err
}
