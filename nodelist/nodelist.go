package nodelist

// https://developer.mozilla.org/fr/docs/Web/API/NodeList

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/node"
)

var singleton sync.Once

var nodelistinterface js.Value

//NodeList struct
type NodeList struct {
	baseobject.BaseObject
}

type NodeListFrom interface {
	NodeList() NodeList
}

func (n NodeList) NodeList() NodeList {
	return n
}

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if nodelistinterface, err = js.Global().GetWithErr("NodeList"); err != nil {
			nodelistinterface = js.Null()
		}
		baseobject.Register(nodelistinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return nodelistinterface
}

func NewFromJSObject(obj js.Value) (NodeList, error) {
	var n NodeList

	if nli := GetInterface(); !nli.IsNull() {
		if obj.InstanceOf(nli) {
			n.BaseObject = n.SetObject(obj)
			return n, nil
		}
	}
	return n, ErrNotAnNodeList
}

func (n NodeList) Item(index int) (node.Node, error) {

	return node.NewFromJSObject(n.JSObject().Index(index))
}
